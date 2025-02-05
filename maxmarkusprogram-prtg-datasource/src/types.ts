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
}

export const DEFAULT_QUERY: Partial<MyQuery> = {
  group: '',
  device: '',
  sensor: '',
  channel: '',
  queryType: QueryType.Metrics,
};

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
