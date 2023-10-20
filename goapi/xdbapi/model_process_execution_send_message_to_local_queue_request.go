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

// checks if the ProcessExecutionSendMessageToLocalQueueRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessExecutionSendMessageToLocalQueueRequest{}

// ProcessExecutionSendMessageToLocalQueueRequest the request for sending message(s) to be consumed within a single process execution
type ProcessExecutionSendMessageToLocalQueueRequest struct {
	Namespace string `json:"namespace"`
	ProcessId string `json:"processId"`
	QueueName string `json:"queueName"`
	// UUID to uniquely distinguish different messages. If not specified, the server will generate a UUID instead.
	DedupId *string        `json:"dedupId,omitempty"`
	Payload *EncodedObject `json:"payload,omitempty"`
}

// NewProcessExecutionSendMessageToLocalQueueRequest instantiates a new ProcessExecutionSendMessageToLocalQueueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessExecutionSendMessageToLocalQueueRequest(namespace string, processId string, queueName string) *ProcessExecutionSendMessageToLocalQueueRequest {
	this := ProcessExecutionSendMessageToLocalQueueRequest{}
	this.Namespace = namespace
	this.ProcessId = processId
	this.QueueName = queueName
	return &this
}

// NewProcessExecutionSendMessageToLocalQueueRequestWithDefaults instantiates a new ProcessExecutionSendMessageToLocalQueueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessExecutionSendMessageToLocalQueueRequestWithDefaults() *ProcessExecutionSendMessageToLocalQueueRequest {
	this := ProcessExecutionSendMessageToLocalQueueRequest{}
	return &this
}

// GetNamespace returns the Namespace field value
func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *ProcessExecutionSendMessageToLocalQueueRequest) SetNamespace(v string) {
	o.Namespace = v
}

// GetProcessId returns the ProcessId field value
func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetProcessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value
// and a boolean to check if the value has been set.
func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetProcessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessId, true
}

// SetProcessId sets field value
func (o *ProcessExecutionSendMessageToLocalQueueRequest) SetProcessId(v string) {
	o.ProcessId = v
}

// GetQueueName returns the QueueName field value
func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetQueueName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueueName
}

// GetQueueNameOk returns a tuple with the QueueName field value
// and a boolean to check if the value has been set.
func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetQueueNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueueName, true
}

// SetQueueName sets field value
func (o *ProcessExecutionSendMessageToLocalQueueRequest) SetQueueName(v string) {
	o.QueueName = v
}

// GetDedupId returns the DedupId field value if set, zero value otherwise.
func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetDedupId() string {
	if o == nil || IsNil(o.DedupId) {
		var ret string
		return ret
	}
	return *o.DedupId
}

// GetDedupIdOk returns a tuple with the DedupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetDedupIdOk() (*string, bool) {
	if o == nil || IsNil(o.DedupId) {
		return nil, false
	}
	return o.DedupId, true
}

// HasDedupId returns a boolean if a field has been set.
func (o *ProcessExecutionSendMessageToLocalQueueRequest) HasDedupId() bool {
	if o != nil && !IsNil(o.DedupId) {
		return true
	}

	return false
}

// SetDedupId gets a reference to the given string and assigns it to the DedupId field.
func (o *ProcessExecutionSendMessageToLocalQueueRequest) SetDedupId(v string) {
	o.DedupId = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetPayload() EncodedObject {
	if o == nil || IsNil(o.Payload) {
		var ret EncodedObject
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessExecutionSendMessageToLocalQueueRequest) GetPayloadOk() (*EncodedObject, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ProcessExecutionSendMessageToLocalQueueRequest) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given EncodedObject and assigns it to the Payload field.
func (o *ProcessExecutionSendMessageToLocalQueueRequest) SetPayload(v EncodedObject) {
	o.Payload = &v
}

func (o ProcessExecutionSendMessageToLocalQueueRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessExecutionSendMessageToLocalQueueRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["namespace"] = o.Namespace
	toSerialize["processId"] = o.ProcessId
	toSerialize["queueName"] = o.QueueName
	if !IsNil(o.DedupId) {
		toSerialize["dedupId"] = o.DedupId
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	return toSerialize, nil
}

type NullableProcessExecutionSendMessageToLocalQueueRequest struct {
	value *ProcessExecutionSendMessageToLocalQueueRequest
	isSet bool
}

func (v NullableProcessExecutionSendMessageToLocalQueueRequest) Get() *ProcessExecutionSendMessageToLocalQueueRequest {
	return v.value
}

func (v *NullableProcessExecutionSendMessageToLocalQueueRequest) Set(val *ProcessExecutionSendMessageToLocalQueueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessExecutionSendMessageToLocalQueueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessExecutionSendMessageToLocalQueueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessExecutionSendMessageToLocalQueueRequest(val *ProcessExecutionSendMessageToLocalQueueRequest) *NullableProcessExecutionSendMessageToLocalQueueRequest {
	return &NullableProcessExecutionSendMessageToLocalQueueRequest{value: val, isSet: true}
}

func (v NullableProcessExecutionSendMessageToLocalQueueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessExecutionSendMessageToLocalQueueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}