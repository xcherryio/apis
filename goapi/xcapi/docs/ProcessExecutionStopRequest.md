# ProcessExecutionStopRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** |  | 
**ProcessId** | **string** |  | 
**StopType** | Pointer to [**ProcessExecutionStopType**](ProcessExecutionStopType.md) |  | [optional] 

## Methods

### NewProcessExecutionStopRequest

`func NewProcessExecutionStopRequest(namespace string, processId string, ) *ProcessExecutionStopRequest`

NewProcessExecutionStopRequest instantiates a new ProcessExecutionStopRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessExecutionStopRequestWithDefaults

`func NewProcessExecutionStopRequestWithDefaults() *ProcessExecutionStopRequest`

NewProcessExecutionStopRequestWithDefaults instantiates a new ProcessExecutionStopRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *ProcessExecutionStopRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ProcessExecutionStopRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ProcessExecutionStopRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetProcessId

`func (o *ProcessExecutionStopRequest) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *ProcessExecutionStopRequest) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *ProcessExecutionStopRequest) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.


### GetStopType

`func (o *ProcessExecutionStopRequest) GetStopType() ProcessExecutionStopType`

GetStopType returns the StopType field if non-nil, zero value otherwise.

### GetStopTypeOk

`func (o *ProcessExecutionStopRequest) GetStopTypeOk() (*ProcessExecutionStopType, bool)`

GetStopTypeOk returns a tuple with the StopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopType

`func (o *ProcessExecutionStopRequest) SetStopType(v ProcessExecutionStopType)`

SetStopType sets StopType field to given value.

### HasStopType

`func (o *ProcessExecutionStopRequest) HasStopType() bool`

HasStopType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


