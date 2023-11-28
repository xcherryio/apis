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

// checks if the AppDatabaseTableReadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppDatabaseTableReadRequest{}

// AppDatabaseTableReadRequest struct for AppDatabaseTableReadRequest
type AppDatabaseTableReadRequest struct {
	TableName *string              `json:"tableName,omitempty"`
	LockType  *DatabaseLockingType `json:"lockType,omitempty"`
	Columns   []string             `json:"columns,omitempty"`
}

// NewAppDatabaseTableReadRequest instantiates a new AppDatabaseTableReadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppDatabaseTableReadRequest() *AppDatabaseTableReadRequest {
	this := AppDatabaseTableReadRequest{}
	return &this
}

// NewAppDatabaseTableReadRequestWithDefaults instantiates a new AppDatabaseTableReadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppDatabaseTableReadRequestWithDefaults() *AppDatabaseTableReadRequest {
	this := AppDatabaseTableReadRequest{}
	return &this
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *AppDatabaseTableReadRequest) GetTableName() string {
	if o == nil || IsNil(o.TableName) {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDatabaseTableReadRequest) GetTableNameOk() (*string, bool) {
	if o == nil || IsNil(o.TableName) {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *AppDatabaseTableReadRequest) HasTableName() bool {
	if o != nil && !IsNil(o.TableName) {
		return true
	}

	return false
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *AppDatabaseTableReadRequest) SetTableName(v string) {
	o.TableName = &v
}

// GetLockType returns the LockType field value if set, zero value otherwise.
func (o *AppDatabaseTableReadRequest) GetLockType() DatabaseLockingType {
	if o == nil || IsNil(o.LockType) {
		var ret DatabaseLockingType
		return ret
	}
	return *o.LockType
}

// GetLockTypeOk returns a tuple with the LockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDatabaseTableReadRequest) GetLockTypeOk() (*DatabaseLockingType, bool) {
	if o == nil || IsNil(o.LockType) {
		return nil, false
	}
	return o.LockType, true
}

// HasLockType returns a boolean if a field has been set.
func (o *AppDatabaseTableReadRequest) HasLockType() bool {
	if o != nil && !IsNil(o.LockType) {
		return true
	}

	return false
}

// SetLockType gets a reference to the given DatabaseLockingType and assigns it to the LockType field.
func (o *AppDatabaseTableReadRequest) SetLockType(v DatabaseLockingType) {
	o.LockType = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *AppDatabaseTableReadRequest) GetColumns() []string {
	if o == nil || IsNil(o.Columns) {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDatabaseTableReadRequest) GetColumnsOk() ([]string, bool) {
	if o == nil || IsNil(o.Columns) {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *AppDatabaseTableReadRequest) HasColumns() bool {
	if o != nil && !IsNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *AppDatabaseTableReadRequest) SetColumns(v []string) {
	o.Columns = v
}

func (o AppDatabaseTableReadRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppDatabaseTableReadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TableName) {
		toSerialize["tableName"] = o.TableName
	}
	if !IsNil(o.LockType) {
		toSerialize["lockType"] = o.LockType
	}
	if !IsNil(o.Columns) {
		toSerialize["columns"] = o.Columns
	}
	return toSerialize, nil
}

type NullableAppDatabaseTableReadRequest struct {
	value *AppDatabaseTableReadRequest
	isSet bool
}

func (v NullableAppDatabaseTableReadRequest) Get() *AppDatabaseTableReadRequest {
	return v.value
}

func (v *NullableAppDatabaseTableReadRequest) Set(val *AppDatabaseTableReadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDatabaseTableReadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDatabaseTableReadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDatabaseTableReadRequest(val *AppDatabaseTableReadRequest) *NullableAppDatabaseTableReadRequest {
	return &NullableAppDatabaseTableReadRequest{value: val, isSet: true}
}

func (v NullableAppDatabaseTableReadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDatabaseTableReadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
