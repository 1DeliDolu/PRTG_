import { DataSourceJsonData } from '@grafana/data';
import { DataQuery } from '@grafana/schema';

export enum QueryType {
  Metrics = 'metrics',
  Raw = 'raw',
  Text = 'text'
}

export interface MyQuery extends DataQuery {
  group: string;
  device: string;
  sensor: string;
  channel: string;
  queryType: QueryType;
  property: string;
  filterProperty: string;
  includeGroupName: boolean;
  includeDeviceName: boolean;
  includeSensorName: boolean;
  groups: Array<string>;
  devices: Array<string>;
  sensors: Array<string>;
  channels: Array<string>;
}

export interface DataPoint {
  Time: number;
  Value: number | string;
}

export interface DataSourceResponse {
  datapoints: DataPoint[];
}

/**
 * These are options configured for each DataSource instance
 */
export interface MyDataSourceOptions extends DataSourceJsonData {
  path?: string;
  cacheTime?: number;
}

export interface MySecureJsonData {
  apiKey?: string;
}

export interface QueryTypeOptions {
  label: string;
  value: QueryType;
}

export const queryTypeOptions = Object.keys(QueryType).map((key) => ({
  label: key,
  value: QueryType[key as keyof typeof QueryType],
}));

export interface ListItem {
  name: string;
  visible_name: string;
}


export interface PRTGItem {
  objid: number;
  objid_raw: number;
  group: string;
  group_raw: string;
  device: string;
  device_raw: string;
  sensor: string;
  sensor_raw: string;
  channel: string;
  channel_raw: string;
  active: boolean;
  active_raw: number;
  message: string;
  message_raw: string;
  priority: string;
  priority_raw: number;
  status: string;
  status_raw: number;
  tags: string;
  tags_raw: string;
  datetime: string;
  datetime_raw: number;
}

export interface PRTGGroupListResponse {
  prtgversion: string;
  treesize: number;
  groups: PRTGItem[];
}

export interface PRTGGroupResponse {
  groups: PRTGItem[];
}

export interface PRTGDeviceListResponse {
  prtgversion: string;
  treesize: number;
  devices: PRTGItem[];
}

export interface PRTGDeviceResponse {
  devices: PRTGItem[];
}

export interface PRTGSensorListResponse {
  prtgversion: string;
  treesize: number;
  sensors: PRTGItem[];
}

export interface PRTGSensorResponse {
  sensors: PRTGItem[];
}

export interface PRTGChannelListResponse {
  prtgversion: string;
  treesize: number;
  histdata: PRTGItemChannel[];
}

export interface PRTGItemChannel {
  channel: string;
  channel_raw: string;
  datetime: string;
  datetime_raw: number;
  [key: string]: string | number; 
}

export const filterPropertyList = [
  { name: 'active', visible_name: 'Active' },
  { name: 'message_raw', visible_name: 'Message' },
  { name: 'priority', visible_name: 'Priority' },
  { name: 'status', visible_name: 'Status' },
  { name: 'tags', visible_name: 'Tags' },
] as const;

export type FilterPropertyItem = typeof filterPropertyList[number];

export interface FilterPropertyOption {
  label: string;
  value: FilterPropertyItem['name'];
}

export const propertyList = [
  { name: 'group', visible_name: 'Group' },
  { name: 'device', visible_name: 'Device' },
  { name: 'sensor', visible_name: 'Sensor' },
] as const;

export type PropertyItem = typeof propertyList[number];

export interface PropertyOption {
  label: string;
  value: PropertyItem['name'];
}


//##################################################################################################
export type SensorColumnItem = typeof sensorColumnList[number];
export type GroupColumnItem = typeof groupColumnList[number];
export type DeviceColumnItem = typeof deviceColumnList[number];

export interface ColumnOption {
  label: string;
  value: string;
}

export const sensorColumnList = [
  { name: 'accessrights', visible_name: 'Access Rights' },
  { name: 'active', visible_name: 'Active' },
  { name: 'downtime', visible_name: 'Downtime Percentage' },
  { name: 'downtimesince', visible_name: 'Elapsed Time Since Last Up' },
  { name: 'downtimetime', visible_name: 'Downtime Duration' },
  { name: 'downsens', visible_name: 'Down Sensors Count' },
  { name: 'interval', visible_name: 'Effective Interval' },
  { name: 'lastcheck', visible_name: 'Last Check Time' },
  { name: 'lastdown', visible_name: 'Last Down Time' },
  { name: 'lastup', visible_name: 'Last Up Time' },
  { name: 'message', visible_name: 'Detailed Message' },
  { name: 'pausedsens', visible_name: 'Paused Sensors Count' },
  { name: 'priority', visible_name: 'Priority' },
  { name: 'status', visible_name: 'Status' },
  { name: 'tags', visible_name: 'Tags' },
  { name: 'totalsens', visible_name: 'Total Sensors Count' },
  { name: 'unusualsens', visible_name: 'Unusual Sensors Count' },
  { name: 'uptime', visible_name: 'Uptime Percentage' },
  { name: 'uptimesince', visible_name: 'Elapsed Time Since Last Down' },
  { name: 'uptimetime', visible_name: 'Uptime Duration' },
  { name: 'upsens', visible_name: 'Up Sensors Count' },
  { name: 'warnsens', visible_name: 'Warning Sensors Count' },
] as const;

export const groupColumnList = [
  { name: 'accessrights', visible_name: 'Access Rights' },
  { name: 'active', visible_name: 'Active' },
  { name: 'downsens', visible_name: 'Down Sensors Count' },
  { name: 'message', visible_name: 'Detailed Message' },
  { name: 'pausedsens', visible_name: 'Paused Sensors Count' },
  { name: 'priority', visible_name: 'Priority' },
  { name: 'status', visible_name: 'Status' },
  { name: 'tags', visible_name: 'Tags' },
  { name: 'totalsens', visible_name: 'Total Sensors Count' },
  { name: 'unusualsens', visible_name: 'Unusual Sensors Count' },
  { name: 'upsens', visible_name: 'Up Sensors Count' },
  { name: 'warnsens', visible_name: 'Warning Sensors Count' },
] as const;

export const deviceColumnList = [
  { name: 'accessrights', visible_name: 'Access Rights' },
  { name: 'active', visible_name: 'Active' },
  { name: 'deviceicon', visible_name: 'Device Icon' },
  { name: 'downsens', visible_name: 'Down Sensors Count' },
  { name: 'location', visible_name: 'Location' },
  { name: 'message', visible_name: 'Detailed Message' },
  { name: 'pausedsens', visible_name: 'Paused Sensors Count' },
  { name: 'priority', visible_name: 'Priority' },
  { name: 'status', visible_name: 'Status' },
  { name: 'tags', visible_name: 'Tags' },
  { name: 'totalsens', visible_name: 'Total Sensors Count' },
  { name: 'unusualsens', visible_name: 'Unusual Sensors Count' },
  { name: 'upsens', visible_name: 'Up Sensors Count' },
  { name: 'warnsens', visible_name: 'Warning Sensors Count' }
] as const;
