# ProcessExecutionRpcRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** |  | 
**ProcessId** | **string** |  | 
**RpcName** | **string** |  | 
**Input** | Pointer to [**EncodedObject**](EncodedObject.md) |  | [optional] 
**TimeoutSeconds** | Pointer to **int32** | the timeout for the single attempt of the Process RPC API | [optional] 
**LoadGlobalAttributesRequest** | Pointer to [**LoadGlobalAttributesRequest**](LoadGlobalAttributesRequest.md) |  | [optional] 

## Methods

### NewProcessExecutionRpcRequest

`func NewProcessExecutionRpcRequest(namespace string, processId string, rpcName string, ) *ProcessExecutionRpcRequest`

NewProcessExecutionRpcRequest instantiates a new ProcessExecutionRpcRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessExecutionRpcRequestWithDefaults

`func NewProcessExecutionRpcRequestWithDefaults() *ProcessExecutionRpcRequest`

NewProcessExecutionRpcRequestWithDefaults instantiates a new ProcessExecutionRpcRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *ProcessExecutionRpcRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ProcessExecutionRpcRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ProcessExecutionRpcRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetProcessId

`func (o *ProcessExecutionRpcRequest) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *ProcessExecutionRpcRequest) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *ProcessExecutionRpcRequest) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.


### GetRpcName

`func (o *ProcessExecutionRpcRequest) GetRpcName() string`

GetRpcName returns the RpcName field if non-nil, zero value otherwise.

### GetRpcNameOk

`func (o *ProcessExecutionRpcRequest) GetRpcNameOk() (*string, bool)`

GetRpcNameOk returns a tuple with the RpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcName

`func (o *ProcessExecutionRpcRequest) SetRpcName(v string)`

SetRpcName sets RpcName field to given value.


### GetInput

`func (o *ProcessExecutionRpcRequest) GetInput() EncodedObject`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ProcessExecutionRpcRequest) GetInputOk() (*EncodedObject, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ProcessExecutionRpcRequest) SetInput(v EncodedObject)`

SetInput sets Input field to given value.

### HasInput

`func (o *ProcessExecutionRpcRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *ProcessExecutionRpcRequest) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *ProcessExecutionRpcRequest) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *ProcessExecutionRpcRequest) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *ProcessExecutionRpcRequest) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetLoadGlobalAttributesRequest

`func (o *ProcessExecutionRpcRequest) GetLoadGlobalAttributesRequest() LoadGlobalAttributesRequest`

GetLoadGlobalAttributesRequest returns the LoadGlobalAttributesRequest field if non-nil, zero value otherwise.

### GetLoadGlobalAttributesRequestOk

`func (o *ProcessExecutionRpcRequest) GetLoadGlobalAttributesRequestOk() (*LoadGlobalAttributesRequest, bool)`

GetLoadGlobalAttributesRequestOk returns a tuple with the LoadGlobalAttributesRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadGlobalAttributesRequest

`func (o *ProcessExecutionRpcRequest) SetLoadGlobalAttributesRequest(v LoadGlobalAttributesRequest)`

SetLoadGlobalAttributesRequest sets LoadGlobalAttributesRequest field to given value.

### HasLoadGlobalAttributesRequest

`func (o *ProcessExecutionRpcRequest) HasLoadGlobalAttributesRequest() bool`

HasLoadGlobalAttributesRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


