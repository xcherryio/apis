# ProcessRpcResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Output** | Pointer to [**EncodedObject**](EncodedObject.md) |  | [optional] 
**StateDecision** | [**StateDecision**](StateDecision.md) |  | 
**PublishToLocalQueue** | Pointer to [**[]LocalQueueMessage**](LocalQueueMessage.md) |  | [optional] 
**WriteToGlobalAttributes** | Pointer to [**[]GlobalAttributeTableRowUpdate**](GlobalAttributeTableRowUpdate.md) |  | [optional] 

## Methods

### NewProcessRpcResponse

`func NewProcessRpcResponse(stateDecision StateDecision, ) *ProcessRpcResponse`

NewProcessRpcResponse instantiates a new ProcessRpcResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessRpcResponseWithDefaults

`func NewProcessRpcResponseWithDefaults() *ProcessRpcResponse`

NewProcessRpcResponseWithDefaults instantiates a new ProcessRpcResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutput

`func (o *ProcessRpcResponse) GetOutput() EncodedObject`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ProcessRpcResponse) GetOutputOk() (*EncodedObject, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ProcessRpcResponse) SetOutput(v EncodedObject)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ProcessRpcResponse) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetStateDecision

`func (o *ProcessRpcResponse) GetStateDecision() StateDecision`

GetStateDecision returns the StateDecision field if non-nil, zero value otherwise.

### GetStateDecisionOk

`func (o *ProcessRpcResponse) GetStateDecisionOk() (*StateDecision, bool)`

GetStateDecisionOk returns a tuple with the StateDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDecision

`func (o *ProcessRpcResponse) SetStateDecision(v StateDecision)`

SetStateDecision sets StateDecision field to given value.


### GetPublishToLocalQueue

`func (o *ProcessRpcResponse) GetPublishToLocalQueue() []LocalQueueMessage`

GetPublishToLocalQueue returns the PublishToLocalQueue field if non-nil, zero value otherwise.

### GetPublishToLocalQueueOk

`func (o *ProcessRpcResponse) GetPublishToLocalQueueOk() (*[]LocalQueueMessage, bool)`

GetPublishToLocalQueueOk returns a tuple with the PublishToLocalQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishToLocalQueue

`func (o *ProcessRpcResponse) SetPublishToLocalQueue(v []LocalQueueMessage)`

SetPublishToLocalQueue sets PublishToLocalQueue field to given value.

### HasPublishToLocalQueue

`func (o *ProcessRpcResponse) HasPublishToLocalQueue() bool`

HasPublishToLocalQueue returns a boolean if a field has been set.

### GetWriteToGlobalAttributes

`func (o *ProcessRpcResponse) GetWriteToGlobalAttributes() []GlobalAttributeTableRowUpdate`

GetWriteToGlobalAttributes returns the WriteToGlobalAttributes field if non-nil, zero value otherwise.

### GetWriteToGlobalAttributesOk

`func (o *ProcessRpcResponse) GetWriteToGlobalAttributesOk() (*[]GlobalAttributeTableRowUpdate, bool)`

GetWriteToGlobalAttributesOk returns a tuple with the WriteToGlobalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteToGlobalAttributes

`func (o *ProcessRpcResponse) SetWriteToGlobalAttributes(v []GlobalAttributeTableRowUpdate)`

SetWriteToGlobalAttributes sets WriteToGlobalAttributes field to given value.

### HasWriteToGlobalAttributes

`func (o *ProcessRpcResponse) HasWriteToGlobalAttributes() bool`

HasWriteToGlobalAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


