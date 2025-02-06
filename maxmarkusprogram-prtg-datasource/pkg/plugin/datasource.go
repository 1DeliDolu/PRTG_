package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/maxmarkusprogram/prtg/pkg/models"
)

// Make sure Datasource implements required interfaces. This is important to do
// since otherwise we will only get a not implemented error response from plugin in
// runtime. In this example datasource instance implements backend.QueryDataHandler,
// backend.CheckHealthHandler interfaces. Plugin should not implement all these
// interfaces - only those which are required for a particular task.
var (
	_ backend.QueryDataHandler      = (*Datasource)(nil)
	_ backend.CheckHealthHandler    = (*Datasource)(nil)
	_ instancemgmt.InstanceDisposer = (*Datasource)(nil)
	_ backend.CallResourceHandler   = (*Datasource)(nil)
)



// NewDatasource creates a new datasource instance.
func NewDatasource(_ context.Context, settings backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	config, err := models.LoadPluginSettings(settings)
	if err != nil {
		return nil, err
	}

	baseURL := fmt.Sprintf("https://%s", config.Path)

	fmt.Println("baseURL: ", baseURL)

	// Default cache time if not set
	cacheTime := config.CacheTime
	if cacheTime <= 0 {
		cacheTime = 30 * time.Second // default cache time
	}

	return &Datasource{
		baseURL: baseURL,
		api:     NewApi(baseURL, config.Secrets.ApiKey, cacheTime, 10*time.Second),
	}, nil
}

// Dispose here tells plugin SDK that plugin wants to clean up resources when a new instance
// created. As soon as datasource settings change detected by SDK old datasource instance will
// be disposed and a new one will be created using NewSampleDatasource factory function.
func (d *Datasource) Dispose() {
	// Clean up datasource instance resources.
}

// QueryData handles multiple queries and returns multiple responses.
// req contains the queries []DataQuery (where each query contains RefID as a unique identifier).
// The QueryDataResponse contains a map of RefID to the response for each query, and each response
// contains Frames ([]*Frame).
func (d *Datasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	// create response struct
	response := backend.NewQueryDataResponse()

	// loop over queries and execute them individually.
	for _, q := range req.Queries {
		res := d.query(ctx, req.PluginContext, q)

		// save the response in a hashmap
		// based on with RefID as identifier
		response.Responses[q.RefID] = res
	}

	return response, nil
}



