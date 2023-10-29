# StateFailureRecoveryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | [**StateFailureRecoveryPolicy**](StateFailureRecoveryPolicy.md) |  | 
**StateFailureProceedStateId** | Pointer to **string** |  | [optional] 
**StateFailureProceedStateConfig** | Pointer to [**AsyncStateConfig**](AsyncStateConfig.md) |  | [optional] 

## Methods

### NewStateFailureRecoveryOptions

`func NewStateFailureRecoveryOptions(policy StateFailureRecoveryPolicy, ) *StateFailureRecoveryOptions`

NewStateFailureRecoveryOptions instantiates a new StateFailureRecoveryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateFailureRecoveryOptionsWithDefaults

`func NewStateFailureRecoveryOptionsWithDefaults() *StateFailureRecoveryOptions`

NewStateFailureRecoveryOptionsWithDefaults instantiates a new StateFailureRecoveryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *StateFailureRecoveryOptions) GetPolicy() StateFailureRecoveryPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *StateFailureRecoveryOptions) GetPolicyOk() (*StateFailureRecoveryPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *StateFailureRecoveryOptions) SetPolicy(v StateFailureRecoveryPolicy)`

SetPolicy sets Policy field to given value.


### GetStateFailureProceedStateId

`func (o *StateFailureRecoveryOptions) GetStateFailureProceedStateId() string`

GetStateFailureProceedStateId returns the StateFailureProceedStateId field if non-nil, zero value otherwise.

### GetStateFailureProceedStateIdOk

`func (o *StateFailureRecoveryOptions) GetStateFailureProceedStateIdOk() (*string, bool)`

GetStateFailureProceedStateIdOk returns a tuple with the StateFailureProceedStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateFailureProceedStateId

`func (o *StateFailureRecoveryOptions) SetStateFailureProceedStateId(v string)`

SetStateFailureProceedStateId sets StateFailureProceedStateId field to given value.

### HasStateFailureProceedStateId

`func (o *StateFailureRecoveryOptions) HasStateFailureProceedStateId() bool`

HasStateFailureProceedStateId returns a boolean if a field has been set.

### GetStateFailureProceedStateConfig

`func (o *StateFailureRecoveryOptions) GetStateFailureProceedStateConfig() AsyncStateConfig`

GetStateFailureProceedStateConfig returns the StateFailureProceedStateConfig field if non-nil, zero value otherwise.

### GetStateFailureProceedStateConfigOk

`func (o *StateFailureRecoveryOptions) GetStateFailureProceedStateConfigOk() (*AsyncStateConfig, bool)`

GetStateFailureProceedStateConfigOk returns a tuple with the StateFailureProceedStateConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateFailureProceedStateConfig

`func (o *StateFailureRecoveryOptions) SetStateFailureProceedStateConfig(v AsyncStateConfig)`

SetStateFailureProceedStateConfig sets StateFailureProceedStateConfig field to given value.

### HasStateFailureProceedStateConfig

`func (o *StateFailureRecoveryOptions) HasStateFailureProceedStateConfig() bool`

HasStateFailureProceedStateConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


