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

// checks if the AppDatabaseReadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppDatabaseReadRequest{}

// AppDatabaseReadRequest the request to read the selected rows of configured app database tables
type AppDatabaseReadRequest struct {
	Tables []AppDatabaseTableReadRequest `json:"tables,omitempty"`
}

// NewAppDatabaseReadRequest instantiates a new AppDatabaseReadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppDatabaseReadRequest() *AppDatabaseReadRequest {
	this := AppDatabaseReadRequest{}
	return &this
}

// NewAppDatabaseReadRequestWithDefaults instantiates a new AppDatabaseReadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppDatabaseReadRequestWithDefaults() *AppDatabaseReadRequest {
	this := AppDatabaseReadRequest{}
	return &this
}

// GetTables returns the Tables field value if set, zero value otherwise.
func (o *AppDatabaseReadRequest) GetTables() []AppDatabaseTableReadRequest {
	if o == nil || IsNil(o.Tables) {
		var ret []AppDatabaseTableReadRequest
		return ret
	}
	return o.Tables
}

// GetTablesOk returns a tuple with the Tables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDatabaseReadRequest) GetTablesOk() ([]AppDatabaseTableReadRequest, bool) {
	if o == nil || IsNil(o.Tables) {
		return nil, false
	}
	return o.Tables, true
}

// HasTables returns a boolean if a field has been set.
func (o *AppDatabaseReadRequest) HasTables() bool {
	if o != nil && !IsNil(o.Tables) {
		return true
	}

	return false
}

// SetTables gets a reference to the given []AppDatabaseTableReadRequest and assigns it to the Tables field.
func (o *AppDatabaseReadRequest) SetTables(v []AppDatabaseTableReadRequest) {
	o.Tables = v
}

func (o AppDatabaseReadRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppDatabaseReadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tables) {
		toSerialize["tables"] = o.Tables
	}
	return toSerialize, nil
}

type NullableAppDatabaseReadRequest struct {
	value *AppDatabaseReadRequest
	isSet bool
}

func (v NullableAppDatabaseReadRequest) Get() *AppDatabaseReadRequest {
	return v.value
}

func (v *NullableAppDatabaseReadRequest) Set(val *AppDatabaseReadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDatabaseReadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDatabaseReadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDatabaseReadRequest(val *AppDatabaseReadRequest) *NullableAppDatabaseReadRequest {
	return &NullableAppDatabaseReadRequest{value: val, isSet: true}
}

func (v NullableAppDatabaseReadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDatabaseReadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
