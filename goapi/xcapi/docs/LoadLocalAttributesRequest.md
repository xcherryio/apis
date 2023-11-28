# LoadLocalAttributesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeysToLoadNoLock** | Pointer to **[]string** |  | [optional] 
**KeysToLoadWithLock** | Pointer to **[]string** |  | [optional] 
**LockType** | Pointer to [**DatabaseLockingType**](DatabaseLockingType.md) |  | [optional] 

## Methods

### NewLoadLocalAttributesRequest

`func NewLoadLocalAttributesRequest() *LoadLocalAttributesRequest`

NewLoadLocalAttributesRequest instantiates a new LoadLocalAttributesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadLocalAttributesRequestWithDefaults

`func NewLoadLocalAttributesRequestWithDefaults() *LoadLocalAttributesRequest`

NewLoadLocalAttributesRequestWithDefaults instantiates a new LoadLocalAttributesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeysToLoadNoLock

`func (o *LoadLocalAttributesRequest) GetKeysToLoadNoLock() []string`

GetKeysToLoadNoLock returns the KeysToLoadNoLock field if non-nil, zero value otherwise.

### GetKeysToLoadNoLockOk

`func (o *LoadLocalAttributesRequest) GetKeysToLoadNoLockOk() (*[]string, bool)`

GetKeysToLoadNoLockOk returns a tuple with the KeysToLoadNoLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeysToLoadNoLock

`func (o *LoadLocalAttributesRequest) SetKeysToLoadNoLock(v []string)`

SetKeysToLoadNoLock sets KeysToLoadNoLock field to given value.

### HasKeysToLoadNoLock

`func (o *LoadLocalAttributesRequest) HasKeysToLoadNoLock() bool`

HasKeysToLoadNoLock returns a boolean if a field has been set.

### GetKeysToLoadWithLock

`func (o *LoadLocalAttributesRequest) GetKeysToLoadWithLock() []string`

GetKeysToLoadWithLock returns the KeysToLoadWithLock field if non-nil, zero value otherwise.

### GetKeysToLoadWithLockOk

`func (o *LoadLocalAttributesRequest) GetKeysToLoadWithLockOk() (*[]string, bool)`

GetKeysToLoadWithLockOk returns a tuple with the KeysToLoadWithLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeysToLoadWithLock

`func (o *LoadLocalAttributesRequest) SetKeysToLoadWithLock(v []string)`

SetKeysToLoadWithLock sets KeysToLoadWithLock field to given value.

### HasKeysToLoadWithLock

`func (o *LoadLocalAttributesRequest) HasKeysToLoadWithLock() bool`

HasKeysToLoadWithLock returns a boolean if a field has been set.

### GetLockType

`func (o *LoadLocalAttributesRequest) GetLockType() DatabaseLockingType`

GetLockType returns the LockType field if non-nil, zero value otherwise.

### GetLockTypeOk

`func (o *LoadLocalAttributesRequest) GetLockTypeOk() (*DatabaseLockingType, bool)`

GetLockTypeOk returns a tuple with the LockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockType

`func (o *LoadLocalAttributesRequest) SetLockType(v DatabaseLockingType)`

SetLockType sets LockType field to given value.

### HasLockType

`func (o *LoadLocalAttributesRequest) HasLockType() bool`

HasLockType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


