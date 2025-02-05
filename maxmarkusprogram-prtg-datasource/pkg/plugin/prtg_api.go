package plugin

import (
	"fmt"
	"net/url"
	"time"
)

// Client verwaltet die Serverdaten und HTTP-Client-Parameter für PRTG.
type Client struct {
	Server       string // PRTG-Server-URL
	Username     string // Benutzername
	Password     string // Passwort
	PasswordHash string // Passwort-Hash (optional)
	Timeout      time.Duration // Timeout für Anfragen
}

var (
	defaultTimeout         = 10 * time.Second
	deltaHistoricThreshold = 31 * 24 * time.Hour // 31 Tage
	dateFormat             = "2006-01-02-15-04-05"
)

const (
	GetSensorDetailsEndpoint     = "/api/getsensordetails.json"
	GetSensorDetailsEndpointXML  = "/api/getsensordetails.xml"
	GetTableListsEndpoint        = "/api/table.json"
	GetHistoricDatasEndpoint     = "/api/historicdata.json"
	GetHistoricDatasEndpointXML  = "/api/historicdata.xml"
	GetSensorTreesEndpoint       = "/api/table.xml"
	userAgent                    = "golang-prtg-api"
)

// PrtgSensorTreeResponse represents the response structure for sensor tree data
type PrtgSensorTreeResponse struct {
	Prtg struct {
		TreeSize int    `json:"treesize"`
		Version  string `json:"version"`
		Items    []struct {
			Name     string `json:"name"`
			SensorID int64  `json:"objid"`
			Type     string `json:"type"`
			Tags     string `json:"tags"`
		} `json:"items"`
	} `json:"prtg"`
}

// NewClient erstellt eine neue PRTG-Client-Instanz.
func NewClient(server, username, password string) *Client {
	return &Client{
		Server:   server,
		Username: username,
		Password: password,
		Timeout:  defaultTimeout,
	}
}

// NewClientWithHashedPass erstellt eine neue Client-Instanz mit Passwort-Hash.
func NewClientWithHashedPass(server, username, passwordHash string) *Client {
	return &Client{
		Server:       server,
		Username:     username,
		PasswordHash: passwordHash,
		Timeout:      defaultTimeout,
	}
}

// SetContextTimeout setzt das Timeout für HTTP-Anfragen.
func (c *Client) SetContextTimeout(timeout time.Duration) {
	if timeout <= 0 {
		c.Timeout = defaultTimeout
	} else {
		c.Timeout = timeout
	}
}

func (c *Client) getTemplateUrlQuery() url.Values {
	q := url.Values{}
	q.Set("username", c.Username)
	if (c.Password != "") {
		q.Set("password", c.Password)
	}
	if (c.PasswordHash != "") {
		q.Set("passhash", c.PasswordHash)
	}
	return q
}

func (c *Client) getCompleteUrl(p string, q url.Values) (string, error) {
	u, err := url.Parse(c.Server)
	if err != nil {
		return "", fmt.Errorf("ungültige URL: %w", err)
	}
	u.Path = p
	u.RawQuery = q.Encode()
	return u.String(), nil
}

// GetPrtgVersion gibt die PRTG-Version des angegebenen Servers zurück.
func (c *Client) GetPrtgVersion() (string, error) {
	q := c.getTemplateUrlQuery()
	q.Set("id", "0")

	u, err := c.getCompleteUrl(GetSensorDetailsEndpoint, q)
	if err != nil {
		return "", err
	}

	var sensorDetailResp prtgSensorDetailsResponse
	if err := getPrtgResponse(u, c.Timeout.Nanoseconds(), &sensorDetailResp); err != nil {
		return "", err
	}
	return sensorDetailResp.PrtgVersion, nil
}

// GetSensorDetail ruft die Details eines Sensors ab.
func (c *Client) GetSensorDetail(id int64) (*PrtgSensorData, error) {
	q := c.getTemplateUrlQuery()
	q.Set("id", fmt.Sprintf("%v", id))

	u, err := c.getCompleteUrl(GetSensorDetailsEndpoint, q)
	if err != nil {
		return nil, err
	}

	var sensorDetailResp prtgSensorDetailsResponse
	if err := getPrtgResponse(u, c.Timeout.Nanoseconds(), &sensorDetailResp); err != nil {
		return nil, err
	}
	return &sensorDetailResp.SensorData, nil
}

// GetHistoricData ruft historische Sensordaten ab.
func (c *Client) GetHistoricData(id, average int64, startDate, endDate time.Time) ([]PrtgHistoricData, error) {
	if id < 0 || average < 0 {
		return nil, fmt.Errorf("ID und Durchschnittswert müssen größer oder gleich null sein")
	}
	if delta := endDate.Sub(startDate); delta < 0 || delta > deltaHistoricThreshold {
		return nil, fmt.Errorf("datenbereich überschreitet 31 tage")
	}

	q := c.getTemplateUrlQuery()
	q.Set("id", fmt.Sprintf("%v", id))
	q.Set("avg", fmt.Sprintf("%v", average))
	q.Set("sDate", startDate.Format(dateFormat))
	q.Set("eDate", endDate.Format(dateFormat))
	q.Set("usecaption", "1")

	u, err := c.getCompleteUrl(GetHistoricDatasEndpoint, q)
	if err != nil {
		return nil, err
	}

	var histDataResp prtgHistoricDataResponse
	if err := getPrtgResponse(u, c.Timeout.Nanoseconds(), &histDataResp); err != nil {
		return nil, err
	}
	if len(histDataResp.HistoricData) == 0 {
		return nil, fmt.Errorf("keine daten gefunden")
	}
	return histDataResp.HistoricData, nil
}

// GetSensorTree gibt die Baumstruktur eines Sensors zurück.
func (c *Client) GetSensorTree(id int64) (*PrtgSensorTreeResponse, error) {
	if id < 0 {
		return nil, fmt.Errorf("ID muss größer oder gleich null sein")
	}

	q := c.getTemplateUrlQuery()
	q.Set("id", fmt.Sprintf("%v", id))
	q.Set("content", "sensortree")

	u, err := c.getCompleteUrl(GetSensorTreesEndpoint, q)
	if err != nil {
		return nil, err
	}

	var tableTreeResp PrtgSensorTreeResponse
	if err := getPrtgResponse(u, c.Timeout.Nanoseconds(), &tableTreeResp); err != nil {
		return nil, err
	}
	return &tableTreeResp, nil
}
