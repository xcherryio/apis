# ProcessStartConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeoutSeconds** | Pointer to **int32** |  | [optional] 
**IdReusePolicy** | Pointer to [**ProcessIdReusePolicy**](ProcessIdReusePolicy.md) |  | [optional] 
**GlobalAttributeConfig** | Pointer to [**GlobalAttributeConfig**](GlobalAttributeConfig.md) |  | [optional] 

## Methods

### NewProcessStartConfig

`func NewProcessStartConfig() *ProcessStartConfig`

NewProcessStartConfig instantiates a new ProcessStartConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessStartConfigWithDefaults

`func NewProcessStartConfigWithDefaults() *ProcessStartConfig`

NewProcessStartConfigWithDefaults instantiates a new ProcessStartConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeoutSeconds

`func (o *ProcessStartConfig) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *ProcessStartConfig) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *ProcessStartConfig) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *ProcessStartConfig) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetIdReusePolicy

`func (o *ProcessStartConfig) GetIdReusePolicy() ProcessIdReusePolicy`

GetIdReusePolicy returns the IdReusePolicy field if non-nil, zero value otherwise.

### GetIdReusePolicyOk

`func (o *ProcessStartConfig) GetIdReusePolicyOk() (*ProcessIdReusePolicy, bool)`

GetIdReusePolicyOk returns a tuple with the IdReusePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdReusePolicy

`func (o *ProcessStartConfig) SetIdReusePolicy(v ProcessIdReusePolicy)`

SetIdReusePolicy sets IdReusePolicy field to given value.

### HasIdReusePolicy

`func (o *ProcessStartConfig) HasIdReusePolicy() bool`

HasIdReusePolicy returns a boolean if a field has been set.

### GetGlobalAttributeConfig

`func (o *ProcessStartConfig) GetGlobalAttributeConfig() GlobalAttributeConfig`

GetGlobalAttributeConfig returns the GlobalAttributeConfig field if non-nil, zero value otherwise.

### GetGlobalAttributeConfigOk

`func (o *ProcessStartConfig) GetGlobalAttributeConfigOk() (*GlobalAttributeConfig, bool)`

GetGlobalAttributeConfigOk returns a tuple with the GlobalAttributeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAttributeConfig

`func (o *ProcessStartConfig) SetGlobalAttributeConfig(v GlobalAttributeConfig)`

SetGlobalAttributeConfig sets GlobalAttributeConfig field to given value.

### HasGlobalAttributeConfig

`func (o *ProcessStartConfig) HasGlobalAttributeConfig() bool`

HasGlobalAttributeConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


