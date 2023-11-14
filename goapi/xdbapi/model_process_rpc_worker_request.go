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

// checks if the ProcessRpcWorkerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessRpcWorkerRequest{}

// ProcessRpcWorkerRequest the request of the worker RPC API
type ProcessRpcWorkerRequest struct {
	Context                Context                      `json:"context"`
	ProcessType            string                       `json:"processType"`
	RpcName                string                       `json:"rpcName"`
	Input                  *EncodedObject               `json:"input,omitempty"`
	LoadedGlobalAttributes *LoadGlobalAttributeResponse `json:"loadedGlobalAttributes,omitempty"`
}

// NewProcessRpcWorkerRequest instantiates a new ProcessRpcWorkerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessRpcWorkerRequest(context Context, processType string, rpcName string) *ProcessRpcWorkerRequest {
	this := ProcessRpcWorkerRequest{}
	this.Context = context
	this.ProcessType = processType
	this.RpcName = rpcName
	return &this
}

// NewProcessRpcWorkerRequestWithDefaults instantiates a new ProcessRpcWorkerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessRpcWorkerRequestWithDefaults() *ProcessRpcWorkerRequest {
	this := ProcessRpcWorkerRequest{}
	return &this
}

// GetContext returns the Context field value
func (o *ProcessRpcWorkerRequest) GetContext() Context {
	if o == nil {
		var ret Context
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *ProcessRpcWorkerRequest) GetContextOk() (*Context, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *ProcessRpcWorkerRequest) SetContext(v Context) {
	o.Context = v
}

// GetProcessType returns the ProcessType field value
func (o *ProcessRpcWorkerRequest) GetProcessType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessType
}

// GetProcessTypeOk returns a tuple with the ProcessType field value
// and a boolean to check if the value has been set.
func (o *ProcessRpcWorkerRequest) GetProcessTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessType, true
}

// SetProcessType sets field value
func (o *ProcessRpcWorkerRequest) SetProcessType(v string) {
	o.ProcessType = v
}

// GetRpcName returns the RpcName field value
func (o *ProcessRpcWorkerRequest) GetRpcName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RpcName
}

// GetRpcNameOk returns a tuple with the RpcName field value
// and a boolean to check if the value has been set.
func (o *ProcessRpcWorkerRequest) GetRpcNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RpcName, true
}

// SetRpcName sets field value
func (o *ProcessRpcWorkerRequest) SetRpcName(v string) {
	o.RpcName = v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *ProcessRpcWorkerRequest) GetInput() EncodedObject {
	if o == nil || IsNil(o.Input) {
		var ret EncodedObject
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessRpcWorkerRequest) GetInputOk() (*EncodedObject, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *ProcessRpcWorkerRequest) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given EncodedObject and assigns it to the Input field.
func (o *ProcessRpcWorkerRequest) SetInput(v EncodedObject) {
	o.Input = &v
}

// GetLoadedGlobalAttributes returns the LoadedGlobalAttributes field value if set, zero value otherwise.
func (o *ProcessRpcWorkerRequest) GetLoadedGlobalAttributes() LoadGlobalAttributeResponse {
	if o == nil || IsNil(o.LoadedGlobalAttributes) {
		var ret LoadGlobalAttributeResponse
		return ret
	}
	return *o.LoadedGlobalAttributes
}

// GetLoadedGlobalAttributesOk returns a tuple with the LoadedGlobalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessRpcWorkerRequest) GetLoadedGlobalAttributesOk() (*LoadGlobalAttributeResponse, bool) {
	if o == nil || IsNil(o.LoadedGlobalAttributes) {
		return nil, false
	}
	return o.LoadedGlobalAttributes, true
}

// HasLoadedGlobalAttributes returns a boolean if a field has been set.
func (o *ProcessRpcWorkerRequest) HasLoadedGlobalAttributes() bool {
	if o != nil && !IsNil(o.LoadedGlobalAttributes) {
		return true
	}

	return false
}

// SetLoadedGlobalAttributes gets a reference to the given LoadGlobalAttributeResponse and assigns it to the LoadedGlobalAttributes field.
func (o *ProcessRpcWorkerRequest) SetLoadedGlobalAttributes(v LoadGlobalAttributeResponse) {
	o.LoadedGlobalAttributes = &v
}

func (o ProcessRpcWorkerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessRpcWorkerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["context"] = o.Context
	toSerialize["processType"] = o.ProcessType
	toSerialize["rpcName"] = o.RpcName
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	if !IsNil(o.LoadedGlobalAttributes) {
		toSerialize["loadedGlobalAttributes"] = o.LoadedGlobalAttributes
	}
	return toSerialize, nil
}

type NullableProcessRpcWorkerRequest struct {
	value *ProcessRpcWorkerRequest
	isSet bool
}

func (v NullableProcessRpcWorkerRequest) Get() *ProcessRpcWorkerRequest {
	return v.value
}

func (v *NullableProcessRpcWorkerRequest) Set(val *ProcessRpcWorkerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessRpcWorkerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessRpcWorkerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessRpcWorkerRequest(val *ProcessRpcWorkerRequest) *NullableProcessRpcWorkerRequest {
	return &NullableProcessRpcWorkerRequest{value: val, isSet: true}
}

func (v NullableProcessRpcWorkerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessRpcWorkerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
