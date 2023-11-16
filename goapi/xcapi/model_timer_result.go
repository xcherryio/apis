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

// checks if the TimerResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimerResult{}

// TimerResult struct for TimerResult
type TimerResult struct {
	Status CommandStatus `json:"status"`
}

type _TimerResult TimerResult

// NewTimerResult instantiates a new TimerResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimerResult(status CommandStatus) *TimerResult {
	this := TimerResult{}
	this.Status = status
	return &this
}

// NewTimerResultWithDefaults instantiates a new TimerResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimerResultWithDefaults() *TimerResult {
	this := TimerResult{}
	return &this
}

// GetStatus returns the Status field value
func (o *TimerResult) GetStatus() CommandStatus {
	if o == nil {
		var ret CommandStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TimerResult) GetStatusOk() (*CommandStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TimerResult) SetStatus(v CommandStatus) {
	o.Status = v
}

func (o TimerResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimerResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *TimerResult) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varTimerResult := _TimerResult{}

	err = json.Unmarshal(bytes, &varTimerResult)

	if err != nil {
		return err
	}

	*o = TimerResult(varTimerResult)

	return err
}

type NullableTimerResult struct {
	value *TimerResult
	isSet bool
}

func (v NullableTimerResult) Get() *TimerResult {
	return v.value
}

func (v *NullableTimerResult) Set(val *TimerResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTimerResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTimerResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimerResult(val *TimerResult) *NullableTimerResult {
	return &NullableTimerResult{value: val, isSet: true}
}

func (v NullableTimerResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimerResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}