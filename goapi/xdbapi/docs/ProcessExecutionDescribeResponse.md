# ProcessExecutionDescribeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessExecutionId** | Pointer to **string** |  | [optional] 
**ProcessType** | Pointer to **string** | the process type for SDK to lookup the process definition class | [optional] 
**WorkerUrl** | Pointer to **string** | the URL for XDB async service to make callback to worker | [optional] 
**StartTimestamp** | Pointer to **int32** | start time of the process execution | [optional] 
**Status** | Pointer to [**ProcessStatus**](ProcessStatus.md) |  | [optional] 

## Methods

### NewProcessExecutionDescribeResponse

`func NewProcessExecutionDescribeResponse() *ProcessExecutionDescribeResponse`

NewProcessExecutionDescribeResponse instantiates a new ProcessExecutionDescribeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessExecutionDescribeResponseWithDefaults

`func NewProcessExecutionDescribeResponseWithDefaults() *ProcessExecutionDescribeResponse`

NewProcessExecutionDescribeResponseWithDefaults instantiates a new ProcessExecutionDescribeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessExecutionId

`func (o *ProcessExecutionDescribeResponse) GetProcessExecutionId() string`

GetProcessExecutionId returns the ProcessExecutionId field if non-nil, zero value otherwise.

### GetProcessExecutionIdOk

`func (o *ProcessExecutionDescribeResponse) GetProcessExecutionIdOk() (*string, bool)`

GetProcessExecutionIdOk returns a tuple with the ProcessExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessExecutionId

`func (o *ProcessExecutionDescribeResponse) SetProcessExecutionId(v string)`

SetProcessExecutionId sets ProcessExecutionId field to given value.

### HasProcessExecutionId

`func (o *ProcessExecutionDescribeResponse) HasProcessExecutionId() bool`

HasProcessExecutionId returns a boolean if a field has been set.

### GetProcessType

`func (o *ProcessExecutionDescribeResponse) GetProcessType() string`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *ProcessExecutionDescribeResponse) GetProcessTypeOk() (*string, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *ProcessExecutionDescribeResponse) SetProcessType(v string)`

SetProcessType sets ProcessType field to given value.

### HasProcessType

`func (o *ProcessExecutionDescribeResponse) HasProcessType() bool`

HasProcessType returns a boolean if a field has been set.

### GetWorkerUrl

`func (o *ProcessExecutionDescribeResponse) GetWorkerUrl() string`

GetWorkerUrl returns the WorkerUrl field if non-nil, zero value otherwise.

### GetWorkerUrlOk

`func (o *ProcessExecutionDescribeResponse) GetWorkerUrlOk() (*string, bool)`

GetWorkerUrlOk returns a tuple with the WorkerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerUrl

`func (o *ProcessExecutionDescribeResponse) SetWorkerUrl(v string)`

SetWorkerUrl sets WorkerUrl field to given value.

### HasWorkerUrl

`func (o *ProcessExecutionDescribeResponse) HasWorkerUrl() bool`

HasWorkerUrl returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *ProcessExecutionDescribeResponse) GetStartTimestamp() int32`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *ProcessExecutionDescribeResponse) GetStartTimestampOk() (*int32, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *ProcessExecutionDescribeResponse) SetStartTimestamp(v int32)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *ProcessExecutionDescribeResponse) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.

### GetStatus

`func (o *ProcessExecutionDescribeResponse) GetStatus() ProcessStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProcessExecutionDescribeResponse) GetStatusOk() (*ProcessStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProcessExecutionDescribeResponse) SetStatus(v ProcessStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProcessExecutionDescribeResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


