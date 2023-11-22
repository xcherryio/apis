/*
xCherry APIs

This APIs between xCherry service and SDKs

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xcapi

import (
	"encoding/json"
	"fmt"
)

// checks if the AppDatabaseTableWrite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppDatabaseTableWrite{}

// AppDatabaseTableWrite struct for AppDatabaseTableWrite
type AppDatabaseTableWrite struct {
	TableName string                `json:"tableName"`
	Rows      []AppDatabaseRowWrite `json:"rows,omitempty"`
}

type _AppDatabaseTableWrite AppDatabaseTableWrite

// NewAppDatabaseTableWrite instantiates a new AppDatabaseTableWrite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppDatabaseTableWrite(tableName string) *AppDatabaseTableWrite {
	this := AppDatabaseTableWrite{}
	this.TableName = tableName
	return &this
}

// NewAppDatabaseTableWriteWithDefaults instantiates a new AppDatabaseTableWrite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppDatabaseTableWriteWithDefaults() *AppDatabaseTableWrite {
	this := AppDatabaseTableWrite{}
	return &this
}

// GetTableName returns the TableName field value
func (o *AppDatabaseTableWrite) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *AppDatabaseTableWrite) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value
func (o *AppDatabaseTableWrite) SetTableName(v string) {
	o.TableName = v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *AppDatabaseTableWrite) GetRows() []AppDatabaseRowWrite {
	if o == nil || IsNil(o.Rows) {
		var ret []AppDatabaseRowWrite
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDatabaseTableWrite) GetRowsOk() ([]AppDatabaseRowWrite, bool) {
	if o == nil || IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *AppDatabaseTableWrite) HasRows() bool {
	if o != nil && !IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []AppDatabaseRowWrite and assigns it to the Rows field.
func (o *AppDatabaseTableWrite) SetRows(v []AppDatabaseRowWrite) {
	o.Rows = v
}

func (o AppDatabaseTableWrite) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppDatabaseTableWrite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tableName"] = o.TableName
	if !IsNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}
	return toSerialize, nil
}

func (o *AppDatabaseTableWrite) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tableName",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAppDatabaseTableWrite := _AppDatabaseTableWrite{}

	err = json.Unmarshal(bytes, &varAppDatabaseTableWrite)

	if err != nil {
		return err
	}

	*o = AppDatabaseTableWrite(varAppDatabaseTableWrite)

	return err
}

type NullableAppDatabaseTableWrite struct {
	value *AppDatabaseTableWrite
	isSet bool
}

func (v NullableAppDatabaseTableWrite) Get() *AppDatabaseTableWrite {
	return v.value
}

func (v *NullableAppDatabaseTableWrite) Set(val *AppDatabaseTableWrite) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDatabaseTableWrite) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDatabaseTableWrite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDatabaseTableWrite(val *AppDatabaseTableWrite) *NullableAppDatabaseTableWrite {
	return &NullableAppDatabaseTableWrite{value: val, isSet: true}
}

func (v NullableAppDatabaseTableWrite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDatabaseTableWrite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}