package plugin

// prtgSensorDetailsResponse stellt die Antwort der PRTG Sensor Details API dar.
type prtgSensorDetailsResponse struct {
	PrtgVersion string         `json:"prtgversion" xml:"prtg-version"`
	SensorData  PrtgSensorData `json:"sensordata" xml:"sensordata"`
}

type prtgSensorDetailsResponseXML struct {
	PrtgVersion string `xml:"prtg-version"`
	PrtgSensorData
}

// PrtgSensorData enthält Eigenschaften für jeden Sensor, jedes Gerät und jede Gruppe.
type PrtgSensorData struct {
	Name             string `json:"name" xml:"name"`
	SensorType       string `json:"sensortype" xml:"sensortype"`
	Interval         string `json:"interval" xml:"interval"`
	ProbeName        string `json:"probename" xml:"probename"`
	ParentGroupName  string `json:"parentgroupname" xml:"parentgroupname"`
	ParentDeviceName string `json:"parentdevicename" xml:"parentdevicename"`
	ParentDeviceId   string `json:"parentdeviceid" xml:"parentdeviceid"`
	LastValue        string `json:"lastvalue" xml:"lastvalue"`
	LastMessage      string `json:"lastmessage" xml:"lastmessage"`
	Favorite         string `json:"favorite" xml:"favorite"`
	StatusText       string `json:"statustext" xml:"statustext"`
	StatusId         string `json:"statusid" xml:"statusid"`
	LastUp           string `json:"lastup" xml:"lastup"`
	LastDown         string `json:"lastdown" xml:"lastdown"`
	LastCheck        string `json:"lastcheck" xml:"lastcheck"`
	Uptime           string `json:"uptime" xml:"uptime"`
	UptimeTime       string `json:"uptimetime" xml:"uptimetime"`
	Downtime         string `json:"downtime" xml:"downtime"`
	DowntimeTime     string `json:"downtimetime" xml:"downtimetime"`
	UpDownTotal      string `json:"updowntotal" xml:"updowntotal"`
	UpDownSince      string `json:"updownsince" xml:"updownsince"`
	Info             string `json:"info" xml:"info"`
}

// prtgTableListResponse stellt die Antwort der PRTG Tabellenlisten API dar.
type prtgTableListResponse struct {
	PrtgVersion string          `json:"prtgversion" xml:"prtg-version"`
	TreeSize    int64           `json:"treesize" xml:"treesize"`
	Groups      []PrtgTableList `json:"groups,omitempty" xml:"groups,omitempty"`
	Devices     []PrtgTableList `json:"devices,omitempty" xml:"devices,omitempty"`
	Sensors     []PrtgTableList `json:"sensors,omitempty" xml:"sensors,omitempty"`
}

type PrtgTableList struct {
	ObjectId int64  `json:"objid" xml:"objid"`
	Probe    string `json:"probe" xml:"probe"`
	Group    string `json:"group" xml:"group"`
	Name     string `json:"name" xml:"name"`
	Device   string `json:"device" xml:"device"`
	Host     string `json:"host" xml:"host"`
	Sensor   string `json:"sensor" xml:"sensor"`
}

type prtgHistoricDataResponse struct {
	PrtgVersion  string             `json:"prtgversion" xml:"prtg-version"`
	TreeSize     int64              `json:"treesize" xml:"treesize"`
	HistoricData []PrtgHistoricData `json:"histdata" xml:"histdata"`
}

type PrtgHistoricData map[string]interface{}

type prtgHistoricDataResponseXML struct {
	PrtgVersion  string       `json:"prtgversion" xml:"prtg-version"`
	HistoricData []ItemTagXML `json:"histdata" xml:"item"`
}

type ItemTagXML struct {
	Datetime    string     `xml:"datetime"`
	DatetimeRAW string     `xml:"datetime_raw"`
	Coverage    string     `xml:"coverage"`
	CoverageRAW string     `xml:"coverage_raw"`
	Value       []ValueXML `xml:"value"`
	ValueRAW    []ValueXML `xml:"value_raw"`
}

type ValueXML struct {
	Key   string `xml:"channel,attr"`
	Value string `xml:",chardata"`
}
// ok