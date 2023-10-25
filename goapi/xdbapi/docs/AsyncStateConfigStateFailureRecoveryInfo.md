# AsyncStateConfigStateFailureRecoveryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | [**StateFailureRecoveryPolicy**](StateFailureRecoveryPolicy.md) |  | 
**StateFailureProceedStateId** | Pointer to **string** |  | [optional] 
**StateFailureProceedStateConfig** | Pointer to [**AsyncStateConfig**](AsyncStateConfig.md) |  | [optional] 

## Methods

### NewAsyncStateConfigStateFailureRecoveryInfo

`func NewAsyncStateConfigStateFailureRecoveryInfo(policy StateFailureRecoveryPolicy, ) *AsyncStateConfigStateFailureRecoveryInfo`

NewAsyncStateConfigStateFailureRecoveryInfo instantiates a new AsyncStateConfigStateFailureRecoveryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncStateConfigStateFailureRecoveryInfoWithDefaults

`func NewAsyncStateConfigStateFailureRecoveryInfoWithDefaults() *AsyncStateConfigStateFailureRecoveryInfo`

NewAsyncStateConfigStateFailureRecoveryInfoWithDefaults instantiates a new AsyncStateConfigStateFailureRecoveryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *AsyncStateConfigStateFailureRecoveryInfo) GetPolicy() StateFailureRecoveryPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *AsyncStateConfigStateFailureRecoveryInfo) GetPolicyOk() (*StateFailureRecoveryPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *AsyncStateConfigStateFailureRecoveryInfo) SetPolicy(v StateFailureRecoveryPolicy)`

SetPolicy sets Policy field to given value.


### GetStateFailureProceedStateId

`func (o *AsyncStateConfigStateFailureRecoveryInfo) GetStateFailureProceedStateId() string`

GetStateFailureProceedStateId returns the StateFailureProceedStateId field if non-nil, zero value otherwise.

### GetStateFailureProceedStateIdOk

`func (o *AsyncStateConfigStateFailureRecoveryInfo) GetStateFailureProceedStateIdOk() (*string, bool)`

GetStateFailureProceedStateIdOk returns a tuple with the StateFailureProceedStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateFailureProceedStateId

`func (o *AsyncStateConfigStateFailureRecoveryInfo) SetStateFailureProceedStateId(v string)`

SetStateFailureProceedStateId sets StateFailureProceedStateId field to given value.

### HasStateFailureProceedStateId

`func (o *AsyncStateConfigStateFailureRecoveryInfo) HasStateFailureProceedStateId() bool`

HasStateFailureProceedStateId returns a boolean if a field has been set.

### GetStateFailureProceedStateConfig

`func (o *AsyncStateConfigStateFailureRecoveryInfo) GetStateFailureProceedStateConfig() AsyncStateConfig`

GetStateFailureProceedStateConfig returns the StateFailureProceedStateConfig field if non-nil, zero value otherwise.

### GetStateFailureProceedStateConfigOk

`func (o *AsyncStateConfigStateFailureRecoveryInfo) GetStateFailureProceedStateConfigOk() (*AsyncStateConfig, bool)`

GetStateFailureProceedStateConfigOk returns a tuple with the StateFailureProceedStateConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateFailureProceedStateConfig

`func (o *AsyncStateConfigStateFailureRecoveryInfo) SetStateFailureProceedStateConfig(v AsyncStateConfig)`

SetStateFailureProceedStateConfig sets StateFailureProceedStateConfig field to given value.

### HasStateFailureProceedStateConfig

`func (o *AsyncStateConfigStateFailureRecoveryInfo) HasStateFailureProceedStateConfig() bool`

HasStateFailureProceedStateConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


