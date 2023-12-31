# NotifyTimerTasksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShardId** | **int32** |  | 
**FireTimestamps** | **[]int64** | the fire timestamp of all timer tasks to pull | 
**Namespace** | Pointer to **string** | optional field for distributed database without global secondary index, to pull for specific task rather than a page | [optional] 
**ProcessId** | Pointer to **string** | optional field for distributed database without global secondary index, to pull for specific task rather than a page | [optional] 
**ProcessExecutionId** | Pointer to **string** | optional field for distributed database without global secondary index, to pull for specific task rather than a page | [optional] 

## Methods

### NewNotifyTimerTasksRequest

`func NewNotifyTimerTasksRequest(shardId int32, fireTimestamps []int64, ) *NotifyTimerTasksRequest`

NewNotifyTimerTasksRequest instantiates a new NotifyTimerTasksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyTimerTasksRequestWithDefaults

`func NewNotifyTimerTasksRequestWithDefaults() *NotifyTimerTasksRequest`

NewNotifyTimerTasksRequestWithDefaults instantiates a new NotifyTimerTasksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShardId

`func (o *NotifyTimerTasksRequest) GetShardId() int32`

GetShardId returns the ShardId field if non-nil, zero value otherwise.

### GetShardIdOk

`func (o *NotifyTimerTasksRequest) GetShardIdOk() (*int32, bool)`

GetShardIdOk returns a tuple with the ShardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardId

`func (o *NotifyTimerTasksRequest) SetShardId(v int32)`

SetShardId sets ShardId field to given value.


### GetFireTimestamps

`func (o *NotifyTimerTasksRequest) GetFireTimestamps() []int64`

GetFireTimestamps returns the FireTimestamps field if non-nil, zero value otherwise.

### GetFireTimestampsOk

`func (o *NotifyTimerTasksRequest) GetFireTimestampsOk() (*[]int64, bool)`

GetFireTimestampsOk returns a tuple with the FireTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFireTimestamps

`func (o *NotifyTimerTasksRequest) SetFireTimestamps(v []int64)`

SetFireTimestamps sets FireTimestamps field to given value.


### GetNamespace

`func (o *NotifyTimerTasksRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *NotifyTimerTasksRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *NotifyTimerTasksRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *NotifyTimerTasksRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetProcessId

`func (o *NotifyTimerTasksRequest) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *NotifyTimerTasksRequest) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *NotifyTimerTasksRequest) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *NotifyTimerTasksRequest) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetProcessExecutionId

`func (o *NotifyTimerTasksRequest) GetProcessExecutionId() string`

GetProcessExecutionId returns the ProcessExecutionId field if non-nil, zero value otherwise.

### GetProcessExecutionIdOk

`func (o *NotifyTimerTasksRequest) GetProcessExecutionIdOk() (*string, bool)`

GetProcessExecutionIdOk returns a tuple with the ProcessExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessExecutionId

`func (o *NotifyTimerTasksRequest) SetProcessExecutionId(v string)`

SetProcessExecutionId sets ProcessExecutionId field to given value.

### HasProcessExecutionId

`func (o *NotifyTimerTasksRequest) HasProcessExecutionId() bool`

HasProcessExecutionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


