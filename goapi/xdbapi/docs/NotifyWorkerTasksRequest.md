# NotifyWorkerTasksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShardId** | **int32** |  | 
**Namespace** | Pointer to **string** | optional field for distributed database without global secondary index, to pull for specific task rather than a page | [optional] 
**ProcessId** | Pointer to **string** | optional field for distributed database without global secondary index, to pull for specific task rather than a page | [optional] 
**ProcessExecutionId** | Pointer to **string** | optional field for distributed database without global secondary index, to pull for specific task rather than a page | [optional] 

## Methods

### NewNotifyWorkerTasksRequest

`func NewNotifyWorkerTasksRequest(shardId int32, ) *NotifyWorkerTasksRequest`

NewNotifyWorkerTasksRequest instantiates a new NotifyWorkerTasksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyWorkerTasksRequestWithDefaults

`func NewNotifyWorkerTasksRequestWithDefaults() *NotifyWorkerTasksRequest`

NewNotifyWorkerTasksRequestWithDefaults instantiates a new NotifyWorkerTasksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShardId

`func (o *NotifyWorkerTasksRequest) GetShardId() int32`

GetShardId returns the ShardId field if non-nil, zero value otherwise.

### GetShardIdOk

`func (o *NotifyWorkerTasksRequest) GetShardIdOk() (*int32, bool)`

GetShardIdOk returns a tuple with the ShardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardId

`func (o *NotifyWorkerTasksRequest) SetShardId(v int32)`

SetShardId sets ShardId field to given value.


### GetNamespace

`func (o *NotifyWorkerTasksRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *NotifyWorkerTasksRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *NotifyWorkerTasksRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *NotifyWorkerTasksRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetProcessId

`func (o *NotifyWorkerTasksRequest) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *NotifyWorkerTasksRequest) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *NotifyWorkerTasksRequest) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *NotifyWorkerTasksRequest) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetProcessExecutionId

`func (o *NotifyWorkerTasksRequest) GetProcessExecutionId() string`

GetProcessExecutionId returns the ProcessExecutionId field if non-nil, zero value otherwise.

### GetProcessExecutionIdOk

`func (o *NotifyWorkerTasksRequest) GetProcessExecutionIdOk() (*string, bool)`

GetProcessExecutionIdOk returns a tuple with the ProcessExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessExecutionId

`func (o *NotifyWorkerTasksRequest) SetProcessExecutionId(v string)`

SetProcessExecutionId sets ProcessExecutionId field to given value.

### HasProcessExecutionId

`func (o *NotifyWorkerTasksRequest) HasProcessExecutionId() bool`

HasProcessExecutionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


