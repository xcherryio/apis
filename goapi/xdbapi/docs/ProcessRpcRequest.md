# ProcessRpcRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**Context**](Context.md) |  | 
**ProcessType** | **string** |  | 
**RpcName** | **string** |  | 
**Input** | Pointer to [**EncodedObject**](EncodedObject.md) |  | [optional] 
**CommandResults** | Pointer to [**CommandResults**](CommandResults.md) |  | [optional] 
**LoadedGlobalAttributes** | Pointer to [**LoadGlobalAttributeResponse**](LoadGlobalAttributeResponse.md) |  | [optional] 

## Methods

### NewProcessRpcRequest

`func NewProcessRpcRequest(context Context, processType string, rpcName string, ) *ProcessRpcRequest`

NewProcessRpcRequest instantiates a new ProcessRpcRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessRpcRequestWithDefaults

`func NewProcessRpcRequestWithDefaults() *ProcessRpcRequest`

NewProcessRpcRequestWithDefaults instantiates a new ProcessRpcRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ProcessRpcRequest) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ProcessRpcRequest) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ProcessRpcRequest) SetContext(v Context)`

SetContext sets Context field to given value.


### GetProcessType

`func (o *ProcessRpcRequest) GetProcessType() string`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *ProcessRpcRequest) GetProcessTypeOk() (*string, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *ProcessRpcRequest) SetProcessType(v string)`

SetProcessType sets ProcessType field to given value.


### GetRpcName

`func (o *ProcessRpcRequest) GetRpcName() string`

GetRpcName returns the RpcName field if non-nil, zero value otherwise.

### GetRpcNameOk

`func (o *ProcessRpcRequest) GetRpcNameOk() (*string, bool)`

GetRpcNameOk returns a tuple with the RpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcName

`func (o *ProcessRpcRequest) SetRpcName(v string)`

SetRpcName sets RpcName field to given value.


### GetInput

`func (o *ProcessRpcRequest) GetInput() EncodedObject`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ProcessRpcRequest) GetInputOk() (*EncodedObject, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ProcessRpcRequest) SetInput(v EncodedObject)`

SetInput sets Input field to given value.

### HasInput

`func (o *ProcessRpcRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetCommandResults

`func (o *ProcessRpcRequest) GetCommandResults() CommandResults`

GetCommandResults returns the CommandResults field if non-nil, zero value otherwise.

### GetCommandResultsOk

`func (o *ProcessRpcRequest) GetCommandResultsOk() (*CommandResults, bool)`

GetCommandResultsOk returns a tuple with the CommandResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandResults

`func (o *ProcessRpcRequest) SetCommandResults(v CommandResults)`

SetCommandResults sets CommandResults field to given value.

### HasCommandResults

`func (o *ProcessRpcRequest) HasCommandResults() bool`

HasCommandResults returns a boolean if a field has been set.

### GetLoadedGlobalAttributes

`func (o *ProcessRpcRequest) GetLoadedGlobalAttributes() LoadGlobalAttributeResponse`

GetLoadedGlobalAttributes returns the LoadedGlobalAttributes field if non-nil, zero value otherwise.

### GetLoadedGlobalAttributesOk

`func (o *ProcessRpcRequest) GetLoadedGlobalAttributesOk() (*LoadGlobalAttributeResponse, bool)`

GetLoadedGlobalAttributesOk returns a tuple with the LoadedGlobalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadedGlobalAttributes

`func (o *ProcessRpcRequest) SetLoadedGlobalAttributes(v LoadGlobalAttributeResponse)`

SetLoadedGlobalAttributes sets LoadedGlobalAttributes field to given value.

### HasLoadedGlobalAttributes

`func (o *ProcessRpcRequest) HasLoadedGlobalAttributes() bool`

HasLoadedGlobalAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


