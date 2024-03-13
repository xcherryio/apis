# ProcessExecutionListInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** |  | [optional] 
**ProcessId** | Pointer to **string** |  | [optional] 
**ProcessExecutionId** | Pointer to **string** |  | [optional] 
**ProcessType** | Pointer to **string** |  | [optional] 
**StartTimestamp** | Pointer to **int64** |  | [optional] 
**CloseTimestamp** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to [**ProcessStatus**](ProcessStatus.md) |  | [optional] 

## Methods

### NewProcessExecutionListInfo

`func NewProcessExecutionListInfo() *ProcessExecutionListInfo`

NewProcessExecutionListInfo instantiates a new ProcessExecutionListInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessExecutionListInfoWithDefaults

`func NewProcessExecutionListInfoWithDefaults() *ProcessExecutionListInfo`

NewProcessExecutionListInfoWithDefaults instantiates a new ProcessExecutionListInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *ProcessExecutionListInfo) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ProcessExecutionListInfo) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ProcessExecutionListInfo) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ProcessExecutionListInfo) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetProcessId

`func (o *ProcessExecutionListInfo) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *ProcessExecutionListInfo) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *ProcessExecutionListInfo) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *ProcessExecutionListInfo) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetProcessExecutionId

`func (o *ProcessExecutionListInfo) GetProcessExecutionId() string`

GetProcessExecutionId returns the ProcessExecutionId field if non-nil, zero value otherwise.

### GetProcessExecutionIdOk

`func (o *ProcessExecutionListInfo) GetProcessExecutionIdOk() (*string, bool)`

GetProcessExecutionIdOk returns a tuple with the ProcessExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessExecutionId

`func (o *ProcessExecutionListInfo) SetProcessExecutionId(v string)`

SetProcessExecutionId sets ProcessExecutionId field to given value.

### HasProcessExecutionId

`func (o *ProcessExecutionListInfo) HasProcessExecutionId() bool`

HasProcessExecutionId returns a boolean if a field has been set.

### GetProcessType

`func (o *ProcessExecutionListInfo) GetProcessType() string`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *ProcessExecutionListInfo) GetProcessTypeOk() (*string, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *ProcessExecutionListInfo) SetProcessType(v string)`

SetProcessType sets ProcessType field to given value.

### HasProcessType

`func (o *ProcessExecutionListInfo) HasProcessType() bool`

HasProcessType returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *ProcessExecutionListInfo) GetStartTimestamp() int64`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *ProcessExecutionListInfo) GetStartTimestampOk() (*int64, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *ProcessExecutionListInfo) SetStartTimestamp(v int64)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *ProcessExecutionListInfo) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.

### GetCloseTimestamp

`func (o *ProcessExecutionListInfo) GetCloseTimestamp() int64`

GetCloseTimestamp returns the CloseTimestamp field if non-nil, zero value otherwise.

### GetCloseTimestampOk

`func (o *ProcessExecutionListInfo) GetCloseTimestampOk() (*int64, bool)`

GetCloseTimestampOk returns a tuple with the CloseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTimestamp

`func (o *ProcessExecutionListInfo) SetCloseTimestamp(v int64)`

SetCloseTimestamp sets CloseTimestamp field to given value.

### HasCloseTimestamp

`func (o *ProcessExecutionListInfo) HasCloseTimestamp() bool`

HasCloseTimestamp returns a boolean if a field has been set.

### GetStatus

`func (o *ProcessExecutionListInfo) GetStatus() ProcessStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProcessExecutionListInfo) GetStatusOk() (*ProcessStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProcessExecutionListInfo) SetStatus(v ProcessStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProcessExecutionListInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


