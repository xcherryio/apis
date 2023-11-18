# AsyncStateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkipWaitUntil** | Pointer to **bool** |  | [optional] 
**WaitUntilApiTimeoutSeconds** | Pointer to **int32** | the timeout for the single attempt of AsyncState.waitUntil API | [optional] 
**ExecuteApiTimeoutSeconds** | Pointer to **int32** | the timeout for the single attempt of AsyncState.execute API | [optional] 
**WaitUntilApiRetryPolicy** | Pointer to [**RetryPolicy**](RetryPolicy.md) |  | [optional] 
**ExecuteApiRetryPolicy** | Pointer to [**RetryPolicy**](RetryPolicy.md) |  | [optional] 
**StateFailureRecoveryOptions** | Pointer to [**StateFailureRecoveryOptions**](StateFailureRecoveryOptions.md) |  | [optional] 
**LoadGlobalAttributesRequest** | Pointer to [**LoadGlobalAttributesRequest**](LoadGlobalAttributesRequest.md) |  | [optional] 
**LoadLocalAttributesRequest** | Pointer to [**LoadLocalAttributesRequest**](LoadLocalAttributesRequest.md) |  | [optional] 

## Methods

### NewAsyncStateConfig

`func NewAsyncStateConfig() *AsyncStateConfig`

NewAsyncStateConfig instantiates a new AsyncStateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncStateConfigWithDefaults

`func NewAsyncStateConfigWithDefaults() *AsyncStateConfig`

NewAsyncStateConfigWithDefaults instantiates a new AsyncStateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipWaitUntil

`func (o *AsyncStateConfig) GetSkipWaitUntil() bool`

GetSkipWaitUntil returns the SkipWaitUntil field if non-nil, zero value otherwise.

### GetSkipWaitUntilOk

`func (o *AsyncStateConfig) GetSkipWaitUntilOk() (*bool, bool)`

GetSkipWaitUntilOk returns a tuple with the SkipWaitUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWaitUntil

`func (o *AsyncStateConfig) SetSkipWaitUntil(v bool)`

SetSkipWaitUntil sets SkipWaitUntil field to given value.

### HasSkipWaitUntil

`func (o *AsyncStateConfig) HasSkipWaitUntil() bool`

HasSkipWaitUntil returns a boolean if a field has been set.

### GetWaitUntilApiTimeoutSeconds

`func (o *AsyncStateConfig) GetWaitUntilApiTimeoutSeconds() int32`

GetWaitUntilApiTimeoutSeconds returns the WaitUntilApiTimeoutSeconds field if non-nil, zero value otherwise.

### GetWaitUntilApiTimeoutSecondsOk

`func (o *AsyncStateConfig) GetWaitUntilApiTimeoutSecondsOk() (*int32, bool)`

GetWaitUntilApiTimeoutSecondsOk returns a tuple with the WaitUntilApiTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitUntilApiTimeoutSeconds

`func (o *AsyncStateConfig) SetWaitUntilApiTimeoutSeconds(v int32)`

SetWaitUntilApiTimeoutSeconds sets WaitUntilApiTimeoutSeconds field to given value.

### HasWaitUntilApiTimeoutSeconds

`func (o *AsyncStateConfig) HasWaitUntilApiTimeoutSeconds() bool`

HasWaitUntilApiTimeoutSeconds returns a boolean if a field has been set.

### GetExecuteApiTimeoutSeconds

`func (o *AsyncStateConfig) GetExecuteApiTimeoutSeconds() int32`

GetExecuteApiTimeoutSeconds returns the ExecuteApiTimeoutSeconds field if non-nil, zero value otherwise.

### GetExecuteApiTimeoutSecondsOk

`func (o *AsyncStateConfig) GetExecuteApiTimeoutSecondsOk() (*int32, bool)`

GetExecuteApiTimeoutSecondsOk returns a tuple with the ExecuteApiTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteApiTimeoutSeconds

`func (o *AsyncStateConfig) SetExecuteApiTimeoutSeconds(v int32)`

SetExecuteApiTimeoutSeconds sets ExecuteApiTimeoutSeconds field to given value.

### HasExecuteApiTimeoutSeconds

`func (o *AsyncStateConfig) HasExecuteApiTimeoutSeconds() bool`

HasExecuteApiTimeoutSeconds returns a boolean if a field has been set.

### GetWaitUntilApiRetryPolicy

`func (o *AsyncStateConfig) GetWaitUntilApiRetryPolicy() RetryPolicy`

GetWaitUntilApiRetryPolicy returns the WaitUntilApiRetryPolicy field if non-nil, zero value otherwise.

### GetWaitUntilApiRetryPolicyOk

