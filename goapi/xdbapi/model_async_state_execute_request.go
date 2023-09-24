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

// checks if the AsyncStateExecuteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsyncStateExecuteRequest{}

// AsyncStateExecuteRequest the input of the execute API
type AsyncStateExecuteRequest struct {
	Context     *Context       `json:"context,omitempty"`
	ProcessType *string        `json:"processType,omitempty"`
	StateId     *string        `json:"stateId,omitempty"`
	StateInput  *EncodedObject `json:"stateInput,omitempty"`
}

// NewAsyncStateExecuteRequest instantiates a new AsyncStateExecuteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsyncStateExecuteRequest() *AsyncStateExecuteRequest {
	this := AsyncStateExecuteRequest{}
	return &this
}

// NewAsyncStateExecuteRequestWithDefaults instantiates a new AsyncStateExecuteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsyncStateExecuteRequestWithDefaults() *AsyncStateExecuteRequest {
	this := AsyncStateExecuteRequest{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *AsyncStateExecuteRequest) GetContext() Context {
	if o == nil || IsNil(o.Context) {
		var ret Context
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncStateExecuteRequest) GetContextOk() (*Context, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *AsyncStateExecuteRequest) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given Context and assigns it to the Context field.
func (o *AsyncStateExecuteRequest) SetContext(v Context) {
	o.Context = &v
}

// GetProcessType returns the ProcessType field value if set, zero value otherwise.
func (o *AsyncStateExecuteRequest) GetProcessType() string {
	if o == nil || IsNil(o.ProcessType) {
		var ret string
		return ret
	}
	return *o.ProcessType
}

// GetProcessTypeOk returns a tuple with the ProcessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncStateExecuteRequest) GetProcessTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessType) {
		return nil, false
	}
	return o.ProcessType, true
}

// HasProcessType returns a boolean if a field has been set.
func (o *AsyncStateExecuteRequest) HasProcessType() bool {
	if o != nil && !IsNil(o.ProcessType) {
		return true
	}

	return false
}

// SetProcessType gets a reference to the given string and assigns it to the ProcessType field.
func (o *AsyncStateExecuteRequest) SetProcessType(v string) {
	o.ProcessType = &v
}

// GetStateId returns the StateId field value if set, zero value otherwise.
func (o *AsyncStateExecuteRequest) GetStateId() string {
	if o == nil || IsNil(o.StateId) {
		var ret string
		return ret
	}
	return *o.StateId
}

// GetStateIdOk returns a tuple with the StateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncStateExecuteRequest) GetStateIdOk() (*string, bool) {
	if o == nil || IsNil(o.StateId) {
		return nil, false
	}
	return o.StateId, true
}

// HasStateId returns a boolean if a field has been set.
func (o *AsyncStateExecuteRequest) HasStateId() bool {
	if o != nil && !IsNil(o.StateId) {
		return true
	}

	return false
}

// SetStateId gets a reference to the given string and assigns it to the StateId field.
func (o *AsyncStateExecuteRequest) SetStateId(v string) {
	o.StateId = &v
}

// GetStateInput returns the StateInput field value if set, zero value otherwise.
func (o *AsyncStateExecuteRequest) GetStateInput() EncodedObject {
	if o == nil || IsNil(o.StateInput) {
		var ret EncodedObject
		return ret
	}
	return *o.StateInput
}

// GetStateInputOk returns a tuple with the StateInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncStateExecuteRequest) GetStateInputOk() (*EncodedObject, bool) {
	if o == nil || IsNil(o.StateInput) {
		return nil, false
	}
	return o.StateInput, true
}

// HasStateInput returns a boolean if a field has been set.
func (o *AsyncStateExecuteRequest) HasStateInput() bool {
	if o != nil && !IsNil(o.StateInput) {
		return true
	}

	return false
}

// SetStateInput gets a reference to the given EncodedObject and assigns it to the StateInput field.
func (o *AsyncStateExecuteRequest) SetStateInput(v EncodedObject) {
	o.StateInput = &v
}

func (o AsyncStateExecuteRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AsyncStateExecuteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.ProcessType) {
		toSerialize["processType"] = o.ProcessType
	}
	if !IsNil(o.StateId) {
		toSerialize["stateId"] = o.StateId
	}
	if !IsNil(o.StateInput) {
		toSerialize["stateInput"] = o.StateInput
	}
	return toSerialize, nil
}

type NullableAsyncStateExecuteRequest struct {
	value *AsyncStateExecuteRequest
	isSet bool
}

func (v NullableAsyncStateExecuteRequest) Get() *AsyncStateExecuteRequest {
	return v.value
}

func (v *NullableAsyncStateExecuteRequest) Set(val *AsyncStateExecuteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAsyncStateExecuteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAsyncStateExecuteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsyncStateExecuteRequest(val *AsyncStateExecuteRequest) *NullableAsyncStateExecuteRequest {
	return &NullableAsyncStateExecuteRequest{value: val, isSet: true}
}

func (v NullableAsyncStateExecuteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsyncStateExecuteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
