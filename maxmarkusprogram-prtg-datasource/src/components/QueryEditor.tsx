import React, { ChangeEvent } from 'react';
import { InlineField, Input, Select, Stack} from '@grafana/ui';
import { QueryEditorProps , SelectableValue} from '@grafana/data';
import { DataSource } from '../datasource';
import { MyDataSourceOptions, MyQuery, queryTypeOptions ,QueryType} from '../types';

type Props = QueryEditorProps<DataSource, MyQuery, MyDataSourceOptions>;

export function QueryEditor({ query, onChange, onRunQuery }: Props) {
  
  
  const onQueryTypeChange = (value: SelectableValue<QueryType>) => {
    onChange({ ...query, queryType: value.value! });
    onRunQuery();
  };

  const onGroupChange = (event: ChangeEvent<HTMLInputElement>) => {
    onChange({ ...query, group: event.target.value });
    onRunQuery();
  };

  const onDeviceChange = (event: ChangeEvent<HTMLInputElement>) => {
    onChange({ ...query, device: event.target.value });
    onRunQuery();
  };

  const onSensorChange = (event: ChangeEvent<HTMLInputElement>) => {
    onChange({ ...query, sensor: event.target.value });
    onRunQuery();
  };

  const onChannelChange = (event: ChangeEvent<HTMLInputElement>) => {
    onChange({ ...query, channel: event.target.value });
    onRunQuery();
  };

  return (
    <div style={{ display: 'flex', gap: '32px', padding: '8px' }}>
      {/* Left Column */}
      <div style={{ flex: '1 1 50%' }}>
        <Stack gap={2}>
          <InlineField label="Query Type" labelWidth={16} grow>
            <Select
              options={queryTypeOptions}
              value={query.queryType}
              onChange={onQueryTypeChange}
              width={32}
            />
          </InlineField>
          <InlineField label="Group" labelWidth={16} grow>
            <Input
              id="query-editor-group"
              onChange={onGroupChange}
              value={query.group || ''}
              placeholder="Select a group"
              width={32}
            />
          </InlineField>
          <InlineField label="Device" labelWidth={16} grow>
            <Input
              id="query-editor-device"
              onChange={onDeviceChange}
              value={query.device || ''}
              placeholder="Select a device"
              width={32}
            />
          </InlineField>
        </Stack>
      </div>
      {/* Right Column */}
      <div style={{ flex: '1 1 50%' }}>
        <Stack gap={2}>
          <InlineField label="Sensor" labelWidth={16} grow>
            <Input
              id="query-editor-sensor"
              onChange={onSensorChange}
              value={query.sensor || ''}
              placeholder="Select a sensor"
              width={32}
            />
          </InlineField>
          <InlineField label="Channel" labelWidth={16} grow>
            <Input
              id="query-editor-channel"
              onChange={onChannelChange}
              value={query.channel || ''}
              placeholder="Select a channel"
              width={32}
            />
          </InlineField>
        </Stack>
      </div>
    </div>
  );
}
