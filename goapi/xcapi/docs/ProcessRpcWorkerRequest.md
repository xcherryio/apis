# ProcessRpcWorkerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**Context**](Context.md) |  | 
**ProcessType** | **string** |  | 
**RpcName** | **string** |  | 
**Input** | Pointer to [**EncodedObject**](EncodedObject.md) |  | [optional] 
**AppDatabaseReadResponse** | Pointer to [**AppDatabaseReadResponse**](AppDatabaseReadResponse.md) |  | [optional] 

## Methods

### NewProcessRpcWorkerRequest

`func NewProcessRpcWorkerRequest(context Context, processType string, rpcName string, ) *ProcessRpcWorkerRequest`

NewProcessRpcWorkerRequest instantiates a new ProcessRpcWorkerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessRpcWorkerRequestWithDefaults

`func NewProcessRpcWorkerRequestWithDefaults() *ProcessRpcWorkerRequest`

NewProcessRpcWorkerRequestWithDefaults instantiates a new ProcessRpcWorkerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ProcessRpcWorkerRequest) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ProcessRpcWorkerRequest) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ProcessRpcWorkerRequest) SetContext(v Context)`

SetContext sets Context field to given value.


### GetProcessType

`func (o *ProcessRpcWorkerRequest) GetProcessType() string`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *ProcessRpcWorkerRequest) GetProcessTypeOk() (*string, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *ProcessRpcWorkerRequest) SetProcessType(v string)`

SetProcessType sets ProcessType field to given value.


### GetRpcName

`func (o *ProcessRpcWorkerRequest) GetRpcName() string`

GetRpcName returns the RpcName field if non-nil, zero value otherwise.

### GetRpcNameOk

`func (o *ProcessRpcWorkerRequest) GetRpcNameOk() (*string, bool)`

GetRpcNameOk returns a tuple with the RpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcName

`func (o *ProcessRpcWorkerRequest) SetRpcName(v string)`

SetRpcName sets RpcName field to given value.


### GetInput

`func (o *ProcessRpcWorkerRequest) GetInput() EncodedObject`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ProcessRpcWorkerRequest) GetInputOk() (*EncodedObject, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ProcessRpcWorkerRequest) SetInput(v EncodedObject)`

SetInput sets Input field to given value.

### HasInput

`func (o *ProcessRpcWorkerRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetAppDatabaseReadResponse

`func (o *ProcessRpcWorkerRequest) GetAppDatabaseReadResponse() AppDatabaseReadResponse`

GetAppDatabaseReadResponse returns the AppDatabaseReadResponse field if non-nil, zero value otherwise.

### GetAppDatabaseReadResponseOk

`func (o *ProcessRpcWorkerRequest) GetAppDatabaseReadResponseOk() (*AppDatabaseReadResponse, bool)`

GetAppDatabaseReadResponseOk returns a tuple with the AppDatabaseReadResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDatabaseReadResponse

`func (o *ProcessRpcWorkerRequest) SetAppDatabaseReadResponse(v AppDatabaseReadResponse)`

SetAppDatabaseReadResponse sets AppDatabaseReadResponse field to given value.

### HasAppDatabaseReadResponse

`func (o *ProcessRpcWorkerRequest) HasAppDatabaseReadResponse() bool`

HasAppDatabaseReadResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


