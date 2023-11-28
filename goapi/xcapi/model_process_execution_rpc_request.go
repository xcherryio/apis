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

// checks if the ProcessExecutionRpcRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessExecutionRpcRequest{}

// ProcessExecutionRpcRequest the request for executing a RPC method of a process execution
type ProcessExecutionRpcRequest struct {
	Namespace string         `json:"namespace"`
	ProcessId string         `json:"processId"`
	RpcName   string         `json:"rpcName"`
	Input     *EncodedObject `json:"input,omitempty"`
	// the timeout for the single attempt of the Process RPC API
	TimeoutSeconds         *int32                  `json:"timeoutSeconds,omitempty"`
	AppDatabaseReadRequest *AppDatabaseReadRequest `json:"appDatabaseReadRequest,omitempty"`
}

type _ProcessExecutionRpcRequest ProcessExecutionRpcRequest

// NewProcessExecutionRpcRequest instantiates a new ProcessExecutionRpcRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessExecutionRpcRequest(namespace string, processId string, rpcName string) *ProcessExecutionRpcRequest {
	this := ProcessExecutionRpcRequest{}
	this.Namespace = namespace
	this.ProcessId = processId
	this.RpcName = rpcName
	return &this
}

// NewProcessExecutionRpcRequestWithDefaults instantiates a new ProcessExecutionRpcRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessExecutionRpcRequestWithDefaults() *ProcessExecutionRpcRequest {
	this := ProcessExecutionRpcRequest{}
	return &this
}

// GetNamespace returns the Namespace field value
func (o *ProcessExecutionRpcRequest) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *ProcessExecutionRpcRequest) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *ProcessExecutionRpcRequest) SetNamespace(v string) {
	o.Namespace = v
}

// GetProcessId returns the ProcessId field value
func (o *ProcessExecutionRpcRequest) GetProcessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value
// and a boolean to check if the value has been set.
func (o *ProcessExecutionRpcRequest) GetProcessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessId, true
}

// SetProcessId sets field value
func (o *ProcessExecutionRpcRequest) SetProcessId(v string) {
	o.ProcessId = v
}

// GetRpcName returns the RpcName field value
func (o *ProcessExecutionRpcRequest) GetRpcName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RpcName
}

// GetRpcNameOk returns a tuple with the RpcName field value
// and a boolean to check if the value has been set.
func (o *ProcessExecutionRpcRequest) GetRpcNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RpcName, true
}

// SetRpcName sets field value
func (o *ProcessExecutionRpcRequest) SetRpcName(v string) {
	o.RpcName = v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *ProcessExecutionRpcRequest) GetInput() EncodedObject {
	if o == nil || IsNil(o.Input) {
		var ret EncodedObject
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessExecutionRpcRequest) GetInputOk() (*EncodedObject, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *ProcessExecutionRpcRequest) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given EncodedObject and assigns it to the Input field.
func (o *ProcessExecutionRpcRequest) SetInput(v EncodedObject) {
	o.Input = &v
}

// GetTimeoutSeconds returns the TimeoutSeconds field value if set, zero value otherwise.
func (o *ProcessExecutionRpcRequest) GetTimeoutSeconds() int32 {
	if o == nil || IsNil(o.TimeoutSeconds) {
		var ret int32
		return ret
	}
	return *o.TimeoutSeconds
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessExecutionRpcRequest) GetTimeoutSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeoutSeconds) {
		return nil, false
	}
	return o.TimeoutSeconds, true
}

// HasTimeoutSeconds returns a boolean if a field has been set.
func (o *ProcessExecutionRpcRequest) HasTimeoutSeconds() bool {
	if o != nil && !IsNil(o.TimeoutSeconds) {
		return true
	}

	return false
}

// SetTimeoutSeconds gets a reference to the given int32 and assigns it to the TimeoutSeconds field.
func (o *ProcessExecutionRpcRequest) SetTimeoutSeconds(v int32) {
	o.TimeoutSeconds = &v
}

// GetAppDatabaseReadRequest returns the AppDatabaseReadRequest field value if set, zero value otherwise.
func (o *ProcessExecutionRpcRequest) GetAppDatabaseReadRequest() AppDatabaseReadRequest {
	if o == nil || IsNil(o.AppDatabaseReadRequest) {
		var ret AppDatabaseReadRequest
		return ret
	}
	return *o.AppDatabaseReadRequest
}

// GetAppDatabaseReadRequestOk returns a tuple with the AppDatabaseReadRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessExecutionRpcRequest) GetAppDatabaseReadRequestOk() (*AppDatabaseReadRequest, bool) {
	if o == nil || IsNil(o.AppDatabaseReadRequest) {
		return nil, false
	}
	return o.AppDatabaseReadRequest, true
}

// HasAppDatabaseReadRequest returns a boolean if a field has been set.
func (o *ProcessExecutionRpcRequest) HasAppDatabaseReadRequest() bool {
	if o != nil && !IsNil(o.AppDatabaseReadRequest) {
		return true
	}

	return false
}

// SetAppDatabaseReadRequest gets a reference to the given AppDatabaseReadRequest and assigns it to the AppDatabaseReadRequest field.
func (o *ProcessExecutionRpcRequest) SetAppDatabaseReadRequest(v AppDatabaseReadRequest) {
	o.AppDatabaseReadRequest = &v
}

func (o ProcessExecutionRpcRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessExecutionRpcRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["namespace"] = o.Namespace
	toSerialize["processId"] = o.ProcessId
	toSerialize["rpcName"] = o.RpcName
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	if !IsNil(o.TimeoutSeconds) {
		toSerialize["timeoutSeconds"] = o.TimeoutSeconds
	}
	if !IsNil(o.AppDatabaseReadRequest) {
		toSerialize["appDatabaseReadRequest"] = o.AppDatabaseReadRequest
	}
	return toSerialize, nil
}

func (o *ProcessExecutionRpcRequest) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"namespace",
		"processId",
		"rpcName",
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

	varProcessExecutionRpcRequest := _ProcessExecutionRpcRequest{}

	err = json.Unmarshal(bytes, &varProcessExecutionRpcRequest)

	if err != nil {
		return err
	}

	*o = ProcessExecutionRpcRequest(varProcessExecutionRpcRequest)

	return err
}

type NullableProcessExecutionRpcRequest struct {
	value *ProcessExecutionRpcRequest
	isSet bool
}

func (v NullableProcessExecutionRpcRequest) Get() *ProcessExecutionRpcRequest {
	return v.value
}

func (v *NullableProcessExecutionRpcRequest) Set(val *ProcessExecutionRpcRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessExecutionRpcRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessExecutionRpcRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessExecutionRpcRequest(val *ProcessExecutionRpcRequest) *NullableProcessExecutionRpcRequest {
	return &NullableProcessExecutionRpcRequest{value: val, isSet: true}
}

func (v NullableProcessExecutionRpcRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessExecutionRpcRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
