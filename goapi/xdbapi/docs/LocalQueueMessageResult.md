# LocalQueueMessageResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DedupId** | **string** | UUID to uniquely distinguish different messages. | 
**Payload** | Pointer to [**EncodedObject**](EncodedObject.md) |  | [optional] 

## Methods

### NewLocalQueueMessageResult

`func NewLocalQueueMessageResult(dedupId string, ) *LocalQueueMessageResult`

NewLocalQueueMessageResult instantiates a new LocalQueueMessageResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalQueueMessageResultWithDefaults

`func NewLocalQueueMessageResultWithDefaults() *LocalQueueMessageResult`

NewLocalQueueMessageResultWithDefaults instantiates a new LocalQueueMessageResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDedupId

`func (o *LocalQueueMessageResult) GetDedupId() string`

GetDedupId returns the DedupId field if non-nil, zero value otherwise.

### GetDedupIdOk

`func (o *LocalQueueMessageResult) GetDedupIdOk() (*string, bool)`

GetDedupIdOk returns a tuple with the DedupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupId

`func (o *LocalQueueMessageResult) SetDedupId(v string)`

SetDedupId sets DedupId field to given value.


### GetPayload

`func (o *LocalQueueMessageResult) GetPayload() EncodedObject`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *LocalQueueMessageResult) GetPayloadOk() (*EncodedObject, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *LocalQueueMessageResult) SetPayload(v EncodedObject)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *LocalQueueMessageResult) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


