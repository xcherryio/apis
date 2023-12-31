# ProcessRpcWorkerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Output** | Pointer to [**EncodedObject**](EncodedObject.md) |  | [optional] 
**StateDecision** | [**StateDecision**](StateDecision.md) |  | 
**PublishToLocalQueue** | Pointer to [**[]LocalQueueMessage**](LocalQueueMessage.md) |  | [optional] 
**WriteToAppDatabase** | Pointer to [**AppDatabaseWrite**](AppDatabaseWrite.md) |  | [optional] 

## Methods

### NewProcessRpcWorkerResponse

`func NewProcessRpcWorkerResponse(stateDecision StateDecision, ) *ProcessRpcWorkerResponse`

NewProcessRpcWorkerResponse instantiates a new ProcessRpcWorkerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessRpcWorkerResponseWithDefaults

`func NewProcessRpcWorkerResponseWithDefaults() *ProcessRpcWorkerResponse`

NewProcessRpcWorkerResponseWithDefaults instantiates a new ProcessRpcWorkerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutput

`func (o *ProcessRpcWorkerResponse) GetOutput() EncodedObject`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ProcessRpcWorkerResponse) GetOutputOk() (*EncodedObject, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ProcessRpcWorkerResponse) SetOutput(v EncodedObject)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ProcessRpcWorkerResponse) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetStateDecision

`func (o *ProcessRpcWorkerResponse) GetStateDecision() StateDecision`

GetStateDecision returns the StateDecision field if non-nil, zero value otherwise.

### GetStateDecisionOk

`func (o *ProcessRpcWorkerResponse) GetStateDecisionOk() (*StateDecision, bool)`

GetStateDecisionOk returns a tuple with the StateDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDecision

`func (o *ProcessRpcWorkerResponse) SetStateDecision(v StateDecision)`

SetStateDecision sets StateDecision field to given value.


### GetPublishToLocalQueue

`func (o *ProcessRpcWorkerResponse) GetPublishToLocalQueue() []LocalQueueMessage`

GetPublishToLocalQueue returns the PublishToLocalQueue field if non-nil, zero value otherwise.

### GetPublishToLocalQueueOk

`func (o *ProcessRpcWorkerResponse) GetPublishToLocalQueueOk() (*[]LocalQueueMessage, bool)`

GetPublishToLocalQueueOk returns a tuple with the PublishToLocalQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishToLocalQueue

`func (o *ProcessRpcWorkerResponse) SetPublishToLocalQueue(v []LocalQueueMessage)`

SetPublishToLocalQueue sets PublishToLocalQueue field to given value.

### HasPublishToLocalQueue

`func (o *ProcessRpcWorkerResponse) HasPublishToLocalQueue() bool`

HasPublishToLocalQueue returns a boolean if a field has been set.

### GetWriteToAppDatabase

`func (o *ProcessRpcWorkerResponse) GetWriteToAppDatabase() AppDatabaseWrite`

GetWriteToAppDatabase returns the WriteToAppDatabase field if non-nil, zero value otherwise.

### GetWriteToAppDatabaseOk

`func (o *ProcessRpcWorkerResponse) GetWriteToAppDatabaseOk() (*AppDatabaseWrite, bool)`

GetWriteToAppDatabaseOk returns a tuple with the WriteToAppDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteToAppDatabase

`func (o *ProcessRpcWorkerResponse) SetWriteToAppDatabase(v AppDatabaseWrite)`

SetWriteToAppDatabase sets WriteToAppDatabase field to given value.

### HasWriteToAppDatabase

`func (o *ProcessRpcWorkerResponse) HasWriteToAppDatabase() bool`

HasWriteToAppDatabase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


