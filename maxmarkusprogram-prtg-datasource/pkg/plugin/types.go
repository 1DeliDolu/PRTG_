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
    AccessRights   string  `json:"accessrights" xml:"accessrights"`
    Active         bool    `json:"active" xml:"active"`
    ActiveRAW      int     `json:"active_raw" xml:"active_raw"`
    Channel        string  `json:"channel" xml:"channel"`
    ChannelRAW     string  `json:"channel_raw" xml:"channel_raw"`
    Datetime       string  `json:"datetime" xml:"datetime"`
    DatetimeRAW    float64 `json:"datetime_raw" xml:"datetime_raw"`
    Device         string  `json:"device" xml:"device"`
    DeviceRAW      string  `json:"device_raw" xml:"device_raw"`
    Downsens       int     `json:"downsens" xml:"downsens"`
    Group          string  `json:"group" xml:"group"`
    GroupRAW       string  `json:"group_raw" xml:"group_raw"`
    Message        string  `json:"message" xml:"message"`
    MessageRAW     string  `json:"message_raw" xml:"message_raw"`
    ObjectId       int64   `json:"objid" xml:"objid"`
    ObjectIdRAW    int64   `json:"objid_raw" xml:"objid_raw"`
    Pausedsens     int     `json:"pausedsens" xml:"pausedsens"`
    Priority       string  `json:"priority" xml:"priority"`
    PriorityRAW    int     `json:"priority_raw" xml:"priority_raw"`
    Sensor         string  `json:"sensor" xml:"sensor"`
    SensorRAW      string  `json:"sensor_raw" xml:"sensor_raw"`
    Status         string  `json:"status" xml:"status"`
    StatusRAW      int     `json:"status_raw" xml:"status_raw"`
    Tags           string  `json:"tags" xml:"tags"`
    TagsRAW        string  `json:"tags_raw" xml:"tags_raw"`
    Totalsens      int     `json:"totalsens" xml:"totalsens"`
    Unusualsens    int     `json:"unusualsens" xml:"unusualsens"`
    Upsens         int     `json:"upsens" xml:"upsens"`
    Warnsens       int     `json:"warnsens" xml:"warnsens"`
}

type PrtgDevicesListResponse struct {
    PrtgVersion string         `json:"prtgversion" xml:"prtg-version"`
    TreeSize    int64          `json:"treesize" xml:"treesize"`
    DeviceList  []PrtgDeviceListItemStruct `json:"devices" xml:"devices"`
}
type PrtgDeviceListItemStruct struct {
    AccessRights   string  `json:"accessrights" xml:"accessrights"`
    Active         bool    `json:"active" xml:"active"`
    ActiveRAW      int     `json:"active_raw" xml:"active_raw"`
    Channel        string  `json:"channel" xml:"channel"`
    ChannelRAW     string  `json:"channel_raw" xml:"channel_raw"`
    Datetime       string  `json:"datetime" xml:"datetime"`
    DatetimeRAW    float64 `json:"datetime_raw" xml:"datetime_raw"`
    Device         string  `json:"device" xml:"device"`
    DeviceIcon     string  `json:"deviceicon" xml:"deviceicon"`
    DeviceIconRAW  string  `json:"deviceicon_raw" xml:"deviceicon_raw"`
    DeviceRAW      string  `json:"device_raw" xml:"device_raw"`
    Downsens       int     `json:"downsens" xml:"downsens"`
    Group          string  `json:"group" xml:"group"`
    GroupRAW       string  `json:"group_raw" xml:"group_raw"`
    Location       string  `json:"location" xml:"location"`
    LocationRAW    string  `json:"location_raw" xml:"location_raw"`
    Message        string  `json:"message" xml:"message"`
    MessageRAW     string  `json:"message_raw" xml:"message_raw"`
    ObjectId       int64   `json:"objid" xml:"objid"`
    ObjectIdRAW    int64   `json:"objid_raw" xml:"objid_raw"`
    PausedSens     int     `json:"pausedsens" xml:"pausedsens"`
    Priority       string  `json:"priority" xml:"priority"`
    PriorityRAW    int     `json:"priority_raw" xml:"priority_raw"`
    Sensor         string  `json:"sensor" xml:"sensor"`
    SensorRAW      string  `json:"sensor_raw" xml:"sensor_raw"`
    Status         string  `json:"status" xml:"status"`
    StatusRAW      int     `json:"status_raw" xml:"status_raw"`
    Tags           string  `json:"tags" xml:"tags"`
    TagsRAW        string  `json:"tags_raw" xml:"tags_raw"`
    TotalSens      int     `json:"totalsens" xml:"totalsens"`
    UnusualSens    int     `json:"unusualsens" xml:"unusualsens"`
    UpSens         int     `json:"upsens" xml:"upsens"`
    WarnSens       int     `json:"warnsens" xml:"warnsens"`
}

type PrtgSensorsListResponse struct {
    PrtgVersion string         `json:"prtgversion" xml:"prtg-version"`
    TreeSize    int64          `json:"treesize" xml:"treesize"`
    SensorList  []PrtgSensorListItemStruct `json:"sensors" xml:"sensors"`
}

