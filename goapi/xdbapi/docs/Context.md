# Context

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessId** | **string** |  | 
**ProcessExecutionId** | **string** |  | 
**ProcessStartedTimestamp** | **int64** |  | 
**StateExecutionId** | Pointer to **string** | stateExecutionId is for async state API only | [optional] 
**FirstAttemptTimestamp** | Pointer to **int64** | for async state API only(during backoff retry) | [optional] 
**Attempt** | Pointer to **int32** | for async state API only(during backoff retry) | [optional] 
**RecoverFromStateExecutionId** | Pointer to **string** | for async state API only, state id + sequence number | [optional] 
**RecoverFromApi** | Pointer to [**StateApiType**](StateApiType.md) |  | [optional] 

## Methods

### NewContext

`func NewContext(processId string, processExecutionId string, processStartedTimestamp int64, ) *Context`

NewContext instantiates a new Context object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextWithDefaults

`func NewContextWithDefaults() *Context`

NewContextWithDefaults instantiates a new Context object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessId

`func (o *Context) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *Context) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *Context) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.


### GetProcessExecutionId

`func (o *Context) GetProcessExecutionId() string`

GetProcessExecutionId returns the ProcessExecutionId field if non-nil, zero value otherwise.

### GetProcessExecutionIdOk

`func (o *Context) GetProcessExecutionIdOk() (*string, bool)`

GetProcessExecutionIdOk returns a tuple with the ProcessExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessExecutionId

`func (o *Context) SetProcessExecutionId(v string)`

SetProcessExecutionId sets ProcessExecutionId field to given value.


### GetProcessStartedTimestamp

`func (o *Context) GetProcessStartedTimestamp() int64`

GetProcessStartedTimestamp returns the ProcessStartedTimestamp field if non-nil, zero value otherwise.

### GetProcessStartedTimestampOk

`func (o *Context) GetProcessStartedTimestampOk() (*int64, bool)`

GetProcessStartedTimestampOk returns a tuple with the ProcessStartedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessStartedTimestamp

`func (o *Context) SetProcessStartedTimestamp(v int64)`

SetProcessStartedTimestamp sets ProcessStartedTimestamp field to given value.


### GetStateExecutionId

`func (o *Context) GetStateExecutionId() string`

GetStateExecutionId returns the StateExecutionId field if non-nil, zero value otherwise.

### GetStateExecutionIdOk

`func (o *Context) GetStateExecutionIdOk() (*string, bool)`

GetStateExecutionIdOk returns a tuple with the StateExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateExecutionId

`func (o *Context) SetStateExecutionId(v string)`

SetStateExecutionId sets StateExecutionId field to given value.

### HasStateExecutionId

`func (o *Context) HasStateExecutionId() bool`

HasStateExecutionId returns a boolean if a field has been set.

### GetFirstAttemptTimestamp

`func (o *Context) GetFirstAttemptTimestamp() int64`

GetFirstAttemptTimestamp returns the FirstAttemptTimestamp field if non-nil, zero value otherwise.

### GetFirstAttemptTimestampOk

`func (o *Context) GetFirstAttemptTimestampOk() (*int64, bool)`

GetFirstAttemptTimestampOk returns a tuple with the FirstAttemptTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAttemptTimestamp

`func (o *Context) SetFirstAttemptTimestamp(v int64)`

SetFirstAttemptTimestamp sets FirstAttemptTimestamp field to given value.

### HasFirstAttemptTimestamp

`func (o *Context) HasFirstAttemptTimestamp() bool`

HasFirstAttemptTimestamp returns a boolean if a field has been set.

### GetAttempt

`func (o *Context) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *Context) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *Context) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *Context) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetRecoverFromStateExecutionId

`func (o *Context) GetRecoverFromStateExecutionId() string`

GetRecoverFromStateExecutionId returns the RecoverFromStateExecutionId field if non-nil, zero value otherwise.

### GetRecoverFromStateExecutionIdOk

`func (o *Context) GetRecoverFromStateExecutionIdOk() (*string, bool)`

GetRecoverFromStateExecutionIdOk returns a tuple with the RecoverFromStateExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverFromStateExecutionId

`func (o *Context) SetRecoverFromStateExecutionId(v string)`

SetRecoverFromStateExecutionId sets RecoverFromStateExecutionId field to given value.

### HasRecoverFromStateExecutionId

`func (o *Context) HasRecoverFromStateExecutionId() bool`

HasRecoverFromStateExecutionId returns a boolean if a field has been set.

### GetRecoverFromApi

`func (o *Context) GetRecoverFromApi() StateApiType`

GetRecoverFromApi returns the RecoverFromApi field if non-nil, zero value otherwise.

### GetRecoverFromApiOk

`func (o *Context) GetRecoverFromApiOk() (*StateApiType, bool)`

GetRecoverFromApiOk returns a tuple with the RecoverFromApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverFromApi

`func (o *Context) SetRecoverFromApi(v StateApiType)`

SetRecoverFromApi sets RecoverFromApi field to given value.

### HasRecoverFromApi

`func (o *Context) HasRecoverFromApi() bool`

HasRecoverFromApi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


