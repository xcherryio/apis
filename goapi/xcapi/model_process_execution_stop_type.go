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

// ProcessExecutionStopType the model 'ProcessExecutionStopType'
type ProcessExecutionStopType string

// List of ProcessExecutionStopType
const (
	TERMINATE ProcessExecutionStopType = "TERMINATE"
	FAIL      ProcessExecutionStopType = "FAIL"
)

// All allowed values of ProcessExecutionStopType enum
var AllowedProcessExecutionStopTypeEnumValues = []ProcessExecutionStopType{
	"TERMINATE",
	"FAIL",
}

func (v *ProcessExecutionStopType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProcessExecutionStopType(value)
	for _, existing := range AllowedProcessExecutionStopTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProcessExecutionStopType", value)
}

// NewProcessExecutionStopTypeFromValue returns a pointer to a valid ProcessExecutionStopType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProcessExecutionStopTypeFromValue(v string) (*ProcessExecutionStopType, error) {
	ev := ProcessExecutionStopType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProcessExecutionStopType: valid values are %v", v, AllowedProcessExecutionStopTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProcessExecutionStopType) IsValid() bool {
	for _, existing := range AllowedProcessExecutionStopTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProcessExecutionStopType value
func (v ProcessExecutionStopType) Ptr() *ProcessExecutionStopType {
	return &v
}

type NullableProcessExecutionStopType struct {
	value *ProcessExecutionStopType
	isSet bool
}

func (v NullableProcessExecutionStopType) Get() *ProcessExecutionStopType {
	return v.value
}

func (v *NullableProcessExecutionStopType) Set(val *ProcessExecutionStopType) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessExecutionStopType) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessExecutionStopType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessExecutionStopType(val *ProcessExecutionStopType) *NullableProcessExecutionStopType {
	return &NullableProcessExecutionStopType{value: val, isSet: true}
}

func (v NullableProcessExecutionStopType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessExecutionStopType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