type PrtgSensorListItemStruct struct {
    AccessRights         string  `json:"accessrights" xml:"accessrights"`
    AccessRightsRAW      string  `json:"accessrights_raw" xml:"accessrights_raw"`
    Active              bool    `json:"active" xml:"active"`
    ActiveRAW           int     `json:"active_raw" xml:"active_raw"`
    Channel             string  `json:"channel" xml:"channel"`
    ChannelRAW          int     `json:"channel_raw" xml:"channel_raw"`
    Datetime            string  `json:"datetime" xml:"datetime"`
    DatetimeRAW         float64 `json:"datetime_raw" xml:"datetime_raw"`
    Device              string  `json:"device" xml:"device"`
    DeviceRAW           string  `json:"device_raw" xml:"device_raw"`
    Downtime            float64 `json:"downtime" xml:"downtime"`
    DowntimeRAW         float64 `json:"downtime_raw" xml:"downtime_raw"`
    DowntimeSince       string  `json:"downtimesince" xml:"downtimesince"`
    DowntimeSinceRAW    float64 `json:"downtimesince_raw" xml:"downtimesince_raw"`
    DowntimeTime        string  `json:"downtimetime" xml:"downtimetime"`
    DowntimeTimeRAW     float64 `json:"downtimetime_raw" xml:"downtimetime_raw"`
    Downsens            int     `json:"downsens" xml:"downsens"`
    DownsensRAW         int     `json:"downsens_raw" xml:"downsens_raw"`
    Group               string  `json:"group" xml:"group"`
    GroupRAW            string  `json:"group_raw" xml:"group_raw"`
    Interval            string  `json:"interval" xml:"interval"`
    IntervalRAW         int     `json:"interval_raw" xml:"interval_raw"`
    LastCheck           string  `json:"lastcheck" xml:"lastcheck"`
    LastCheckRAW        float64 `json:"lastcheck_raw" xml:"lastcheck_raw"`
    LastDown            string  `json:"lastdown" xml:"lastdown"`
    LastDownRAW         float64 `json:"lastdown_raw" xml:"lastdown_raw"`
    LastUp              string  `json:"lastup" xml:"lastup"`
    LastUpRAW           float64 `json:"lastup_raw" xml:"lastup_raw"`
    Message             string  `json:"message" xml:"message"`
    MessageRAW          string  `json:"message_raw" xml:"message_raw"`
    ObjectId            int64   `json:"objid" xml:"objid"`
    ObjectIdRAW         int64   `json:"objid_raw" xml:"objid_raw"`
    ParentId            int64   `json:"parentid" xml:"parentid"`
    ParentIdRAW         int64   `json:"parentid_raw" xml:"parentid_raw"`
    PausedSens          int     `json:"pausedsens" xml:"pausedsens"`
    PausedSensRAW       int     `json:"pausedsens_raw" xml:"pausedsens_raw"`
    Priority            string  `json:"priority" xml:"priority"`
    PriorityRAW         int     `json:"priority_raw" xml:"priority_raw"`
    Sensor              string  `json:"sensor" xml:"sensor"`
    SensorRAW           string  `json:"sensor_raw" xml:"sensor_raw"`
    Status              string  `json:"status" xml:"status"`
    StatusRAW           int     `json:"status_raw" xml:"status_raw"`
    Tags                string  `json:"tags" xml:"tags"`
    TagsRAW             string  `json:"tags_raw" xml:"tags_raw"`
    TotalSens           int     `json:"totalsens" xml:"totalsens"`
    TotalSensRAW        int     `json:"totalsens_raw" xml:"totalsens_raw"`
    UnusualSens         int     `json:"unusualsens" xml:"unusualsens"`
    UnusualSensRAW      int     `json:"unusualsens_raw" xml:"unusualsens_raw"`
    Uptime              float64 `json:"uptime" xml:"uptime"`
    UptimeRAW           float64 `json:"uptime_raw" xml:"uptime_raw"`
    UptimeSince         string  `json:"uptimesince" xml:"uptimesince"`
    UptimeSinceRAW      float64 `json:"uptimesince_raw" xml:"uptimesince_raw"`
    UptimeTime          string  `json:"uptimetime" xml:"uptimetime"`
    UptimeTimeRAW       float64 `json:"uptimetime_raw" xml:"uptimetime_raw"`
    UpSens              int     `json:"upsens" xml:"upsens"`
    UpSensRAW           int     `json:"upsens_raw" xml:"upsens_raw"`
    WarnSens            int     `json:"warnsens" xml:"warnsens"`
    WarnSensRAW         int     `json:"warnsens_raw" xml:"warnsens_raw"`
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
// channel list  PrtgChannelsListResponse
type PrtgChannelsListResponse struct {
    DatetimeRAW float64                `json:"datetime_raw" xml:"datetime_raw"`
    ValueRAW    PrtgChannelValueStruct `json:"value_raw" xml:"value_raw"`
    StatusRAW   int                    `json:"status_raw,omitempty" xml:"status_raw,omitempty"`
}

type PrtgChannelValueStruct struct {
    Text    string `json:"text" xml:"text"`
    Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
}

