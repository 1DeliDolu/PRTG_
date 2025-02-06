package plugin

// PrtgTableListResponse represents the PRTG Table List API response
type PrtgTableListResponse struct {
    PrtgVersion []PrtgStatusListResponse   `json:"prtgversion" xml:"prtg-version"`
    TreeSize    int64                      `json:"treesize" xml:"treesize"`
    Groups      []PrtgGroupListResponse    `json:"groups,omitempty" xml:"groups,omitempty"`
    Devices     []PrtgDevicesListResponse  `json:"devices,omitempty" xml:"devices,omitempty"`
    Sensors     []PrtgSensorsListResponse  `json:"sensors,omitempty" xml:"sensors,omitempty"`
}

type PrtgGroupListResponse struct {
    PrtgVersion string                    `json:"prtgversion" xml:"prtg-version"`
    TreeSize    int64                     `json:"treesize" xml:"treesize"`
    GroupList   []PrtgGroupListItemStruct `json:"groups" xml:"groups"`
}

type PrtgGroupListItemStruct struct {
    ObjectId     int64   `json:"objid" xml:"objid"`
    ObjectIdRAW  int64   `json:"objid_raw" xml:"objid_raw"`
    Group        string  `json:"group" xml:"group"`
    GroupRAW     string  `json:"group_raw" xml:"group_raw"`
    Device       string  `json:"device" xml:"device"`
    DeviceRAW    string  `json:"device_raw" xml:"device_raw"`
    Sensor       string  `json:"sensor" xml:"sensor"`
    SensorRAW    string  `json:"sensor_raw" xml:"sensor_raw"`
    Channel      string  `json:"channel" xml:"channel"`
    ChannelRAW   string  `json:"channel_raw" xml:"channel_raw"`
    Active       bool    `json:"active" xml:"active"`
    ActiveRAW    int     `json:"active_raw" xml:"active_raw"`
    Message      string  `json:"message" xml:"message"`
    MessageRAW   string  `json:"message_raw" xml:"message_raw"`
    Priority     string  `json:"priority" xml:"priority"`
    PriorityRAW  int     `json:"priority_raw" xml:"priority_raw"`
    Status       string  `json:"status" xml:"status"`
    StatusRAW    int     `json:"status_raw" xml:"status_raw"`
    Tags         string  `json:"tags" xml:"tags"`
    TagsRAW      string  `json:"tags_raw" xml:"tags_raw"`
    Datetime     string  `json:"datetime" xml:"datetime"`
    DatetimeRAW  float64 `json:"datetime_raw" xml:"datetime_raw"`
}

type PrtgDevicesListResponse struct {
    PrtgVersion string         `json:"prtgversion" xml:"prtg-version"`
    TreeSize    int64          `json:"treesize" xml:"treesize"`
    DeviceList  []PrtgGroupListItemStruct `json:"devices" xml:"devices"`
}

type PrtgSensorsListResponse struct {
    PrtgVersion string         `json:"prtgversion" xml:"prtg-version"`
    TreeSize    int64          `json:"treesize" xml:"treesize"`
    SensorList  []PrtgSensorListItemStruct `json:"sensors" xml:"sensors"`
}

type PrtgSensorListItemStruct struct {
    ObjectId     int64   `json:"objid" xml:"objid"`
    ObjectIdRAW  int64   `json:"objid_raw" xml:"objid_raw"`
    Group        string  `json:"group" xml:"group"`
    GroupRAW     string  `json:"group_raw" xml:"group_raw"`
    Device       string  `json:"device" xml:"device"`
    DeviceRAW    string  `json:"device_raw" xml:"device_raw"`
    Sensor       string  `json:"sensor" xml:"sensor"`
    SensorRAW    string  `json:"sensor_raw" xml:"sensor_raw"`
    Channel      string  `json:"channel" xml:"channel"`
    ChannelRAW   int     `json:"channel_raw" xml:"channel_raw"`
    Active       bool    `json:"active" xml:"active"`
    ActiveRAW    int     `json:"active_raw" xml:"active_raw"`
    Message      string  `json:"message" xml:"message"`
    MessageRAW   string  `json:"message_raw" xml:"message_raw"`
    Priority     string  `json:"priority" xml:"priority"`
    PriorityRAW  int     `json:"priority_raw" xml:"priority_raw"`
    Status       string  `json:"status" xml:"status"`
    StatusRAW    int     `json:"status_raw" xml:"status_raw"`
    Tags         string  `json:"tags" xml:"tags"`
    TagsRAW      string  `json:"tags_raw" xml:"tags_raw"`
    Datetime     string  `json:"datetime" xml:"datetime"`
    DatetimeRAW  float64 `json:"datetime_raw" xml:"datetime_raw"`
}

type PrtgStatusListResponse struct {
    PrtgVersion          string `json:"prtgversion" xml:"prtg-version"`
    AckAlarms            string `json:"ackalarms" xml:"ackalarms"`
    Alarms               string `json:"alarms" xml:"alarms"`
    AutoDiscoTasks       string `json:"autodiscotasks" xml:"autodiscotasks"`
    BackgroundTasks      string `json:"backgroundtasks" xml:"backgroundtasks"`
    Clock                string `json:"clock" xml:"clock"`
    ClusterNodeName      string `json:"clusternodename" xml:"clusternodename"`
    ClusterType          string `json:"clustertype" xml:"clustertype"`
    CommercialExpiryDays int    `json:"commercialexpirydays" xml:"commercialexpirydays"`
    CorrelationTasks     string `json:"correlationtasks" xml:"correlationtasks"`
    DaysInstalled        int    `json:"daysinstalled" xml:"daysinstalled"`
    EditionType          string `json:"editiontype" xml:"editiontype"`
    Favs                 int    `json:"favs" xml:"favs"`
    JsClock              int64  `json:"jsclock" xml:"jsclock"`
    LowMem               bool   `json:"lowmem" xml:"lowmem"`
    MaintExpiryDays      string `json:"maintexpirydays" xml:"maintexpirydays"`
    MaxSensorCount       string `json:"maxsensorcount" xml:"maxsensorcount"`
    NewAlarms            string `json:"newalarms" xml:"newalarms"`
    NewMessages          string `json:"newmessages" xml:"newmessages"`
    NewTickets           string `json:"newtickets" xml:"newtickets"`
    Overloadprotection   bool   `json:"overloadprotection" xml:"overloadprotection"`
    PartialAlarms        string `json:"partialalarms" xml:"partialalarms"`
    PausedSens           string `json:"pausedsens" xml:"pausedsens"`
    PRTGUpdateAvailable  bool   `json:"prtgupdateavailable" xml:"prtgupdateavailable"`
    ReadOnlyUser         string `json:"readonlyuser" xml:"readonlyuser"`
    ReportTasks          string `json:"reporttasks" xml:"reporttasks"`
    TotalSens            int    `json:"totalsens" xml:"totalsens"`
    TrialExpiryDays      int    `json:"trialexpirydays" xml:"trialexpirydays"`
    UnknownSens          string `json:"unknownsens" xml:"unknownsens"`
    UnusualSens          string `json:"unusualsens" xml:"unusualsens"`
    UpSens               string `json:"upsens" xml:"upsens"`
    Version              string `json:"version" xml:"version"`
    WarnSens             string `json:"warnsens" xml:"warnsens"`
}



