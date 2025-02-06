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
		"columns": "objid,group,device,sensor,channel,active,message,priority,status," +
			"tags,datetime,deviceicon,location,upsens,downsens,warnsens,pausedsens,unusualsens,totalsens,accessrights",
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

func (a *Api) GetChanelValues(sensorId string) (*PrtgChannelValueStruct, error) {
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

func (a *Api) GetChannels(sensorId string) (*PrtgChannelValueStruct, error) {
	params := map[string]string{
		"content": "channels",
		"columns": "objid,channel",
		"id":      sensorId,
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
