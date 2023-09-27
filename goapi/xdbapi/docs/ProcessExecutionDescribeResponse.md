# ProcessExecutionDescribeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessExecutionId** | Pointer to **string** |  | [optional] 
**ProcessType** | Pointer to **string** | the process type for SDK to lookup the process definition class | [optional] 
**WorkerUrl** | Pointer to **string** | the URL for XDB async service to make callback to worker | [optional] 
**StartStateId** | Pointer to **string** | StateId of the first AsyncState to start | [optional] 
**StartStateConfig** | Pointer to [**AsyncStateConfig**](AsyncStateConfig.md) |  | [optional] 
**ProcessStartConfig** | Pointer to [**ProcessStartConfig**](ProcessStartConfig.md) |  | [optional] 

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

### GetStartStateId

`func (o *ProcessExecutionDescribeResponse) GetStartStateId() string`

GetStartStateId returns the StartStateId field if non-nil, zero value otherwise.

### GetStartStateIdOk

`func (o *ProcessExecutionDescribeResponse) GetStartStateIdOk() (*string, bool)`

GetStartStateIdOk returns a tuple with the StartStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartStateId

`func (o *ProcessExecutionDescribeResponse) SetStartStateId(v string)`

SetStartStateId sets StartStateId field to given value.

### HasStartStateId

`func (o *ProcessExecutionDescribeResponse) HasStartStateId() bool`

HasStartStateId returns a boolean if a field has been set.

### GetStartStateConfig

`func (o *ProcessExecutionDescribeResponse) GetStartStateConfig() AsyncStateConfig`

GetStartStateConfig returns the StartStateConfig field if non-nil, zero value otherwise.

### GetStartStateConfigOk

`func (o *ProcessExecutionDescribeResponse) GetStartStateConfigOk() (*AsyncStateConfig, bool)`

GetStartStateConfigOk returns a tuple with the StartStateConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartStateConfig

`func (o *ProcessExecutionDescribeResponse) SetStartStateConfig(v AsyncStateConfig)`

SetStartStateConfig sets StartStateConfig field to given value.

### HasStartStateConfig

`func (o *ProcessExecutionDescribeResponse) HasStartStateConfig() bool`

HasStartStateConfig returns a boolean if a field has been set.

### GetProcessStartConfig

`func (o *ProcessExecutionDescribeResponse) GetProcessStartConfig() ProcessStartConfig`

GetProcessStartConfig returns the ProcessStartConfig field if non-nil, zero value otherwise.

### GetProcessStartConfigOk

`func (o *ProcessExecutionDescribeResponse) GetProcessStartConfigOk() (*ProcessStartConfig, bool)`

GetProcessStartConfigOk returns a tuple with the ProcessStartConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessStartConfig

`func (o *ProcessExecutionDescribeResponse) SetProcessStartConfig(v ProcessStartConfig)`

SetProcessStartConfig sets ProcessStartConfig field to given value.

### HasProcessStartConfig

`func (o *ProcessExecutionDescribeResponse) HasProcessStartConfig() bool`

HasProcessStartConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


