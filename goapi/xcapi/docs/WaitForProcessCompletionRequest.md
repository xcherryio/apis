# WaitForProcessCompletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShardId** | **int32** |  | 
**ProcessExecutionId** | **string** |  | 

## Methods

### NewWaitForProcessCompletionRequest

`func NewWaitForProcessCompletionRequest(shardId int32, processExecutionId string, ) *WaitForProcessCompletionRequest`

NewWaitForProcessCompletionRequest instantiates a new WaitForProcessCompletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaitForProcessCompletionRequestWithDefaults

`func NewWaitForProcessCompletionRequestWithDefaults() *WaitForProcessCompletionRequest`

NewWaitForProcessCompletionRequestWithDefaults instantiates a new WaitForProcessCompletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShardId

`func (o *WaitForProcessCompletionRequest) GetShardId() int32`

GetShardId returns the ShardId field if non-nil, zero value otherwise.

### GetShardIdOk

`func (o *WaitForProcessCompletionRequest) GetShardIdOk() (*int32, bool)`

GetShardIdOk returns a tuple with the ShardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardId

`func (o *WaitForProcessCompletionRequest) SetShardId(v int32)`

SetShardId sets ShardId field to given value.


### GetProcessExecutionId

`func (o *WaitForProcessCompletionRequest) GetProcessExecutionId() string`

GetProcessExecutionId returns the ProcessExecutionId field if non-nil, zero value otherwise.

### GetProcessExecutionIdOk

`func (o *WaitForProcessCompletionRequest) GetProcessExecutionIdOk() (*string, bool)`

GetProcessExecutionIdOk returns a tuple with the ProcessExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessExecutionId

`func (o *WaitForProcessCompletionRequest) SetProcessExecutionId(v string)`

SetProcessExecutionId sets ProcessExecutionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


