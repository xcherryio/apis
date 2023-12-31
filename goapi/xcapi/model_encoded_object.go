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

// checks if the EncodedObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EncodedObject{}

// EncodedObject struct for EncodedObject
type EncodedObject struct {
	Encoding string `json:"encoding"`
	Data     string `json:"data"`
}

type _EncodedObject EncodedObject

// NewEncodedObject instantiates a new EncodedObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncodedObject(encoding string, data string) *EncodedObject {
	this := EncodedObject{}
	this.Encoding = encoding
	this.Data = data
	return &this
}

// NewEncodedObjectWithDefaults instantiates a new EncodedObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncodedObjectWithDefaults() *EncodedObject {
	this := EncodedObject{}
	return &this
}

// GetEncoding returns the Encoding field value
func (o *EncodedObject) GetEncoding() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value
// and a boolean to check if the value has been set.
func (o *EncodedObject) GetEncodingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encoding, true
}

// SetEncoding sets field value
func (o *EncodedObject) SetEncoding(v string) {
	o.Encoding = v
}

// GetData returns the Data field value
func (o *EncodedObject) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EncodedObject) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *EncodedObject) SetData(v string) {
	o.Data = v
}

func (o EncodedObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EncodedObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["encoding"] = o.Encoding
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *EncodedObject) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"encoding",
		"data",
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

	varEncodedObject := _EncodedObject{}

	err = json.Unmarshal(bytes, &varEncodedObject)

	if err != nil {
		return err
	}

	*o = EncodedObject(varEncodedObject)

	return err
}

type NullableEncodedObject struct {
	value *EncodedObject
	isSet bool
}

func (v NullableEncodedObject) Get() *EncodedObject {
	return v.value
}

func (v *NullableEncodedObject) Set(val *EncodedObject) {
	v.value = val
	v.isSet = true
}

func (v NullableEncodedObject) IsSet() bool {
	return v.isSet
}

func (v *NullableEncodedObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncodedObject(val *EncodedObject) *NullableEncodedObject {
	return &NullableEncodedObject{value: val, isSet: true}
}

func (v NullableEncodedObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncodedObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
