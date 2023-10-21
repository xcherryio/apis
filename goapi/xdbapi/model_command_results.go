/*
XDB APIs

This APIs between xdb service and SDKs

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xdbapi

import (
	"encoding/json"
)

// checks if the CommandResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommandResults{}

// CommandResults struct for CommandResults
type CommandResults struct {
	TimerResults      []TimerResult       `json:"timerResults,omitempty"`
	LocalQueueResults []LocalQueueMessage `json:"localQueueResults,omitempty"`
}

// NewCommandResults instantiates a new CommandResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommandResults() *CommandResults {
	this := CommandResults{}
	return &this
}

// NewCommandResultsWithDefaults instantiates a new CommandResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandResultsWithDefaults() *CommandResults {
	this := CommandResults{}
	return &this
}

// GetTimerResults returns the TimerResults field value if set, zero value otherwise.
func (o *CommandResults) GetTimerResults() []TimerResult {
	if o == nil || IsNil(o.TimerResults) {
		var ret []TimerResult
		return ret
	}
	return o.TimerResults
}

// GetTimerResultsOk returns a tuple with the TimerResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandResults) GetTimerResultsOk() ([]TimerResult, bool) {
	if o == nil || IsNil(o.TimerResults) {
		return nil, false
	}
	return o.TimerResults, true
}

// HasTimerResults returns a boolean if a field has been set.
func (o *CommandResults) HasTimerResults() bool {
	if o != nil && !IsNil(o.TimerResults) {
		return true
	}

	return false
}

// SetTimerResults gets a reference to the given []TimerResult and assigns it to the TimerResults field.
func (o *CommandResults) SetTimerResults(v []TimerResult) {
	o.TimerResults = v
}

// GetLocalQueueResults returns the LocalQueueResults field value if set, zero value otherwise.
func (o *CommandResults) GetLocalQueueResults() []LocalQueueMessage {
	if o == nil || IsNil(o.LocalQueueResults) {
		var ret []LocalQueueMessage
		return ret
	}
	return o.LocalQueueResults
}

// GetLocalQueueResultsOk returns a tuple with the LocalQueueResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandResults) GetLocalQueueResultsOk() ([]LocalQueueMessage, bool) {
	if o == nil || IsNil(o.LocalQueueResults) {
		return nil, false
	}
	return o.LocalQueueResults, true
}

// HasLocalQueueResults returns a boolean if a field has been set.
func (o *CommandResults) HasLocalQueueResults() bool {
	if o != nil && !IsNil(o.LocalQueueResults) {
		return true
	}

	return false
}

// SetLocalQueueResults gets a reference to the given []LocalQueueMessage and assigns it to the LocalQueueResults field.
func (o *CommandResults) SetLocalQueueResults(v []LocalQueueMessage) {
	o.LocalQueueResults = v
}

func (o CommandResults) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommandResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TimerResults) {
		toSerialize["timerResults"] = o.TimerResults
	}
	if !IsNil(o.LocalQueueResults) {
		toSerialize["localQueueResults"] = o.LocalQueueResults
	}
	return toSerialize, nil
}

type NullableCommandResults struct {
	value *CommandResults
	isSet bool
}

func (v NullableCommandResults) Get() *CommandResults {
	return v.value
}

func (v *NullableCommandResults) Set(val *CommandResults) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandResults) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandResults(val *CommandResults) *NullableCommandResults {
	return &NullableCommandResults{value: val, isSet: true}
}

func (v NullableCommandResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
