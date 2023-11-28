/*
xCherry APIs

This APIs between xCherry service and SDKs

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xcapi

import (
	"encoding/json"
)

// checks if the ApiErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiErrorResponse{}

// ApiErrorResponse struct for ApiErrorResponse
type ApiErrorResponse struct {
	ErrorSubType *ErrorSubType `json:"errorSubType,omitempty"`
	// for WORKER_EXECUTION_ERROR, it's the value from WorkerErrorResponse.errorType; for APP_DATABASE_READ/WRITE_ERROR, it's the error code from database driver
	AppErrorType *string `json:"appErrorType,omitempty"`
	// for WORKER_EXECUTION_ERROR, it's the value from WorkerErrorResponse.details; for APP_DATABASE_READ/WRITE_ERROR, it's the error message from database driver; for other apiErrorType, it's the detailed message from server.
	Details *string `json:"details,omitempty"`
}

// NewApiErrorResponse instantiates a new ApiErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiErrorResponse() *ApiErrorResponse {
	this := ApiErrorResponse{}
	return &this
}

// NewApiErrorResponseWithDefaults instantiates a new ApiErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiErrorResponseWithDefaults() *ApiErrorResponse {
	this := ApiErrorResponse{}
	return &this
}

// GetErrorSubType returns the ErrorSubType field value if set, zero value otherwise.
func (o *ApiErrorResponse) GetErrorSubType() ErrorSubType {
	if o == nil || IsNil(o.ErrorSubType) {
		var ret ErrorSubType
		return ret
	}
	return *o.ErrorSubType
}

// GetErrorSubTypeOk returns a tuple with the ErrorSubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiErrorResponse) GetErrorSubTypeOk() (*ErrorSubType, bool) {
	if o == nil || IsNil(o.ErrorSubType) {
		return nil, false
	}
	return o.ErrorSubType, true
}

// HasErrorSubType returns a boolean if a field has been set.
func (o *ApiErrorResponse) HasErrorSubType() bool {
	if o != nil && !IsNil(o.ErrorSubType) {
		return true
	}

	return false
}

// SetErrorSubType gets a reference to the given ErrorSubType and assigns it to the ErrorSubType field.
func (o *ApiErrorResponse) SetErrorSubType(v ErrorSubType) {
	o.ErrorSubType = &v
}

// GetAppErrorType returns the AppErrorType field value if set, zero value otherwise.
func (o *ApiErrorResponse) GetAppErrorType() string {
	if o == nil || IsNil(o.AppErrorType) {
		var ret string
		return ret
	}
	return *o.AppErrorType
}

// GetAppErrorTypeOk returns a tuple with the AppErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiErrorResponse) GetAppErrorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AppErrorType) {
		return nil, false
	}
	return o.AppErrorType, true
}

// HasAppErrorType returns a boolean if a field has been set.
func (o *ApiErrorResponse) HasAppErrorType() bool {
	if o != nil && !IsNil(o.AppErrorType) {
		return true
	}

	return false
}

// SetAppErrorType gets a reference to the given string and assigns it to the AppErrorType field.
func (o *ApiErrorResponse) SetAppErrorType(v string) {
	o.AppErrorType = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ApiErrorResponse) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiErrorResponse) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ApiErrorResponse) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ApiErrorResponse) SetDetails(v string) {
	o.Details = &v
}

func (o ApiErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorSubType) {
		toSerialize["errorSubType"] = o.ErrorSubType
	}
	if !IsNil(o.AppErrorType) {
		toSerialize["appErrorType"] = o.AppErrorType
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableApiErrorResponse struct {
	value *ApiErrorResponse
	isSet bool
}

func (v NullableApiErrorResponse) Get() *ApiErrorResponse {
	return v.value
}

func (v *NullableApiErrorResponse) Set(val *ApiErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiErrorResponse(val *ApiErrorResponse) *NullableApiErrorResponse {
	return &NullableApiErrorResponse{value: val, isSet: true}
}

func (v NullableApiErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
