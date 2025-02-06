import React, { useEffect, useState } from 'react'
import { InlineField, Select, Stack, FieldSet, InlineSwitch } from '@grafana/ui'
import { QueryEditorProps, SelectableValue } from '@grafana/data'
import { DataSource } from '../datasource'
import { MyDataSourceOptions, MyQuery, queryTypeOptions, QueryType} from '../types'

type Props = QueryEditorProps<DataSource, MyQuery, MyDataSourceOptions>

export function QueryEditor({ query, onChange, onRunQuery, datasource }: Props) {
  const isMetricsMode = query.queryType === QueryType.Metrics
  const isRawMode = query.queryType === QueryType.Raw
  const isTextMode = query.queryType === QueryType.Text

  const [lists, setLists] = useState({
    groups: [] as Array<SelectableValue<string>>,
    devices: [] as Array<SelectableValue<string>>,
    sensors: [] as Array<SelectableValue<string>>,
    channels: [] as Array<SelectableValue<string>>,
    values: [] as Array<SelectableValue<string>>,
    properties: [] as Array<SelectableValue<string>>,
    filterProperties: [] as Array<SelectableValue<string>>,
  });

  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    async function fetchGroups() {
      setIsLoading(true);
      try {
        const response = await datasource.getGroups();
        if (response && Array.isArray(response.groups)) {
          const groupOptions = response.groups.map(group => ({
            label: group.group,
            value: group.objid.toString(),
          }));
          setLists(prev => ({
            ...prev,
            groups: groupOptions,
          }));
        } else {
          console.error('Invalid response format:', response);
        }
      } catch (error) {
        console.error('Error fetching groups:', error);
      }
      setIsLoading(false);
    }

    fetchGroups();
  }, [datasource]);

  const onQueryTypeChange = (value: SelectableValue<QueryType>) => {
    onChange({ ...query, queryType: value.value! })
    onRunQuery()
  }

  // Add other onChange handlers for Select components
  const onGroupChange = (value: SelectableValue<string>) => {
    onChange({ ...query, group: value.value! })
    onRunQuery()
  }

  const onDeviceChange = (value: SelectableValue<string>) => {
    onChange({ ...query, device: value.value! })
    onRunQuery()
  }

  const onSensorChange = (value: SelectableValue<string>) => {
    onChange({ ...query, sensor: value.value! })
    onRunQuery()
  }

  const onChannelChange = (value: SelectableValue<string>) => {
    onChange({ ...query, channel: value.value! })
    onRunQuery()
  }

  const onPropertyChange = (value: SelectableValue<string>) => {
    onChange({ ...query, property: value.value! })
    onRunQuery()
  }

  const onFilterPropertyChange = (value: SelectableValue<string>) => {
    onChange({ ...query, filterProperty: value.value! })
    onRunQuery()
  }

  const onIncludeGroupName = (e: React.ChangeEvent<HTMLInputElement>) => {
    onChange({ ...query, includeGroupName: e.currentTarget.checked })
    onRunQuery()
  }

  const onIncludeDevice = (e: React.ChangeEvent<HTMLInputElement>) => {
    onChange({ ...query, includeDeviceName: e.currentTarget.checked })
    onRunQuery()
  }

  const onIncludeSensor = (e: React.ChangeEvent<HTMLInputElement>) => {
    onChange({ ...query, includeSensorName: e.currentTarget.checked })
    onRunQuery()
  }

  

  return (
    <Stack direction="column" gap={1}>
      <Stack direction="row" gap={4}>
        <Stack direction="column" gap={1}>
          <InlineField label="Query Type"
            labelWidth={20} grow>
            <Select options={queryTypeOptions}
              value={query.queryType}
              onChange={onQueryTypeChange}
              width={47} />
          </InlineField>

            <InlineField label="Group" labelWidth={20} grow>
            <Select
              isLoading={isLoading}
              options={lists.groups}
              value={query.group}
              onChange={onGroupChange}
              width={47}
              allowCustomValue
              isClearable
              isDisabled={!query.queryType}
              placeholder="Select Group or type '*'"
            />
            </InlineField>
            <InlineField
            label="Device"
            labelWidth={20} grow>
            <Select
              isLoading={!lists.devices.length}
              options={lists.devices}
              value={query.device}
              onChange={onDeviceChange}
              width={47}
              allowCustomValue
              placeholder="Select Device or type '*'"
              isClearable
              isDisabled={!query.group}
            />
            </InlineField>
        </Stack>
        <Stack direction="column" gap={1}>
          <InlineField
            label="Sensor"
            labelWidth={20} grow>
            <Select
              isLoading={!lists.sensors.length}
              options={lists.sensors}
              value={query.sensor}
              onChange={onSensorChange}
              width={47}
              allowCustomValue
              placeholder="Select Sensor or type '*'"
              isClearable
              isDisabled={!query.device}
            />
          </InlineField>

            <InlineField
            label="Channel"
            labelWidth={20} grow>
            <Select
              isLoading={!lists.channels.length}
              options={lists.channels}
              value={query.channel}
              onChange={onChannelChange}
              width={47}
              allowCustomValue
              placeholder="Select Channel or type '*'"
              isClearable
              isDisabled={!query.sensor}
            />
            </InlineField>
        </Stack>
      </Stack>

      {isMetricsMode && (
        <FieldSet
          label="Options">
          <Stack
            direction="row"
            gap={1}>
            <InlineField
              label="Include Group"
              labelWidth={16}>
              <InlineSwitch
                value={query.includeGroupName || false}
                onChange={onIncludeGroupName} />
            </InlineField>

            <InlineField label="Include Device"
              labelWidth={15}>
              <InlineSwitch
                value={query.includeDeviceName || false}
                onChange={onIncludeDevice} />
            </InlineField>

            <InlineField
              label="Include Sensor"
              labelWidth={15}>
              <InlineSwitch
                value={query.includeSensorName || false}
                onChange={onIncludeSensor} />
            </InlineField>
          </Stack>
        </FieldSet>
      )}

      {isTextMode && isRawMode && (
        <FieldSet label="Options">
          <Stack direction="row" gap={1}>
            <InlineField label="Property"
              labelWidth={16}>
              <Select
                options={lists.properties}
                value={query.property}
                onChange={onPropertyChange}
                width={32} />
            </InlineField>
            <InlineField
              label="Filter Property"
              labelWidth={16}>
              <Select
                options={lists.filterProperties}
                value={query.filterProperty}
                onChange={onFilterPropertyChange}
                width={32}
              />
            </InlineField>
          </Stack>
        </FieldSet>
      )}
    </Stack>
  )
}
