# AsyncStateExecuteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StateDecision** | [**StateDecision**](StateDecision.md) |  | 
**PublishToLocalQueue** | Pointer to [**[]LocalQueueMessage**](LocalQueueMessage.md) |  | [optional] 
**WriteToAppDatabase** | Pointer to [**AppDatabaseWrite**](AppDatabaseWrite.md) |  | [optional] 
**WriteToLocalAttributes** | Pointer to [**[]KeyValue**](KeyValue.md) |  | [optional] 
**UpdatedAppDatabaseReadRequest** | Pointer to [**AppDatabaseReadRequest**](AppDatabaseReadRequest.md) |  | [optional] 

## Methods

### NewAsyncStateExecuteResponse

`func NewAsyncStateExecuteResponse(stateDecision StateDecision, ) *AsyncStateExecuteResponse`

NewAsyncStateExecuteResponse instantiates a new AsyncStateExecuteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncStateExecuteResponseWithDefaults

`func NewAsyncStateExecuteResponseWithDefaults() *AsyncStateExecuteResponse`

NewAsyncStateExecuteResponseWithDefaults instantiates a new AsyncStateExecuteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStateDecision

`func (o *AsyncStateExecuteResponse) GetStateDecision() StateDecision`

GetStateDecision returns the StateDecision field if non-nil, zero value otherwise.

### GetStateDecisionOk

`func (o *AsyncStateExecuteResponse) GetStateDecisionOk() (*StateDecision, bool)`

GetStateDecisionOk returns a tuple with the StateDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDecision

`func (o *AsyncStateExecuteResponse) SetStateDecision(v StateDecision)`

SetStateDecision sets StateDecision field to given value.


### GetPublishToLocalQueue

`func (o *AsyncStateExecuteResponse) GetPublishToLocalQueue() []LocalQueueMessage`

GetPublishToLocalQueue returns the PublishToLocalQueue field if non-nil, zero value otherwise.

### GetPublishToLocalQueueOk

`func (o *AsyncStateExecuteResponse) GetPublishToLocalQueueOk() (*[]LocalQueueMessage, bool)`

GetPublishToLocalQueueOk returns a tuple with the PublishToLocalQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishToLocalQueue

`func (o *AsyncStateExecuteResponse) SetPublishToLocalQueue(v []LocalQueueMessage)`

SetPublishToLocalQueue sets PublishToLocalQueue field to given value.

### HasPublishToLocalQueue

`func (o *AsyncStateExecuteResponse) HasPublishToLocalQueue() bool`

HasPublishToLocalQueue returns a boolean if a field has been set.

### GetWriteToAppDatabase

`func (o *AsyncStateExecuteResponse) GetWriteToAppDatabase() AppDatabaseWrite`

GetWriteToAppDatabase returns the WriteToAppDatabase field if non-nil, zero value otherwise.

### GetWriteToAppDatabaseOk

`func (o *AsyncStateExecuteResponse) GetWriteToAppDatabaseOk() (*AppDatabaseWrite, bool)`

GetWriteToAppDatabaseOk returns a tuple with the WriteToAppDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteToAppDatabase

`func (o *AsyncStateExecuteResponse) SetWriteToAppDatabase(v AppDatabaseWrite)`

SetWriteToAppDatabase sets WriteToAppDatabase field to given value.

### HasWriteToAppDatabase

`func (o *AsyncStateExecuteResponse) HasWriteToAppDatabase() bool`

HasWriteToAppDatabase returns a boolean if a field has been set.

### GetWriteToLocalAttributes

`func (o *AsyncStateExecuteResponse) GetWriteToLocalAttributes() []KeyValue`

GetWriteToLocalAttributes returns the WriteToLocalAttributes field if non-nil, zero value otherwise.

### GetWriteToLocalAttributesOk

`func (o *AsyncStateExecuteResponse) GetWriteToLocalAttributesOk() (*[]KeyValue, bool)`

GetWriteToLocalAttributesOk returns a tuple with the WriteToLocalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteToLocalAttributes

`func (o *AsyncStateExecuteResponse) SetWriteToLocalAttributes(v []KeyValue)`

SetWriteToLocalAttributes sets WriteToLocalAttributes field to given value.

### HasWriteToLocalAttributes

`func (o *AsyncStateExecuteResponse) HasWriteToLocalAttributes() bool`

HasWriteToLocalAttributes returns a boolean if a field has been set.

### GetUpdatedAppDatabaseReadRequest

`func (o *AsyncStateExecuteResponse) GetUpdatedAppDatabaseReadRequest() AppDatabaseReadRequest`

GetUpdatedAppDatabaseReadRequest returns the UpdatedAppDatabaseReadRequest field if non-nil, zero value otherwise.

### GetUpdatedAppDatabaseReadRequestOk

`func (o *AsyncStateExecuteResponse) GetUpdatedAppDatabaseReadRequestOk() (*AppDatabaseReadRequest, bool)`

GetUpdatedAppDatabaseReadRequestOk returns a tuple with the UpdatedAppDatabaseReadRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAppDatabaseReadRequest

`func (o *AsyncStateExecuteResponse) SetUpdatedAppDatabaseReadRequest(v AppDatabaseReadRequest)`

SetUpdatedAppDatabaseReadRequest sets UpdatedAppDatabaseReadRequest field to given value.

### HasUpdatedAppDatabaseReadRequest

`func (o *AsyncStateExecuteResponse) HasUpdatedAppDatabaseReadRequest() bool`

HasUpdatedAppDatabaseReadRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


