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

// checks if the PublishToLocalQueueRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublishToLocalQueueRequest{}

// PublishToLocalQueueRequest the request for sending messages to be consumed within a single process execution
type PublishToLocalQueueRequest struct {
	Namespace string              `json:"namespace"`
	ProcessId string              `json:"processId"`
	Messages  []LocalQueueMessage `json:"messages,omitempty"`
}

// NewPublishToLocalQueueRequest instantiates a new PublishToLocalQueueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublishToLocalQueueRequest(namespace string, processId string) *PublishToLocalQueueRequest {
	this := PublishToLocalQueueRequest{}
	this.Namespace = namespace
	this.ProcessId = processId
	return &this
}

// NewPublishToLocalQueueRequestWithDefaults instantiates a new PublishToLocalQueueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublishToLocalQueueRequestWithDefaults() *PublishToLocalQueueRequest {
	this := PublishToLocalQueueRequest{}
	return &this
}

// GetNamespace returns the Namespace field value
func (o *PublishToLocalQueueRequest) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *PublishToLocalQueueRequest) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *PublishToLocalQueueRequest) SetNamespace(v string) {
	o.Namespace = v
}

// GetProcessId returns the ProcessId field value
func (o *PublishToLocalQueueRequest) GetProcessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value
// and a boolean to check if the value has been set.
func (o *PublishToLocalQueueRequest) GetProcessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessId, true
}

// SetProcessId sets field value
func (o *PublishToLocalQueueRequest) SetProcessId(v string) {
	o.ProcessId = v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *PublishToLocalQueueRequest) GetMessages() []LocalQueueMessage {
	if o == nil || IsNil(o.Messages) {
		var ret []LocalQueueMessage
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishToLocalQueueRequest) GetMessagesOk() ([]LocalQueueMessage, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *PublishToLocalQueueRequest) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []LocalQueueMessage and assigns it to the Messages field.
func (o *PublishToLocalQueueRequest) SetMessages(v []LocalQueueMessage) {
	o.Messages = v
}

func (o PublishToLocalQueueRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublishToLocalQueueRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["namespace"] = o.Namespace
	toSerialize["processId"] = o.ProcessId
	if !IsNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	return toSerialize, nil
}

type NullablePublishToLocalQueueRequest struct {
	value *PublishToLocalQueueRequest
	isSet bool
}

func (v NullablePublishToLocalQueueRequest) Get() *PublishToLocalQueueRequest {
	return v.value
}

func (v *NullablePublishToLocalQueueRequest) Set(val *PublishToLocalQueueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePublishToLocalQueueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePublishToLocalQueueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublishToLocalQueueRequest(val *PublishToLocalQueueRequest) *NullablePublishToLocalQueueRequest {
	return &NullablePublishToLocalQueueRequest{value: val, isSet: true}
}

func (v NullablePublishToLocalQueueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublishToLocalQueueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