`func (o *AsyncStateConfig) GetWaitUntilApiRetryPolicyOk() (*RetryPolicy, bool)`

GetWaitUntilApiRetryPolicyOk returns a tuple with the WaitUntilApiRetryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitUntilApiRetryPolicy

`func (o *AsyncStateConfig) SetWaitUntilApiRetryPolicy(v RetryPolicy)`

SetWaitUntilApiRetryPolicy sets WaitUntilApiRetryPolicy field to given value.

### HasWaitUntilApiRetryPolicy

`func (o *AsyncStateConfig) HasWaitUntilApiRetryPolicy() bool`

HasWaitUntilApiRetryPolicy returns a boolean if a field has been set.

### GetExecuteApiRetryPolicy

`func (o *AsyncStateConfig) GetExecuteApiRetryPolicy() RetryPolicy`

GetExecuteApiRetryPolicy returns the ExecuteApiRetryPolicy field if non-nil, zero value otherwise.

### GetExecuteApiRetryPolicyOk

`func (o *AsyncStateConfig) GetExecuteApiRetryPolicyOk() (*RetryPolicy, bool)`

GetExecuteApiRetryPolicyOk returns a tuple with the ExecuteApiRetryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteApiRetryPolicy

`func (o *AsyncStateConfig) SetExecuteApiRetryPolicy(v RetryPolicy)`

SetExecuteApiRetryPolicy sets ExecuteApiRetryPolicy field to given value.

### HasExecuteApiRetryPolicy

`func (o *AsyncStateConfig) HasExecuteApiRetryPolicy() bool`

HasExecuteApiRetryPolicy returns a boolean if a field has been set.

### GetStateFailureRecoveryOptions

`func (o *AsyncStateConfig) GetStateFailureRecoveryOptions() StateFailureRecoveryOptions`

GetStateFailureRecoveryOptions returns the StateFailureRecoveryOptions field if non-nil, zero value otherwise.

### GetStateFailureRecoveryOptionsOk

`func (o *AsyncStateConfig) GetStateFailureRecoveryOptionsOk() (*StateFailureRecoveryOptions, bool)`

GetStateFailureRecoveryOptionsOk returns a tuple with the StateFailureRecoveryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateFailureRecoveryOptions

`func (o *AsyncStateConfig) SetStateFailureRecoveryOptions(v StateFailureRecoveryOptions)`

SetStateFailureRecoveryOptions sets StateFailureRecoveryOptions field to given value.

### HasStateFailureRecoveryOptions

`func (o *AsyncStateConfig) HasStateFailureRecoveryOptions() bool`

HasStateFailureRecoveryOptions returns a boolean if a field has been set.

### GetLoadGlobalAttributesRequest

`func (o *AsyncStateConfig) GetLoadGlobalAttributesRequest() LoadGlobalAttributesRequest`

GetLoadGlobalAttributesRequest returns the LoadGlobalAttributesRequest field if non-nil, zero value otherwise.

### GetLoadGlobalAttributesRequestOk

`func (o *AsyncStateConfig) GetLoadGlobalAttributesRequestOk() (*LoadGlobalAttributesRequest, bool)`

GetLoadGlobalAttributesRequestOk returns a tuple with the LoadGlobalAttributesRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadGlobalAttributesRequest

`func (o *AsyncStateConfig) SetLoadGlobalAttributesRequest(v LoadGlobalAttributesRequest)`

SetLoadGlobalAttributesRequest sets LoadGlobalAttributesRequest field to given value.

### HasLoadGlobalAttributesRequest

`func (o *AsyncStateConfig) HasLoadGlobalAttributesRequest() bool`

HasLoadGlobalAttributesRequest returns a boolean if a field has been set.

### GetLoadLocalAttributesRequest

`func (o *AsyncStateConfig) GetLoadLocalAttributesRequest() LoadLocalAttributesRequest`

GetLoadLocalAttributesRequest returns the LoadLocalAttributesRequest field if non-nil, zero value otherwise.

### GetLoadLocalAttributesRequestOk

`func (o *AsyncStateConfig) GetLoadLocalAttributesRequestOk() (*LoadLocalAttributesRequest, bool)`

GetLoadLocalAttributesRequestOk returns a tuple with the LoadLocalAttributesRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadLocalAttributesRequest

`func (o *AsyncStateConfig) SetLoadLocalAttributesRequest(v LoadLocalAttributesRequest)`

SetLoadLocalAttributesRequest sets LoadLocalAttributesRequest field to given value.

### HasLoadLocalAttributesRequest

`func (o *AsyncStateConfig) HasLoadLocalAttributesRequest() bool`

HasLoadLocalAttributesRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


