# ProcessExecutionPublishToLocalQueueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** |  | 
**ProcessId** | **string** |  | 
**QueueName** | **string** |  | 
**DedupId** | Pointer to **string** | UUID to uniquely distinguish different messages. If not specified, the server will generate a UUID instead. | [optional] 
**Payload** | Pointer to [**EncodedObject**](EncodedObject.md) |  | [optional] 

## Methods

### NewProcessExecutionPublishToLocalQueueRequest

`func NewProcessExecutionPublishToLocalQueueRequest(namespace string, processId string, queueName string, ) *ProcessExecutionPublishToLocalQueueRequest`

NewProcessExecutionPublishToLocalQueueRequest instantiates a new ProcessExecutionPublishToLocalQueueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessExecutionPublishToLocalQueueRequestWithDefaults

`func NewProcessExecutionPublishToLocalQueueRequestWithDefaults() *ProcessExecutionPublishToLocalQueueRequest`

NewProcessExecutionPublishToLocalQueueRequestWithDefaults instantiates a new ProcessExecutionPublishToLocalQueueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *ProcessExecutionPublishToLocalQueueRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ProcessExecutionPublishToLocalQueueRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ProcessExecutionPublishToLocalQueueRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetProcessId

`func (o *ProcessExecutionPublishToLocalQueueRequest) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *ProcessExecutionPublishToLocalQueueRequest) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *ProcessExecutionPublishToLocalQueueRequest) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.


### GetQueueName

`func (o *ProcessExecutionPublishToLocalQueueRequest) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *ProcessExecutionPublishToLocalQueueRequest) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *ProcessExecutionPublishToLocalQueueRequest) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.


### GetDedupId

`func (o *ProcessExecutionPublishToLocalQueueRequest) GetDedupId() string`

GetDedupId returns the DedupId field if non-nil, zero value otherwise.

### GetDedupIdOk

`func (o *ProcessExecutionPublishToLocalQueueRequest) GetDedupIdOk() (*string, bool)`

GetDedupIdOk returns a tuple with the DedupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupId

`func (o *ProcessExecutionPublishToLocalQueueRequest) SetDedupId(v string)`

SetDedupId sets DedupId field to given value.

### HasDedupId

`func (o *ProcessExecutionPublishToLocalQueueRequest) HasDedupId() bool`

HasDedupId returns a boolean if a field has been set.

### GetPayload

`func (o *ProcessExecutionPublishToLocalQueueRequest) GetPayload() EncodedObject`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ProcessExecutionPublishToLocalQueueRequest) GetPayloadOk() (*EncodedObject, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ProcessExecutionPublishToLocalQueueRequest) SetPayload(v EncodedObject)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ProcessExecutionPublishToLocalQueueRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


