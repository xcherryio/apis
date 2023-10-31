# AsyncStateExecuteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StateDecision** | [**StateDecision**](StateDecision.md) |  | 
**PublishToLocalQueue** | Pointer to [**[]LocalQueueMessage**](LocalQueueMessage.md) |  | [optional] 
**WriteToGlobalAttributes** | Pointer to [**[]GlobalAttributeTableRowUpdate**](GlobalAttributeTableRowUpdate.md) |  | [optional] 

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

### GetWriteToGlobalAttributes

`func (o *AsyncStateExecuteResponse) GetWriteToGlobalAttributes() []GlobalAttributeTableRowUpdate`

GetWriteToGlobalAttributes returns the WriteToGlobalAttributes field if non-nil, zero value otherwise.

### GetWriteToGlobalAttributesOk

`func (o *AsyncStateExecuteResponse) GetWriteToGlobalAttributesOk() (*[]GlobalAttributeTableRowUpdate, bool)`

GetWriteToGlobalAttributesOk returns a tuple with the WriteToGlobalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteToGlobalAttributes

`func (o *AsyncStateExecuteResponse) SetWriteToGlobalAttributes(v []GlobalAttributeTableRowUpdate)`

SetWriteToGlobalAttributes sets WriteToGlobalAttributes field to given value.

### HasWriteToGlobalAttributes

`func (o *AsyncStateExecuteResponse) HasWriteToGlobalAttributes() bool`

HasWriteToGlobalAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


