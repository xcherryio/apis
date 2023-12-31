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

// checks if the AsyncStateWaitUntilRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsyncStateWaitUntilRequest{}

// AsyncStateWaitUntilRequest the input of the waitUntil API
type AsyncStateWaitUntilRequest struct {
	Context     Context        `json:"context"`
	ProcessType string         `json:"processType"`
	StateId     string         `json:"stateId"`
	StateInput  *EncodedObject `json:"stateInput,omitempty"`
}

type _AsyncStateWaitUntilRequest AsyncStateWaitUntilRequest

// NewAsyncStateWaitUntilRequest instantiates a new AsyncStateWaitUntilRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsyncStateWaitUntilRequest(context Context, processType string, stateId string) *AsyncStateWaitUntilRequest {
	this := AsyncStateWaitUntilRequest{}
	this.Context = context
	this.ProcessType = processType
	this.StateId = stateId
	return &this
}

// NewAsyncStateWaitUntilRequestWithDefaults instantiates a new AsyncStateWaitUntilRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsyncStateWaitUntilRequestWithDefaults() *AsyncStateWaitUntilRequest {
	this := AsyncStateWaitUntilRequest{}
	return &this
}

// GetContext returns the Context field value
func (o *AsyncStateWaitUntilRequest) GetContext() Context {
	if o == nil {
		var ret Context
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *AsyncStateWaitUntilRequest) GetContextOk() (*Context, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *AsyncStateWaitUntilRequest) SetContext(v Context) {
	o.Context = v
}

// GetProcessType returns the ProcessType field value
func (o *AsyncStateWaitUntilRequest) GetProcessType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessType
}

// GetProcessTypeOk returns a tuple with the ProcessType field value
// and a boolean to check if the value has been set.
func (o *AsyncStateWaitUntilRequest) GetProcessTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessType, true
}

// SetProcessType sets field value
func (o *AsyncStateWaitUntilRequest) SetProcessType(v string) {
	o.ProcessType = v
}

// GetStateId returns the StateId field value
func (o *AsyncStateWaitUntilRequest) GetStateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateId
}

// GetStateIdOk returns a tuple with the StateId field value
// and a boolean to check if the value has been set.
func (o *AsyncStateWaitUntilRequest) GetStateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateId, true
}

// SetStateId sets field value
func (o *AsyncStateWaitUntilRequest) SetStateId(v string) {
	o.StateId = v
}

// GetStateInput returns the StateInput field value if set, zero value otherwise.
func (o *AsyncStateWaitUntilRequest) GetStateInput() EncodedObject {
	if o == nil || IsNil(o.StateInput) {
		var ret EncodedObject
		return ret
	}
	return *o.StateInput
}

// GetStateInputOk returns a tuple with the StateInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncStateWaitUntilRequest) GetStateInputOk() (*EncodedObject, bool) {
	if o == nil || IsNil(o.StateInput) {
		return nil, false
	}
	return o.StateInput, true
}

// HasStateInput returns a boolean if a field has been set.
func (o *AsyncStateWaitUntilRequest) HasStateInput() bool {
	if o != nil && !IsNil(o.StateInput) {
		return true
	}

	return false
}

// SetStateInput gets a reference to the given EncodedObject and assigns it to the StateInput field.
func (o *AsyncStateWaitUntilRequest) SetStateInput(v EncodedObject) {
	o.StateInput = &v
}

func (o AsyncStateWaitUntilRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AsyncStateWaitUntilRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["context"] = o.Context
	toSerialize["processType"] = o.ProcessType
	toSerialize["stateId"] = o.StateId
	if !IsNil(o.StateInput) {
		toSerialize["stateInput"] = o.StateInput
	}
	return toSerialize, nil
}

func (o *AsyncStateWaitUntilRequest) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"context",
		"processType",
		"stateId",
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

	varAsyncStateWaitUntilRequest := _AsyncStateWaitUntilRequest{}

	err = json.Unmarshal(bytes, &varAsyncStateWaitUntilRequest)

	if err != nil {
		return err
	}

	*o = AsyncStateWaitUntilRequest(varAsyncStateWaitUntilRequest)

	return err
}

type NullableAsyncStateWaitUntilRequest struct {
	value *AsyncStateWaitUntilRequest
	isSet bool
}

func (v NullableAsyncStateWaitUntilRequest) Get() *AsyncStateWaitUntilRequest {
	return v.value
}

func (v *NullableAsyncStateWaitUntilRequest) Set(val *AsyncStateWaitUntilRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAsyncStateWaitUntilRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAsyncStateWaitUntilRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsyncStateWaitUntilRequest(val *AsyncStateWaitUntilRequest) *NullableAsyncStateWaitUntilRequest {
	return &NullableAsyncStateWaitUntilRequest{value: val, isSet: true}
}

func (v NullableAsyncStateWaitUntilRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsyncStateWaitUntilRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
