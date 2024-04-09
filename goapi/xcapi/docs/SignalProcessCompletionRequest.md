# SignalProcessCompletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShardId** | **int32** |  | 
**ProcessExecutionId** | **string** |  | 
**Status** | **string** |  | 

## Methods

### NewSignalProcessCompletionRequest

`func NewSignalProcessCompletionRequest(shardId int32, processExecutionId string, status string, ) *SignalProcessCompletionRequest`

NewSignalProcessCompletionRequest instantiates a new SignalProcessCompletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalProcessCompletionRequestWithDefaults

`func NewSignalProcessCompletionRequestWithDefaults() *SignalProcessCompletionRequest`

NewSignalProcessCompletionRequestWithDefaults instantiates a new SignalProcessCompletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShardId

`func (o *SignalProcessCompletionRequest) GetShardId() int32`

GetShardId returns the ShardId field if non-nil, zero value otherwise.

### GetShardIdOk

`func (o *SignalProcessCompletionRequest) GetShardIdOk() (*int32, bool)`

GetShardIdOk returns a tuple with the ShardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardId

`func (o *SignalProcessCompletionRequest) SetShardId(v int32)`

SetShardId sets ShardId field to given value.


### GetProcessExecutionId

`func (o *SignalProcessCompletionRequest) GetProcessExecutionId() string`

GetProcessExecutionId returns the ProcessExecutionId field if non-nil, zero value otherwise.

### GetProcessExecutionIdOk

`func (o *SignalProcessCompletionRequest) GetProcessExecutionIdOk() (*string, bool)`

GetProcessExecutionIdOk returns a tuple with the ProcessExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessExecutionId

`func (o *SignalProcessCompletionRequest) SetProcessExecutionId(v string)`

SetProcessExecutionId sets ProcessExecutionId field to given value.


### GetStatus

`func (o *SignalProcessCompletionRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SignalProcessCompletionRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SignalProcessCompletionRequest) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


