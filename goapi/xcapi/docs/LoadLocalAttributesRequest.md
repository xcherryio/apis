# LoadLocalAttributesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeysToLoadNoLock** | Pointer to **[]string** |  | [optional] 
**KeysToLockWithLock** | Pointer to **[]string** |  | [optional] 
**LockingPolicy** | Pointer to [**TableReadLockingPolicy**](TableReadLockingPolicy.md) |  | [optional] 

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

### GetKeysToLockWithLock

`func (o *LoadLocalAttributesRequest) GetKeysToLockWithLock() []string`

GetKeysToLockWithLock returns the KeysToLockWithLock field if non-nil, zero value otherwise.

### GetKeysToLockWithLockOk

`func (o *LoadLocalAttributesRequest) GetKeysToLockWithLockOk() (*[]string, bool)`

GetKeysToLockWithLockOk returns a tuple with the KeysToLockWithLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeysToLockWithLock

`func (o *LoadLocalAttributesRequest) SetKeysToLockWithLock(v []string)`

SetKeysToLockWithLock sets KeysToLockWithLock field to given value.

### HasKeysToLockWithLock

`func (o *LoadLocalAttributesRequest) HasKeysToLockWithLock() bool`

HasKeysToLockWithLock returns a boolean if a field has been set.

### GetLockingPolicy

`func (o *LoadLocalAttributesRequest) GetLockingPolicy() TableReadLockingPolicy`

GetLockingPolicy returns the LockingPolicy field if non-nil, zero value otherwise.

### GetLockingPolicyOk

`func (o *LoadLocalAttributesRequest) GetLockingPolicyOk() (*TableReadLockingPolicy, bool)`

GetLockingPolicyOk returns a tuple with the LockingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockingPolicy

`func (o *LoadLocalAttributesRequest) SetLockingPolicy(v TableReadLockingPolicy)`

SetLockingPolicy sets LockingPolicy field to given value.

### HasLockingPolicy

`func (o *LoadLocalAttributesRequest) HasLockingPolicy() bool`

HasLockingPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


