# AsyncStateWaitUntilRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**Context**](Context.md) |  | 
**ProcessType** | **string** |  | 
**StateId** | **string** |  | 
**StateInput** | Pointer to [**EncodedObject**](EncodedObject.md) |  | [optional] 

## Methods

### NewAsyncStateWaitUntilRequest

`func NewAsyncStateWaitUntilRequest(context Context, processType string, stateId string, ) *AsyncStateWaitUntilRequest`

NewAsyncStateWaitUntilRequest instantiates a new AsyncStateWaitUntilRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncStateWaitUntilRequestWithDefaults

`func NewAsyncStateWaitUntilRequestWithDefaults() *AsyncStateWaitUntilRequest`

NewAsyncStateWaitUntilRequestWithDefaults instantiates a new AsyncStateWaitUntilRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *AsyncStateWaitUntilRequest) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AsyncStateWaitUntilRequest) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AsyncStateWaitUntilRequest) SetContext(v Context)`

SetContext sets Context field to given value.


### GetProcessType

`func (o *AsyncStateWaitUntilRequest) GetProcessType() string`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *AsyncStateWaitUntilRequest) GetProcessTypeOk() (*string, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *AsyncStateWaitUntilRequest) SetProcessType(v string)`

SetProcessType sets ProcessType field to given value.


### GetStateId

`func (o *AsyncStateWaitUntilRequest) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *AsyncStateWaitUntilRequest) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *AsyncStateWaitUntilRequest) SetStateId(v string)`

SetStateId sets StateId field to given value.


### GetStateInput

`func (o *AsyncStateWaitUntilRequest) GetStateInput() EncodedObject`

GetStateInput returns the StateInput field if non-nil, zero value otherwise.

### GetStateInputOk

`func (o *AsyncStateWaitUntilRequest) GetStateInputOk() (*EncodedObject, bool)`

GetStateInputOk returns a tuple with the StateInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateInput

`func (o *AsyncStateWaitUntilRequest) SetStateInput(v EncodedObject)`

SetStateInput sets StateInput field to given value.

### HasStateInput

`func (o *AsyncStateWaitUntilRequest) HasStateInput() bool`

HasStateInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