func (d *Datasource) query(_ context.Context, _ backend.PluginContext, query backend.DataQuery) backend.DataResponse {
	var response backend.DataResponse

	// Unmarshal the JSON into our queryModel.
	var qm queryModel

	err := json.Unmarshal(query.JSON, &qm)
	if err != nil {
		return backend.ErrDataResponse(backend.StatusBadRequest, fmt.Sprintf("json unmarshal: %v", err.Error()))
	}

	// create data frame response.
	frame := data.NewFrame("response")

	// add fields based on query type
	switch qm.QueryType {
	case "Metrics":
		historicalData, err := d.api.GetHistoricalData(qm.Sensor, query.TimeRange.From, query.TimeRange.To)
		if err != nil {
			return backend.ErrDataResponse(backend.StatusInternal, fmt.Sprintf("error fetching historical data: %v", err))
		}

		times := make([]time.Time, len(historicalData.HistData))
		values := make([]float64, len(historicalData.HistData))
		for i, data := range historicalData.HistData {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", data.Datetime)
			if err != nil {
				return backend.ErrDataResponse(backend.StatusInternal, fmt.Sprintf("error parsing datetime: %v", err))
			}
			times[i] = parsedTime
			values[i] = data.Value
		}

		frame.Fields = append(frame.Fields,
			data.NewField("time", nil, times),
			data.NewField("values", nil, values),
		)
		if qm.IncludeGroupName {
			frame.Fields = append(frame.Fields, data.NewField("group", nil, []string{qm.Group}))
		}
		if qm.IncludeDeviceName {
			frame.Fields = append(frame.Fields, data.NewField("device", nil, []string{qm.Device}))
		}
		if qm.IncludeSensorName {
			frame.Fields = append(frame.Fields, data.NewField("sensor", nil, []string{qm.Sensor}))
		}
	case "Text":
		switch qm.Property {
		case "group":
			groupData, err := d.api.GetGroups()
			if err != nil {
				return backend.ErrDataResponse(backend.StatusInternal, fmt.Sprintf("error fetching group data: %v", err))
			}
			filterProperties := extractFilterProperties(groupData.Groups, qm.FilterProperty)
			frame.Fields = append(frame.Fields,
				data.NewField("property", nil, []string{qm.Property}),
				data.NewField("filterProperty", nil, filterProperties),
			)
		case "device":
			deviceData, err := d.api.GetDevices()
			if err != nil {
				return backend.ErrDataResponse(backend.StatusInternal, fmt.Sprintf("error fetching device data: %v", err))
			}
			filterProperties := extractFilterProperties(deviceData.Devices, qm.FilterProperty)
			frame.Fields = append(frame.Fields,
				data.NewField("property", nil, []string{qm.Property}),
				data.NewField("filterProperty", nil, filterProperties),
			)
		case "sensor":
			sensorData, err := d.api.GetSensors()
			if err != nil {
				return backend.ErrDataResponse(backend.StatusInternal, fmt.Sprintf("error fetching sensor data: %v", err))
			}
			filterProperties := extractFilterProperties(sensorData.Sensors, qm.FilterProperty)
			frame.Fields = append(frame.Fields,
				data.NewField("property", nil, []string{qm.Property}),
				data.NewField("filterProperty", nil, filterProperties),
			)
		}
	case "Raw":
		switch qm.Property {
		case "group":
			groupData, err := d.api.GetGroups()
			if err != nil {
				return backend.ErrDataResponse(backend.StatusInternal, fmt.Sprintf("error fetching group data: %v", err))
			}
			filterProperties := extractFilterProperties(groupData.Groups, qm.FilterProperty + "raw")
			frame.Fields = append(frame.Fields,
				data.NewField("property", nil, []string{qm.Property + "raw"}),
				data.NewField("filterProperty", nil, filterProperties),
			)
		case "device":
			deviceData, err := d.api.GetDevices()
			if err != nil {
				return backend.ErrDataResponse(backend.StatusInternal, fmt.Sprintf("error fetching device data: %v", err))
			}
			filterProperties := extractFilterProperties(deviceData.Devices, qm.FilterProperty + "raw")
			frame.Fields = append(frame.Fields,
				data.NewField("property", nil, []string{qm.Property + "raw"}),
				data.NewField("filterProperty", nil, filterProperties),
			)
		case "sensor":
			sensorData, err := d.api.GetSensors()
			if err != nil {
				return backend.ErrDataResponse(backend.StatusInternal, fmt.Sprintf("error fetching sensor data: %v", err))
			}
			filterProperties := extractFilterProperties(sensorData.Sensors, qm.FilterProperty + "raw")
			frame.Fields = append(frame.Fields,
				data.NewField("property", nil, []string{qm.Property + "raw"}),
				data.NewField("filterProperty", nil, filterProperties),
			)
		}
	}

	// add the frames to the response.
	response.Frames = append(response.Frames, frame)

	return response
}

// extractFilterProperties extracts the filter properties from the given data
func extractFilterProperties(data interface{}, _ string) []string {
	var filterProperties []string
	switch v := data.(type) {
	case []Group:
		filterProperties = make([]string, len(v))
		for i, item := range v {
			filterProperties[i] = item.Group
		}
	case []Device:
		filterProperties = make([]string, len(v))
		for i, item := range v {
			filterProperties[i] = item.Device
		}
	case []Sensor:
		filterProperties = make([]string, len(v))
		for i, item := range v {
			filterProperties[i] = item.Sensor
		}
	}
	return filterProperties
}

/* ########################################## CHECK HEALTH  ############################################ */

