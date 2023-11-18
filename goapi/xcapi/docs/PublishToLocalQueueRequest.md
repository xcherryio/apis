# PublishToLocalQueueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** |  | 
**ProcessId** | **string** |  | 
**Messages** | Pointer to [**[]LocalQueueMessage**](LocalQueueMessage.md) |  | [optional] 

## Methods

### NewPublishToLocalQueueRequest

`func NewPublishToLocalQueueRequest(namespace string, processId string, ) *PublishToLocalQueueRequest`

NewPublishToLocalQueueRequest instantiates a new PublishToLocalQueueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishToLocalQueueRequestWithDefaults

`func NewPublishToLocalQueueRequestWithDefaults() *PublishToLocalQueueRequest`

NewPublishToLocalQueueRequestWithDefaults instantiates a new PublishToLocalQueueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *PublishToLocalQueueRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PublishToLocalQueueRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PublishToLocalQueueRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetProcessId

`func (o *PublishToLocalQueueRequest) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *PublishToLocalQueueRequest) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *PublishToLocalQueueRequest) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.


### GetMessages

`func (o *PublishToLocalQueueRequest) GetMessages() []LocalQueueMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *PublishToLocalQueueRequest) GetMessagesOk() (*[]LocalQueueMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *PublishToLocalQueueRequest) SetMessages(v []LocalQueueMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *PublishToLocalQueueRequest) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


