/*
xCherry APIs

This APIs between xCherry service and SDKs

API version: 0.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xcapi

import (
	"encoding/json"
)

// checks if the LoadLocalAttributesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadLocalAttributesRequest{}

// LoadLocalAttributesRequest struct for LoadLocalAttributesRequest
type LoadLocalAttributesRequest struct {
	KeysToLoadNoLock   []string  `json:"keysToLoadNoLock,omitempty"`
	KeysToLoadWithLock []string  `json:"keysToLoadWithLock,omitempty"`
	LockType           *LockType `json:"lockType,omitempty"`
}

// NewLoadLocalAttributesRequest instantiates a new LoadLocalAttributesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadLocalAttributesRequest() *LoadLocalAttributesRequest {
	this := LoadLocalAttributesRequest{}
	return &this
}

// NewLoadLocalAttributesRequestWithDefaults instantiates a new LoadLocalAttributesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadLocalAttributesRequestWithDefaults() *LoadLocalAttributesRequest {
	this := LoadLocalAttributesRequest{}
	return &this
}

// GetKeysToLoadNoLock returns the KeysToLoadNoLock field value if set, zero value otherwise.
func (o *LoadLocalAttributesRequest) GetKeysToLoadNoLock() []string {
	if o == nil || IsNil(o.KeysToLoadNoLock) {
		var ret []string
		return ret
	}
	return o.KeysToLoadNoLock
}

// GetKeysToLoadNoLockOk returns a tuple with the KeysToLoadNoLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadLocalAttributesRequest) GetKeysToLoadNoLockOk() ([]string, bool) {
	if o == nil || IsNil(o.KeysToLoadNoLock) {
		return nil, false
	}
	return o.KeysToLoadNoLock, true
}

// HasKeysToLoadNoLock returns a boolean if a field has been set.
func (o *LoadLocalAttributesRequest) HasKeysToLoadNoLock() bool {
	if o != nil && !IsNil(o.KeysToLoadNoLock) {
		return true
	}

	return false
}

// SetKeysToLoadNoLock gets a reference to the given []string and assigns it to the KeysToLoadNoLock field.
func (o *LoadLocalAttributesRequest) SetKeysToLoadNoLock(v []string) {
	o.KeysToLoadNoLock = v
}

// GetKeysToLoadWithLock returns the KeysToLoadWithLock field value if set, zero value otherwise.
func (o *LoadLocalAttributesRequest) GetKeysToLoadWithLock() []string {
	if o == nil || IsNil(o.KeysToLoadWithLock) {
		var ret []string
		return ret
	}
	return o.KeysToLoadWithLock
}

// GetKeysToLoadWithLockOk returns a tuple with the KeysToLoadWithLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadLocalAttributesRequest) GetKeysToLoadWithLockOk() ([]string, bool) {
	if o == nil || IsNil(o.KeysToLoadWithLock) {
		return nil, false
	}
	return o.KeysToLoadWithLock, true
}

// HasKeysToLoadWithLock returns a boolean if a field has been set.
func (o *LoadLocalAttributesRequest) HasKeysToLoadWithLock() bool {
	if o != nil && !IsNil(o.KeysToLoadWithLock) {
		return true
	}

	return false
}

// SetKeysToLoadWithLock gets a reference to the given []string and assigns it to the KeysToLoadWithLock field.
func (o *LoadLocalAttributesRequest) SetKeysToLoadWithLock(v []string) {
	o.KeysToLoadWithLock = v
}

// GetLockType returns the LockType field value if set, zero value otherwise.
func (o *LoadLocalAttributesRequest) GetLockType() LockType {
	if o == nil || IsNil(o.LockType) {
		var ret LockType
		return ret
	}
	return *o.LockType
}

// GetLockTypeOk returns a tuple with the LockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadLocalAttributesRequest) GetLockTypeOk() (*LockType, bool) {
	if o == nil || IsNil(o.LockType) {
		return nil, false
	}
	return o.LockType, true
}

// HasLockType returns a boolean if a field has been set.
func (o *LoadLocalAttributesRequest) HasLockType() bool {
	if o != nil && !IsNil(o.LockType) {
		return true
	}

	return false
}

// SetLockType gets a reference to the given LockType and assigns it to the LockType field.
func (o *LoadLocalAttributesRequest) SetLockType(v LockType) {
	o.LockType = &v
}

func (o LoadLocalAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadLocalAttributesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KeysToLoadNoLock) {
		toSerialize["keysToLoadNoLock"] = o.KeysToLoadNoLock
	}
	if !IsNil(o.KeysToLoadWithLock) {
		toSerialize["keysToLoadWithLock"] = o.KeysToLoadWithLock
	}
	if !IsNil(o.LockType) {
		toSerialize["lockType"] = o.LockType
	}
	return toSerialize, nil
}

type NullableLoadLocalAttributesRequest struct {
	value *LoadLocalAttributesRequest
	isSet bool
}

func (v NullableLoadLocalAttributesRequest) Get() *LoadLocalAttributesRequest {
	return v.value
}

func (v *NullableLoadLocalAttributesRequest) Set(val *LoadLocalAttributesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadLocalAttributesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadLocalAttributesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadLocalAttributesRequest(val *LoadLocalAttributesRequest) *NullableLoadLocalAttributesRequest {
	return &NullableLoadLocalAttributesRequest{value: val, isSet: true}
}

func (v NullableLoadLocalAttributesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadLocalAttributesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
