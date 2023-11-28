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

// checks if the AppDatabaseErrorHandling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppDatabaseErrorHandling{}

// AppDatabaseErrorHandling handling the AppDatabase error
type AppDatabaseErrorHandling struct {
	LatestAppDBReqForRevise *AppDatabaseReadRequest `json:"latestAppDBReqForRevise,omitempty"`
}

// NewAppDatabaseErrorHandling instantiates a new AppDatabaseErrorHandling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppDatabaseErrorHandling() *AppDatabaseErrorHandling {
	this := AppDatabaseErrorHandling{}
	return &this
}

// NewAppDatabaseErrorHandlingWithDefaults instantiates a new AppDatabaseErrorHandling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppDatabaseErrorHandlingWithDefaults() *AppDatabaseErrorHandling {
	this := AppDatabaseErrorHandling{}
	return &this
}

// GetLatestAppDBReqForRevise returns the LatestAppDBReqForRevise field value if set, zero value otherwise.
func (o *AppDatabaseErrorHandling) GetLatestAppDBReqForRevise() AppDatabaseReadRequest {
	if o == nil || IsNil(o.LatestAppDBReqForRevise) {
		var ret AppDatabaseReadRequest
		return ret
	}
	return *o.LatestAppDBReqForRevise
}

// GetLatestAppDBReqForReviseOk returns a tuple with the LatestAppDBReqForRevise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDatabaseErrorHandling) GetLatestAppDBReqForReviseOk() (*AppDatabaseReadRequest, bool) {
	if o == nil || IsNil(o.LatestAppDBReqForRevise) {
		return nil, false
	}
	return o.LatestAppDBReqForRevise, true
}

// HasLatestAppDBReqForRevise returns a boolean if a field has been set.
func (o *AppDatabaseErrorHandling) HasLatestAppDBReqForRevise() bool {
	if o != nil && !IsNil(o.LatestAppDBReqForRevise) {
		return true
	}

	return false
}

// SetLatestAppDBReqForRevise gets a reference to the given AppDatabaseReadRequest and assigns it to the LatestAppDBReqForRevise field.
func (o *AppDatabaseErrorHandling) SetLatestAppDBReqForRevise(v AppDatabaseReadRequest) {
	o.LatestAppDBReqForRevise = &v
}

func (o AppDatabaseErrorHandling) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppDatabaseErrorHandling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LatestAppDBReqForRevise) {
		toSerialize["latestAppDBReqForRevise"] = o.LatestAppDBReqForRevise
	}
	return toSerialize, nil
}

type NullableAppDatabaseErrorHandling struct {
	value *AppDatabaseErrorHandling
	isSet bool
}

func (v NullableAppDatabaseErrorHandling) Get() *AppDatabaseErrorHandling {
	return v.value
}

func (v *NullableAppDatabaseErrorHandling) Set(val *AppDatabaseErrorHandling) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDatabaseErrorHandling) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDatabaseErrorHandling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDatabaseErrorHandling(val *AppDatabaseErrorHandling) *NullableAppDatabaseErrorHandling {
	return &NullableAppDatabaseErrorHandling{value: val, isSet: true}
}

func (v NullableAppDatabaseErrorHandling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDatabaseErrorHandling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
