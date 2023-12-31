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

// checks if the AppDatabaseRowReadResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppDatabaseRowReadResponse{}

// AppDatabaseRowReadResponse struct for AppDatabaseRowReadResponse
type AppDatabaseRowReadResponse struct {
	Columns []AppDatabaseColumnValue `json:"columns,omitempty"`
}

// NewAppDatabaseRowReadResponse instantiates a new AppDatabaseRowReadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppDatabaseRowReadResponse() *AppDatabaseRowReadResponse {
	this := AppDatabaseRowReadResponse{}
	return &this
}

// NewAppDatabaseRowReadResponseWithDefaults instantiates a new AppDatabaseRowReadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppDatabaseRowReadResponseWithDefaults() *AppDatabaseRowReadResponse {
	this := AppDatabaseRowReadResponse{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *AppDatabaseRowReadResponse) GetColumns() []AppDatabaseColumnValue {
	if o == nil || IsNil(o.Columns) {
		var ret []AppDatabaseColumnValue
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDatabaseRowReadResponse) GetColumnsOk() ([]AppDatabaseColumnValue, bool) {
	if o == nil || IsNil(o.Columns) {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *AppDatabaseRowReadResponse) HasColumns() bool {
	if o != nil && !IsNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []AppDatabaseColumnValue and assigns it to the Columns field.
func (o *AppDatabaseRowReadResponse) SetColumns(v []AppDatabaseColumnValue) {
	o.Columns = v
}

func (o AppDatabaseRowReadResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppDatabaseRowReadResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Columns) {
		toSerialize["columns"] = o.Columns
	}
	return toSerialize, nil
}

type NullableAppDatabaseRowReadResponse struct {
	value *AppDatabaseRowReadResponse
	isSet bool
}

func (v NullableAppDatabaseRowReadResponse) Get() *AppDatabaseRowReadResponse {
	return v.value
}

func (v *NullableAppDatabaseRowReadResponse) Set(val *AppDatabaseRowReadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDatabaseRowReadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDatabaseRowReadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDatabaseRowReadResponse(val *AppDatabaseRowReadResponse) *NullableAppDatabaseRowReadResponse {
	return &NullableAppDatabaseRowReadResponse{value: val, isSet: true}
}

func (v NullableAppDatabaseRowReadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDatabaseRowReadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
