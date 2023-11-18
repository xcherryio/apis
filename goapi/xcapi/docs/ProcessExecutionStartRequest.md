# ProcessExecutionStartRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** |  | 
**ProcessId** | **string** | the user business identifier for the process, which can be used for multiple ProcessExecution based on ProcessIdReusePolicy | 
**ProcessType** | **string** | the process type for SDK to lookup the process definition class | 
**WorkerUrl** | **string** | the URL for xcherry async service to make callback to worker | 
**StartStateId** | Pointer to **string** | StateId of the first AsyncState to start | [optional] 
**StartStateInput** | Pointer to [**EncodedObject**](EncodedObject.md) |  | [optional] 
**StartStateConfig** | Pointer to [**AsyncStateConfig**](AsyncStateConfig.md) |  | [optional] 
**ProcessStartConfig** | Pointer to [**ProcessStartConfig**](ProcessStartConfig.md) |  | [optional] 

## Methods

### NewProcessExecutionStartRequest

`func NewProcessExecutionStartRequest(namespace string, processId string, processType string, workerUrl string, ) *ProcessExecutionStartRequest`

NewProcessExecutionStartRequest instantiates a new ProcessExecutionStartRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessExecutionStartRequestWithDefaults

`func NewProcessExecutionStartRequestWithDefaults() *ProcessExecutionStartRequest`

NewProcessExecutionStartRequestWithDefaults instantiates a new ProcessExecutionStartRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *ProcessExecutionStartRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ProcessExecutionStartRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ProcessExecutionStartRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetProcessId

`func (o *ProcessExecutionStartRequest) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *ProcessExecutionStartRequest) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *ProcessExecutionStartRequest) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.


### GetProcessType

`func (o *ProcessExecutionStartRequest) GetProcessType() string`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *ProcessExecutionStartRequest) GetProcessTypeOk() (*string, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *ProcessExecutionStartRequest) SetProcessType(v string)`

SetProcessType sets ProcessType field to given value.


### GetWorkerUrl

`func (o *ProcessExecutionStartRequest) GetWorkerUrl() string`

GetWorkerUrl returns the WorkerUrl field if non-nil, zero value otherwise.

### GetWorkerUrlOk

`func (o *ProcessExecutionStartRequest) GetWorkerUrlOk() (*string, bool)`

GetWorkerUrlOk returns a tuple with the WorkerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerUrl

`func (o *ProcessExecutionStartRequest) SetWorkerUrl(v string)`

SetWorkerUrl sets WorkerUrl field to given value.


### GetStartStateId

`func (o *ProcessExecutionStartRequest) GetStartStateId() string`

GetStartStateId returns the StartStateId field if non-nil, zero value otherwise.

### GetStartStateIdOk

`func (o *ProcessExecutionStartRequest) GetStartStateIdOk() (*string, bool)`

GetStartStateIdOk returns a tuple with the StartStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartStateId

`func (o *ProcessExecutionStartRequest) SetStartStateId(v string)`

SetStartStateId sets StartStateId field to given value.

### HasStartStateId

`func (o *ProcessExecutionStartRequest) HasStartStateId() bool`

HasStartStateId returns a boolean if a field has been set.

### GetStartStateInput

`func (o *ProcessExecutionStartRequest) GetStartStateInput() EncodedObject`

GetStartStateInput returns the StartStateInput field if non-nil, zero value otherwise.

### GetStartStateInputOk

`func (o *ProcessExecutionStartRequest) GetStartStateInputOk() (*EncodedObject, bool)`

GetStartStateInputOk returns a tuple with the StartStateInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartStateInput

`func (o *ProcessExecutionStartRequest) SetStartStateInput(v EncodedObject)`

SetStartStateInput sets StartStateInput field to given value.

### HasStartStateInput

`func (o *ProcessExecutionStartRequest) HasStartStateInput() bool`

HasStartStateInput returns a boolean if a field has been set.

### GetStartStateConfig

`func (o *ProcessExecutionStartRequest) GetStartStateConfig() AsyncStateConfig`

GetStartStateConfig returns the StartStateConfig field if non-nil, zero value otherwise.

### GetStartStateConfigOk

`func (o *ProcessExecutionStartRequest) GetStartStateConfigOk() (*AsyncStateConfig, bool)`

GetStartStateConfigOk returns a tuple with the StartStateConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartStateConfig

`func (o *ProcessExecutionStartRequest) SetStartStateConfig(v AsyncStateConfig)`

SetStartStateConfig sets StartStateConfig field to given value.

### HasStartStateConfig

`func (o *ProcessExecutionStartRequest) HasStartStateConfig() bool`

HasStartStateConfig returns a boolean if a field has been set.

### GetProcessStartConfig

`func (o *ProcessExecutionStartRequest) GetProcessStartConfig() ProcessStartConfig`

GetProcessStartConfig returns the ProcessStartConfig field if non-nil, zero value otherwise.

### GetProcessStartConfigOk

`func (o *ProcessExecutionStartRequest) GetProcessStartConfigOk() (*ProcessStartConfig, bool)`

GetProcessStartConfigOk returns a tuple with the ProcessStartConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessStartConfig

`func (o *ProcessExecutionStartRequest) SetProcessStartConfig(v ProcessStartConfig)`

SetProcessStartConfig sets ProcessStartConfig field to given value.

### HasProcessStartConfig

`func (o *ProcessExecutionStartRequest) HasProcessStartConfig() bool`

HasProcessStartConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


