# LocalQueueMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueueName** | **string** |  | 
**DedupId** | Pointer to **string** | UUID to uniquely distinguish different messages. If not specified, the server will generate a UUID instead. | [optional] 
**Payload** | Pointer to [**EncodedObject**](EncodedObject.md) |  | [optional] 

## Methods

### NewLocalQueueMessage

`func NewLocalQueueMessage(queueName string, ) *LocalQueueMessage`

NewLocalQueueMessage instantiates a new LocalQueueMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalQueueMessageWithDefaults

`func NewLocalQueueMessageWithDefaults() *LocalQueueMessage`

NewLocalQueueMessageWithDefaults instantiates a new LocalQueueMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueueName

`func (o *LocalQueueMessage) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *LocalQueueMessage) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *LocalQueueMessage) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.


### GetDedupId

`func (o *LocalQueueMessage) GetDedupId() string`

GetDedupId returns the DedupId field if non-nil, zero value otherwise.

### GetDedupIdOk

`func (o *LocalQueueMessage) GetDedupIdOk() (*string, bool)`

GetDedupIdOk returns a tuple with the DedupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupId

`func (o *LocalQueueMessage) SetDedupId(v string)`

SetDedupId sets DedupId field to given value.

### HasDedupId

`func (o *LocalQueueMessage) HasDedupId() bool`

HasDedupId returns a boolean if a field has been set.

### GetPayload

`func (o *LocalQueueMessage) GetPayload() EncodedObject`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *LocalQueueMessage) GetPayloadOk() (*EncodedObject, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *LocalQueueMessage) SetPayload(v EncodedObject)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *LocalQueueMessage) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


