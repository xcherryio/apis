# AsyncStateWaitUntilResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandRequest** | [**CommandRequest**](CommandRequest.md) |  | 
**PublishToLocalQueue** | Pointer to [**[]LocalQueueMessage**](LocalQueueMessage.md) |  | [optional] 

## Methods

### NewAsyncStateWaitUntilResponse

`func NewAsyncStateWaitUntilResponse(commandRequest CommandRequest, ) *AsyncStateWaitUntilResponse`

NewAsyncStateWaitUntilResponse instantiates a new AsyncStateWaitUntilResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncStateWaitUntilResponseWithDefaults

`func NewAsyncStateWaitUntilResponseWithDefaults() *AsyncStateWaitUntilResponse`

NewAsyncStateWaitUntilResponseWithDefaults instantiates a new AsyncStateWaitUntilResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandRequest

`func (o *AsyncStateWaitUntilResponse) GetCommandRequest() CommandRequest`

GetCommandRequest returns the CommandRequest field if non-nil, zero value otherwise.

### GetCommandRequestOk

`func (o *AsyncStateWaitUntilResponse) GetCommandRequestOk() (*CommandRequest, bool)`

GetCommandRequestOk returns a tuple with the CommandRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandRequest

`func (o *AsyncStateWaitUntilResponse) SetCommandRequest(v CommandRequest)`

SetCommandRequest sets CommandRequest field to given value.


### GetPublishToLocalQueue

`func (o *AsyncStateWaitUntilResponse) GetPublishToLocalQueue() []LocalQueueMessage`

GetPublishToLocalQueue returns the PublishToLocalQueue field if non-nil, zero value otherwise.

### GetPublishToLocalQueueOk

`func (o *AsyncStateWaitUntilResponse) GetPublishToLocalQueueOk() (*[]LocalQueueMessage, bool)`

GetPublishToLocalQueueOk returns a tuple with the PublishToLocalQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishToLocalQueue

`func (o *AsyncStateWaitUntilResponse) SetPublishToLocalQueue(v []LocalQueueMessage)`

SetPublishToLocalQueue sets PublishToLocalQueue field to given value.

### HasPublishToLocalQueue

`func (o *AsyncStateWaitUntilResponse) HasPublishToLocalQueue() bool`

HasPublishToLocalQueue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


