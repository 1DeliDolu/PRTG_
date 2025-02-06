package plugin

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Api struct to hold API related configurations
type Api struct {
	baseURL string
	apiKey  string
	timeout time.Duration
}

// NewApi creates a new Api instance
func NewApi(baseURL, apiKey string, timeout, requestTimeout time.Duration) *Api {
	return &Api{
		baseURL: baseURL,
		apiKey:  apiKey,
		timeout: requestTimeout,
	}
}

// buildApiUrl creates a standardized PRTG API URL
func (a *Api) buildApiUrl(method string, params map[string]string) (string, error) {
	baseUrl := fmt.Sprintf("%s/api/%s", a.baseURL, method)
	u, err := url.Parse(baseUrl)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	// Add query parameters
	q := url.Values{}
	q.Set("apitoken", a.apiKey)

	// Add additional parameters
	for key, value := range params {
		q.Set(key, value)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}

// SetTimeout sets the API request timeout
func (a *Api) SetTimeout(timeout time.Duration) {
	if timeout > 0 {
		a.timeout = timeout
	}
}

// baseExecuteRequest handles the common HTTP request logic
func (a *Api) baseExecuteRequest(endpoint string, params map[string]string) ([]byte, error) {
	apiUrl, err := a.buildApiUrl(endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("failed to build URL: %w", err)
	}

	// Disable TLS verification (for self-signed certificates)
	client := &http.Client{
		Timeout: a.timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden {
		return nil, fmt.Errorf("access denied: please verify API token and permissions")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func (a *Api) GetStatusList() (*PrtgStatusListResponse, error) {
	body, err := a.baseExecuteRequest("status.json", nil)
	if err != nil {
		return nil, err
	}

	var response PrtgStatusListResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &response, nil
}

func (a *Api) GetGroups() (*PrtgGroupListResponse, error) {
	params := map[string]string{
		"content": "groups",
		"columns": "objid,group,device,sensor,channel,active,message,priority,status,tags,datetime,upsens,downsens,warnsens,pausedsens,unusualsens,totalsens,accessrights",
	}

	body, err := a.baseExecuteRequest("table.json", params)
	if err != nil {
		return nil, err
	}

	var response PrtgGroupListResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &response, nil
}

func (a *Api) GetDevices() (*PrtgDevicesListResponse, error) {
	params := map[string]string{
		"content": "devices",
		"columns": "accessrights,active,channel,datetime,device,deviceicon,downsens,group,location,message,objid,pausedsens,priority,sensor,status,tags,totalsens,unusualsens,upsens,warnsens",
	}

	body, err := a.baseExecuteRequest("table.json", params)
	if err != nil {
		return nil, err
	}

	var response PrtgDevicesListResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &response, nil
}

func (a *Api) GetSensors() (*PrtgSensorsListResponse, error) {
	params := map[string]string{
		"content": "sensors",
		"columns": "objid,group,device,sensor,channel,active,status,message,priority,tags,datetime,lastcheck,lastup,lastdown,interval,uptime,uptimetime,uptimesince,downtime,downtimetime,downtimesince,upsens,downsens,warnsens,pausedsens,unusualsens,totalsens,accessrights,parentid",
	}

	body, err := a.baseExecuteRequest("table.json", params)
	if err != nil {
		return nil, err
	}

	var response PrtgSensorsListResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &response, nil
}

func (a *Api) GetChannels(sensorId string) (*PrtgChannelValueStruct, error) {
	params := map[string]string{
		"content":    "values",
		"columns":    "value_,datetime",
		"usecaption": "true",
		"count":      "1",
		"id":         sensorId,
	}

	body, err := a.baseExecuteRequest("table.json", params)
	if err != nil {
		return nil, err
	}

	var response PrtgChannelValueStruct
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &response, nil
}



// GetHistoricalData performs a historical data query for a specific sensor
func (a *Api) GetHistoricalData(sensorID string, startDate, endDate time.Time) (*PrtgHistoricalDataResponse, error) {
	if sensorID == "" {
		return nil, fmt.Errorf("invalid query: missing sensor ID")
	}

	// Calculate time difference in hours
	hours := endDate.Sub(startDate).Hours()

	// Determine averaging interval based on time range
	var avg string
	switch {
	case hours > 12 && hours < 36:
		avg = "300"
	case hours > 36 && hours < 745:
		avg = "3600"
	case hours > 745:
		avg = "86400"
	default:
		avg = "0"
	}

	// Format dates in PRTG format (YYYY-MM-DD-HH-mm-ss)
	formatDate := func(t time.Time) string {
		return t.Format("2025-02-06-15-04-05") // PRTG format yyyy-mm-dd-hh-mm-ss
	}

	// Prepare parameters
	params := map[string]string{
		"id":         sensorID,
		"avg":        avg,
		"sdate":      formatDate(startDate),
		"edate":      formatDate(endDate),
		"count":      "50000",
		"usecaption": "1",
		"columns":    "datetime,value_",
	}

	body, err := a.baseExecuteRequest("historicdata.json", params)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch historical data: %w", err)
	}

	var response PrtgHistoricalDataResponse
		if err := json.Unmarshal(body, &response); err != nil {
			return nil, fmt.Errorf("failed to parse historical data response: %w", err)
		}
	
		if len(response.HistData) == 0 {
			return nil, fmt.Errorf("no historical data received from PRTG")
		}

	return &response, nil
}