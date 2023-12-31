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

// ErrorSubType the model 'ErrorSubType'
type ErrorSubType string

// List of ErrorSubType
const (
	UNCATEGORIZED_ERROR      ErrorSubType = "UNCATEGORIZED_ERROR"
	WORKER_EXECUTION_ERROR   ErrorSubType = "WORKER_EXECUTION_ERROR"
	APP_DATABASE_READ_ERROR  ErrorSubType = "APP_DATABASE_READ_ERROR"
	APP_DATABASE_WRITE_ERROR ErrorSubType = "APP_DATABASE_WRITE_ERROR"
	POLL_TIMEOUT_ERROR       ErrorSubType = "POLL_TIMEOUT_ERROR"
)

// All allowed values of ErrorSubType enum
var AllowedErrorSubTypeEnumValues = []ErrorSubType{
	"UNCATEGORIZED_ERROR",
	"WORKER_EXECUTION_ERROR",
	"APP_DATABASE_READ_ERROR",
	"APP_DATABASE_WRITE_ERROR",
	"POLL_TIMEOUT_ERROR",
}

func (v *ErrorSubType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ErrorSubType(value)
	for _, existing := range AllowedErrorSubTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ErrorSubType", value)
}

// NewErrorSubTypeFromValue returns a pointer to a valid ErrorSubType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewErrorSubTypeFromValue(v string) (*ErrorSubType, error) {
	ev := ErrorSubType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ErrorSubType: valid values are %v", v, AllowedErrorSubTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ErrorSubType) IsValid() bool {
	for _, existing := range AllowedErrorSubTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ErrorSubType value
func (v ErrorSubType) Ptr() *ErrorSubType {
	return &v
}

type NullableErrorSubType struct {
	value *ErrorSubType
	isSet bool
}

func (v NullableErrorSubType) Get() *ErrorSubType {
	return v.value
}

func (v *NullableErrorSubType) Set(val *ErrorSubType) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorSubType) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorSubType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorSubType(val *ErrorSubType) *NullableErrorSubType {
	return &NullableErrorSubType{value: val, isSet: true}
}

func (v NullableErrorSubType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorSubType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
