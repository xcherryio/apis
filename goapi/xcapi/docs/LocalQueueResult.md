# LocalQueueResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**CommandStatus**](CommandStatus.md) |  | 
**QueueName** | **string** |  | 
**Messages** | Pointer to [**[]LocalQueueMessageResult**](LocalQueueMessageResult.md) |  | [optional] 

## Methods

### NewLocalQueueResult

`func NewLocalQueueResult(status CommandStatus, queueName string, ) *LocalQueueResult`

NewLocalQueueResult instantiates a new LocalQueueResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalQueueResultWithDefaults

`func NewLocalQueueResultWithDefaults() *LocalQueueResult`

NewLocalQueueResultWithDefaults instantiates a new LocalQueueResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *LocalQueueResult) GetStatus() CommandStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LocalQueueResult) GetStatusOk() (*CommandStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LocalQueueResult) SetStatus(v CommandStatus)`

SetStatus sets Status field to given value.


### GetQueueName

`func (o *LocalQueueResult) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *LocalQueueResult) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *LocalQueueResult) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.


### GetMessages

`func (o *LocalQueueResult) GetMessages() []LocalQueueMessageResult`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *LocalQueueResult) GetMessagesOk() (*[]LocalQueueMessageResult, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *LocalQueueResult) SetMessages(v []LocalQueueMessageResult)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *LocalQueueResult) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


