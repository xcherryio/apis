# ProcessExecutionSendMessageToLocalQueueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** |  | 
**ProcessId** | **string** |  | 
**QueueName** | **string** |  | 
**DedupId** | Pointer to **string** | UUID to uniquely distinguish different messages. If not specified, the server will generate a UUID instead. | [optional] 
**Payload** | Pointer to [**EncodedObject**](EncodedObject.md) |  | [optional] 

## Methods

### NewProcessExecutionSendMessageToLocalQueueRequest

`func NewProcessExecutionSendMessageToLocalQueueRequest(namespace string, processId string, queueName string, ) *ProcessExecutionSendMessageToLocalQueueRequest`

NewProcessExecutionSendMessageToLocalQueueRequest instantiates a new ProcessExecutionSendMessageToLocalQueueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessExecutionSendMessageToLocalQueueRequestWithDefaults

`func NewProcessExecutionSendMessageToLocalQueueRequestWithDefaults() *ProcessExecutionSendMessageToLocalQueueRequest`

NewProcessExecutionSendMessageToLocalQueueRequestWithDefaults instantiates a new ProcessExecutionSendMessageToLocalQueueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ProcessExecutionSendMessageToLocalQueueRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetProcessId

`func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *ProcessExecutionSendMessageToLocalQueueRequest) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.


### GetQueueName

`func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *ProcessExecutionSendMessageToLocalQueueRequest) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.


### GetDedupId

`func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetDedupId() string`

GetDedupId returns the DedupId field if non-nil, zero value otherwise.

### GetDedupIdOk

`func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetDedupIdOk() (*string, bool)`

GetDedupIdOk returns a tuple with the DedupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupId

`func (o *ProcessExecutionSendMessageToLocalQueueRequest) SetDedupId(v string)`

SetDedupId sets DedupId field to given value.

### HasDedupId

`func (o *ProcessExecutionSendMessageToLocalQueueRequest) HasDedupId() bool`

HasDedupId returns a boolean if a field has been set.

### GetPayload

`func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetPayload() EncodedObject`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetPayloadOk() (*EncodedObject, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ProcessExecutionSendMessageToLocalQueueRequest) SetPayload(v EncodedObject)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ProcessExecutionSendMessageToLocalQueueRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