// CheckHealth handles health checks sent from Grafana to the plugin.
// The main use case for these health checks is the test button on the
// datasource configuration page which allows users to verify that
// a datasource is working as expected.
func (d *Datasource) CheckHealth(_ context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	res := &backend.CheckHealthResult{}
	config, err := models.LoadPluginSettings(*req.PluginContext.DataSourceInstanceSettings)

	if err != nil {
		res.Status = backend.HealthStatusError
		res.Message = "Unable to load settings"
		return res, nil
	}

	if config.Secrets.ApiKey == "" {
		res.Status = backend.HealthStatusError
		res.Message = "API key is missing"
		return res, nil
	}

	return &backend.CheckHealthResult{
		Status:  backend.HealthStatusOk,
		Message: "Data source is working",
	}, nil
}

/* ########################################## CALL RESOURCE   ############################################ */

func (d *Datasource) CallResource(ctx context.Context, req *backend.CallResourceRequest, sender backend.CallResourceResponseSender) error {
	switch req.Path {
	case "groups":
		return d.handleGetGroups(sender)
	case "devices":
		return d.handleGetDevices(sender)
	case "sensors":
		return d.handleGetSensors(sender)
	case "channels":
		objid := req.URL
		return d.handleGetChannel(sender, objid)
	default:
		return sender.Send(&backend.CallResourceResponse{
			Status: http.StatusNotFound,
		})
	}
}

func (d *Datasource) handleGetGroups(sender backend.CallResourceResponseSender) error {
	groups, err := d.api.GetGroups()
	if (err != nil) {
		return sender.Send(&backend.CallResourceResponse{
			Status: http.StatusInternalServerError,
			Body:   []byte(err.Error()),
		})
	}

	body, err := json.Marshal(groups)
	if (err != nil) {
		return sender.Send(&backend.CallResourceResponse{
			Status: http.StatusInternalServerError,
			Body:   []byte(fmt.Sprintf("error marshaling groups: %v", err)),
		})
	}

	return sender.Send(&backend.CallResourceResponse{
		Status: http.StatusOK,
		Headers: map[string][]string{
			"Content-Type": {"application/json"},
		},
		Body: body,
	})
}

func (d *Datasource) handleGetDevices(sender backend.CallResourceResponseSender) error {
	devices, err := d.api.GetDevices()
	if (err != nil) {
		return sender.Send(&backend.CallResourceResponse{
			Status: http.StatusInternalServerError,
			Body:   []byte(err.Error()),
		})
	}

	body, err := json.Marshal(devices)
	if (err != nil) {
		return sender.Send(&backend.CallResourceResponse{
			Status: http.StatusInternalServerError,
			Body:   []byte(fmt.Sprintf("error marshaling devices: %v", err)),
		})
	}

	return sender.Send(&backend.CallResourceResponse{
		Status: http.StatusOK,
		Headers: map[string][]string{
			"Content-Type": {"application/json"},
		},
		Body: body,
	})
}

func (d *Datasource) handleGetSensors(sender backend.CallResourceResponseSender) error {
	sensors, err := d.api.GetSensors()
	if (err != nil) {
		return sender.Send(&backend.CallResourceResponse{
			Status: http.StatusInternalServerError,
			Body:   []byte(err.Error()),
		})
	}

	body, err := json.Marshal(sensors)
	if (err != nil) {
		return sender.Send(&backend.CallResourceResponse{
			Status: http.StatusInternalServerError,
			Body:   []byte(fmt.Sprintf("error marshaling sensors: %v", err)),
		})
	}

	return sender.Send(&backend.CallResourceResponse{
		Status: http.StatusOK,
		Headers: map[string][]string{
			"Content-Type": {"application/json"},
		},
		Body: body,
	})
}

func (d *Datasource) handleGetChannel(sender backend.CallResourceResponseSender, sensorId string) error {
	channels, err := d.api.GetChannels(sensorId)
	if (err != nil) {
		return sender.Send(&backend.CallResourceResponse{
			Status: http.StatusInternalServerError,
			Body:   []byte(err.Error()),
		})
	}

	body, err := json.Marshal(channels)
	if (err != nil) {
		return sender.Send(&backend.CallResourceResponse{
			Status: http.StatusInternalServerError,
			Body:   []byte(fmt.Sprintf("error marshaling channels: %v", err)),
		})
	}

	return sender.Send(&backend.CallResourceResponse{
		Status: http.StatusOK,
		Headers: map[string][]string{
			"Content-Type": {"application/json"},
		},
		Body: body,
	})
}
