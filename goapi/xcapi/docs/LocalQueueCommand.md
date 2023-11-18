# LocalQueueCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueueName** | **string** |  | 
**Count** | Pointer to **int32** | the number of identical messages to await | [optional] [default to 1]

## Methods

### NewLocalQueueCommand

`func NewLocalQueueCommand(queueName string, ) *LocalQueueCommand`

NewLocalQueueCommand instantiates a new LocalQueueCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalQueueCommandWithDefaults

`func NewLocalQueueCommandWithDefaults() *LocalQueueCommand`

NewLocalQueueCommandWithDefaults instantiates a new LocalQueueCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueueName

`func (o *LocalQueueCommand) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *LocalQueueCommand) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *LocalQueueCommand) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.


### GetCount

`func (o *LocalQueueCommand) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *LocalQueueCommand) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *LocalQueueCommand) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *LocalQueueCommand) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


