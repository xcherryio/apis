# LoadGlobalAttributesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**[]GlobalAttributeKey**](GlobalAttributeKey.md) |  | [optional] 
**DefaultReadLockingType** | Pointer to [**AttributeReadLockingType**](AttributeReadLockingType.md) |  | [optional] 
**TableReadLockingPolicyOverrides** | Pointer to [**[]TableReadLockingPolicy**](TableReadLockingPolicy.md) | set a different read policy per table to override the default locking type | [optional] 

## Methods

### NewLoadGlobalAttributesRequest

`func NewLoadGlobalAttributesRequest() *LoadGlobalAttributesRequest`

NewLoadGlobalAttributesRequest instantiates a new LoadGlobalAttributesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadGlobalAttributesRequestWithDefaults

`func NewLoadGlobalAttributesRequestWithDefaults() *LoadGlobalAttributesRequest`

NewLoadGlobalAttributesRequestWithDefaults instantiates a new LoadGlobalAttributesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *LoadGlobalAttributesRequest) GetAttributes() []GlobalAttributeKey`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LoadGlobalAttributesRequest) GetAttributesOk() (*[]GlobalAttributeKey, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LoadGlobalAttributesRequest) SetAttributes(v []GlobalAttributeKey)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *LoadGlobalAttributesRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDefaultReadLockingType

`func (o *LoadGlobalAttributesRequest) GetDefaultReadLockingType() AttributeReadLockingType`

GetDefaultReadLockingType returns the DefaultReadLockingType field if non-nil, zero value otherwise.

### GetDefaultReadLockingTypeOk

`func (o *LoadGlobalAttributesRequest) GetDefaultReadLockingTypeOk() (*AttributeReadLockingType, bool)`

GetDefaultReadLockingTypeOk returns a tuple with the DefaultReadLockingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultReadLockingType

`func (o *LoadGlobalAttributesRequest) SetDefaultReadLockingType(v AttributeReadLockingType)`

SetDefaultReadLockingType sets DefaultReadLockingType field to given value.

### HasDefaultReadLockingType

`func (o *LoadGlobalAttributesRequest) HasDefaultReadLockingType() bool`

HasDefaultReadLockingType returns a boolean if a field has been set.

### GetTableReadLockingPolicyOverrides

`func (o *LoadGlobalAttributesRequest) GetTableReadLockingPolicyOverrides() []TableReadLockingPolicy`

GetTableReadLockingPolicyOverrides returns the TableReadLockingPolicyOverrides field if non-nil, zero value otherwise.

### GetTableReadLockingPolicyOverridesOk

`func (o *LoadGlobalAttributesRequest) GetTableReadLockingPolicyOverridesOk() (*[]TableReadLockingPolicy, bool)`

GetTableReadLockingPolicyOverridesOk returns a tuple with the TableReadLockingPolicyOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableReadLockingPolicyOverrides

`func (o *LoadGlobalAttributesRequest) SetTableReadLockingPolicyOverrides(v []TableReadLockingPolicy)`

SetTableReadLockingPolicyOverrides sets TableReadLockingPolicyOverrides field to given value.

### HasTableReadLockingPolicyOverrides

`func (o *LoadGlobalAttributesRequest) HasTableReadLockingPolicyOverrides() bool`

HasTableReadLockingPolicyOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


