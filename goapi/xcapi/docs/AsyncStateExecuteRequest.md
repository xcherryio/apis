# AsyncStateExecuteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**Context**](Context.md) |  | 
**ProcessType** | **string** |  | 
**StateId** | **string** |  | 
**StateInput** | Pointer to [**EncodedObject**](EncodedObject.md) |  | [optional] 
**CommandResults** | Pointer to [**CommandResults**](CommandResults.md) |  | [optional] 
**AppDatabaseReadResponse** | Pointer to [**AppDatabaseReadResponse**](AppDatabaseReadResponse.md) |  | [optional] 
**AppDatabaseError** | Pointer to [**AppDatabaseError**](AppDatabaseError.md) |  | [optional] 
**LoadedLocalAttributes** | Pointer to [**LoadLocalAttributesResponse**](LoadLocalAttributesResponse.md) |  | [optional] 

## Methods

### NewAsyncStateExecuteRequest

`func NewAsyncStateExecuteRequest(context Context, processType string, stateId string, ) *AsyncStateExecuteRequest`

NewAsyncStateExecuteRequest instantiates a new AsyncStateExecuteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncStateExecuteRequestWithDefaults

`func NewAsyncStateExecuteRequestWithDefaults() *AsyncStateExecuteRequest`

NewAsyncStateExecuteRequestWithDefaults instantiates a new AsyncStateExecuteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *AsyncStateExecuteRequest) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AsyncStateExecuteRequest) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AsyncStateExecuteRequest) SetContext(v Context)`

SetContext sets Context field to given value.


### GetProcessType

`func (o *AsyncStateExecuteRequest) GetProcessType() string`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *AsyncStateExecuteRequest) GetProcessTypeOk() (*string, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *AsyncStateExecuteRequest) SetProcessType(v string)`

SetProcessType sets ProcessType field to given value.


### GetStateId

`func (o *AsyncStateExecuteRequest) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *AsyncStateExecuteRequest) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *AsyncStateExecuteRequest) SetStateId(v string)`

SetStateId sets StateId field to given value.


### GetStateInput

`func (o *AsyncStateExecuteRequest) GetStateInput() EncodedObject`

GetStateInput returns the StateInput field if non-nil, zero value otherwise.

### GetStateInputOk

`func (o *AsyncStateExecuteRequest) GetStateInputOk() (*EncodedObject, bool)`

GetStateInputOk returns a tuple with the StateInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateInput

`func (o *AsyncStateExecuteRequest) SetStateInput(v EncodedObject)`

SetStateInput sets StateInput field to given value.

### HasStateInput

`func (o *AsyncStateExecuteRequest) HasStateInput() bool`

HasStateInput returns a boolean if a field has been set.

### GetCommandResults

`func (o *AsyncStateExecuteRequest) GetCommandResults() CommandResults`

GetCommandResults returns the CommandResults field if non-nil, zero value otherwise.

### GetCommandResultsOk

`func (o *AsyncStateExecuteRequest) GetCommandResultsOk() (*CommandResults, bool)`

GetCommandResultsOk returns a tuple with the CommandResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandResults

`func (o *AsyncStateExecuteRequest) SetCommandResults(v CommandResults)`

SetCommandResults sets CommandResults field to given value.

### HasCommandResults

`func (o *AsyncStateExecuteRequest) HasCommandResults() bool`

HasCommandResults returns a boolean if a field has been set.

### GetAppDatabaseReadResponse

`func (o *AsyncStateExecuteRequest) GetAppDatabaseReadResponse() AppDatabaseReadResponse`

GetAppDatabaseReadResponse returns the AppDatabaseReadResponse field if non-nil, zero value otherwise.

### GetAppDatabaseReadResponseOk

`func (o *AsyncStateExecuteRequest) GetAppDatabaseReadResponseOk() (*AppDatabaseReadResponse, bool)`

GetAppDatabaseReadResponseOk returns a tuple with the AppDatabaseReadResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDatabaseReadResponse

`func (o *AsyncStateExecuteRequest) SetAppDatabaseReadResponse(v AppDatabaseReadResponse)`

SetAppDatabaseReadResponse sets AppDatabaseReadResponse field to given value.

### HasAppDatabaseReadResponse

`func (o *AsyncStateExecuteRequest) HasAppDatabaseReadResponse() bool`

HasAppDatabaseReadResponse returns a boolean if a field has been set.

### GetAppDatabaseError

`func (o *AsyncStateExecuteRequest) GetAppDatabaseError() AppDatabaseError`

GetAppDatabaseError returns the AppDatabaseError field if non-nil, zero value otherwise.

### GetAppDatabaseErrorOk

`func (o *AsyncStateExecuteRequest) GetAppDatabaseErrorOk() (*AppDatabaseError, bool)`

GetAppDatabaseErrorOk returns a tuple with the AppDatabaseError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDatabaseError

`func (o *AsyncStateExecuteRequest) SetAppDatabaseError(v AppDatabaseError)`

SetAppDatabaseError sets AppDatabaseError field to given value.

### HasAppDatabaseError

`func (o *AsyncStateExecuteRequest) HasAppDatabaseError() bool`

HasAppDatabaseError returns a boolean if a field has been set.

### GetLoadedLocalAttributes

`func (o *AsyncStateExecuteRequest) GetLoadedLocalAttributes() LoadLocalAttributesResponse`

GetLoadedLocalAttributes returns the LoadedLocalAttributes field if non-nil, zero value otherwise.

### GetLoadedLocalAttributesOk

`func (o *AsyncStateExecuteRequest) GetLoadedLocalAttributesOk() (*LoadLocalAttributesResponse, bool)`

GetLoadedLocalAttributesOk returns a tuple with the LoadedLocalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadedLocalAttributes

`func (o *AsyncStateExecuteRequest) SetLoadedLocalAttributes(v LoadLocalAttributesResponse)`

SetLoadedLocalAttributes sets LoadedLocalAttributes field to given value.

### HasLoadedLocalAttributes

`func (o *AsyncStateExecuteRequest) HasLoadedLocalAttributes() bool`

HasLoadedLocalAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


