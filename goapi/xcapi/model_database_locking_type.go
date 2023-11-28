/*
xCherry APIs

This APIs between xCherry service and SDKs

API version: 0.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xcapi

import (
	"encoding/json"
	"fmt"
)

// DatabaseLockingType the model 'DatabaseLockingType'
type DatabaseLockingType string

// List of DatabaseLockingType
const (
	NO_LOCKING     DatabaseLockingType = "NO_LOCKING"
	SHARE_LOCK     DatabaseLockingType = "SHARE_LOCK"
	EXCLUSIVE_LOCK DatabaseLockingType = "EXCLUSIVE_LOCK"
)

// All allowed values of DatabaseLockingType enum
var AllowedDatabaseLockingTypeEnumValues = []DatabaseLockingType{
	"NO_LOCKING",
	"SHARE_LOCK",
	"EXCLUSIVE_LOCK",
}

func (v *DatabaseLockingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DatabaseLockingType(value)
	for _, existing := range AllowedDatabaseLockingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DatabaseLockingType", value)
}

// NewDatabaseLockingTypeFromValue returns a pointer to a valid DatabaseLockingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDatabaseLockingTypeFromValue(v string) (*DatabaseLockingType, error) {
	ev := DatabaseLockingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DatabaseLockingType: valid values are %v", v, AllowedDatabaseLockingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DatabaseLockingType) IsValid() bool {
	for _, existing := range AllowedDatabaseLockingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatabaseLockingType value
func (v DatabaseLockingType) Ptr() *DatabaseLockingType {
	return &v
}

type NullableDatabaseLockingType struct {
	value *DatabaseLockingType
	isSet bool
}

func (v NullableDatabaseLockingType) Get() *DatabaseLockingType {
	return v.value
}

func (v *NullableDatabaseLockingType) Set(val *DatabaseLockingType) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseLockingType) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseLockingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseLockingType(val *DatabaseLockingType) *NullableDatabaseLockingType {
	return &NullableDatabaseLockingType{value: val, isSet: true}
}

func (v NullableDatabaseLockingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseLockingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
