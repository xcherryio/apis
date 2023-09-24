# StateMovement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StateId** | **string** |  | 
**StateInput** | Pointer to [**EncodedObject**](EncodedObject.md) |  | [optional] 
**StateConfig** | Pointer to [**AsyncStateConfig**](AsyncStateConfig.md) |  | [optional] 

## Methods

### NewStateMovement

`func NewStateMovement(stateId string, ) *StateMovement`

NewStateMovement instantiates a new StateMovement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateMovementWithDefaults

`func NewStateMovementWithDefaults() *StateMovement`

NewStateMovementWithDefaults instantiates a new StateMovement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStateId

`func (o *StateMovement) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *StateMovement) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *StateMovement) SetStateId(v string)`

SetStateId sets StateId field to given value.


### GetStateInput

`func (o *StateMovement) GetStateInput() EncodedObject`

GetStateInput returns the StateInput field if non-nil, zero value otherwise.

### GetStateInputOk

`func (o *StateMovement) GetStateInputOk() (*EncodedObject, bool)`

GetStateInputOk returns a tuple with the StateInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateInput

`func (o *StateMovement) SetStateInput(v EncodedObject)`

SetStateInput sets StateInput field to given value.

### HasStateInput

`func (o *StateMovement) HasStateInput() bool`

HasStateInput returns a boolean if a field has been set.

### GetStateConfig

`func (o *StateMovement) GetStateConfig() AsyncStateConfig`

GetStateConfig returns the StateConfig field if non-nil, zero value otherwise.

### GetStateConfigOk

`func (o *StateMovement) GetStateConfigOk() (*AsyncStateConfig, bool)`

GetStateConfigOk returns a tuple with the StateConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateConfig

`func (o *StateMovement) SetStateConfig(v AsyncStateConfig)`

SetStateConfig sets StateConfig field to given value.

### HasStateConfig

`func (o *StateMovement) HasStateConfig() bool`

HasStateConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


