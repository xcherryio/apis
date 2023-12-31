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

// StateFailureRecoveryPolicy the model 'StateFailureRecoveryPolicy'
type StateFailureRecoveryPolicy string

// List of StateFailureRecoveryPolicy
const (
	FAIL_PROCESS_ON_STATE_FAILURE StateFailureRecoveryPolicy = "FAIL_PROCESS_ON_STATE_FAILURE"
	PROCEED_TO_CONFIGURED_STATE   StateFailureRecoveryPolicy = "PROCEED_TO_CONFIGURED_STATE"
)

// All allowed values of StateFailureRecoveryPolicy enum
var AllowedStateFailureRecoveryPolicyEnumValues = []StateFailureRecoveryPolicy{
	"FAIL_PROCESS_ON_STATE_FAILURE",
	"PROCEED_TO_CONFIGURED_STATE",
}

func (v *StateFailureRecoveryPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StateFailureRecoveryPolicy(value)
	for _, existing := range AllowedStateFailureRecoveryPolicyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StateFailureRecoveryPolicy", value)
}

// NewStateFailureRecoveryPolicyFromValue returns a pointer to a valid StateFailureRecoveryPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStateFailureRecoveryPolicyFromValue(v string) (*StateFailureRecoveryPolicy, error) {
	ev := StateFailureRecoveryPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StateFailureRecoveryPolicy: valid values are %v", v, AllowedStateFailureRecoveryPolicyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StateFailureRecoveryPolicy) IsValid() bool {
	for _, existing := range AllowedStateFailureRecoveryPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StateFailureRecoveryPolicy value
func (v StateFailureRecoveryPolicy) Ptr() *StateFailureRecoveryPolicy {
	return &v
}

type NullableStateFailureRecoveryPolicy struct {
	value *StateFailureRecoveryPolicy
	isSet bool
}

func (v NullableStateFailureRecoveryPolicy) Get() *StateFailureRecoveryPolicy {
	return v.value
}

func (v *NullableStateFailureRecoveryPolicy) Set(val *StateFailureRecoveryPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableStateFailureRecoveryPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableStateFailureRecoveryPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateFailureRecoveryPolicy(val *StateFailureRecoveryPolicy) *NullableStateFailureRecoveryPolicy {
	return &NullableStateFailureRecoveryPolicy{value: val, isSet: true}
}

func (v NullableStateFailureRecoveryPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateFailureRecoveryPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
