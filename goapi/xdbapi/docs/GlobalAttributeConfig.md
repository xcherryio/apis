# GlobalAttributeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDbTable** | Pointer to **string** |  | [optional] 
**PrimaryGlobalAttribute** | Pointer to [**GlobalAttributeValue**](GlobalAttributeValue.md) |  | [optional] 
**InitialGlobalAttributes** | Pointer to [**[]GlobalAttributeValue**](GlobalAttributeValue.md) |  | [optional] 
**InitialGlobalAttributeWriteMode** | Pointer to [**AttributeWriteConflictMode**](AttributeWriteConflictMode.md) |  | [optional] 

## Methods

### NewGlobalAttributeConfig

`func NewGlobalAttributeConfig() *GlobalAttributeConfig`

NewGlobalAttributeConfig instantiates a new GlobalAttributeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalAttributeConfigWithDefaults

`func NewGlobalAttributeConfigWithDefaults() *GlobalAttributeConfig`

NewGlobalAttributeConfigWithDefaults instantiates a new GlobalAttributeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDbTable

`func (o *GlobalAttributeConfig) GetDefaultDbTable() string`

GetDefaultDbTable returns the DefaultDbTable field if non-nil, zero value otherwise.

### GetDefaultDbTableOk

`func (o *GlobalAttributeConfig) GetDefaultDbTableOk() (*string, bool)`

GetDefaultDbTableOk returns a tuple with the DefaultDbTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDbTable

`func (o *GlobalAttributeConfig) SetDefaultDbTable(v string)`

SetDefaultDbTable sets DefaultDbTable field to given value.

### HasDefaultDbTable

`func (o *GlobalAttributeConfig) HasDefaultDbTable() bool`

HasDefaultDbTable returns a boolean if a field has been set.

### GetPrimaryGlobalAttribute

`func (o *GlobalAttributeConfig) GetPrimaryGlobalAttribute() GlobalAttributeValue`

GetPrimaryGlobalAttribute returns the PrimaryGlobalAttribute field if non-nil, zero value otherwise.

### GetPrimaryGlobalAttributeOk

`func (o *GlobalAttributeConfig) GetPrimaryGlobalAttributeOk() (*GlobalAttributeValue, bool)`

GetPrimaryGlobalAttributeOk returns a tuple with the PrimaryGlobalAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGlobalAttribute

`func (o *GlobalAttributeConfig) SetPrimaryGlobalAttribute(v GlobalAttributeValue)`

SetPrimaryGlobalAttribute sets PrimaryGlobalAttribute field to given value.

### HasPrimaryGlobalAttribute

`func (o *GlobalAttributeConfig) HasPrimaryGlobalAttribute() bool`

HasPrimaryGlobalAttribute returns a boolean if a field has been set.

### GetInitialGlobalAttributes

`func (o *GlobalAttributeConfig) GetInitialGlobalAttributes() []GlobalAttributeValue`

GetInitialGlobalAttributes returns the InitialGlobalAttributes field if non-nil, zero value otherwise.

### GetInitialGlobalAttributesOk

`func (o *GlobalAttributeConfig) GetInitialGlobalAttributesOk() (*[]GlobalAttributeValue, bool)`

GetInitialGlobalAttributesOk returns a tuple with the InitialGlobalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialGlobalAttributes

`func (o *GlobalAttributeConfig) SetInitialGlobalAttributes(v []GlobalAttributeValue)`

SetInitialGlobalAttributes sets InitialGlobalAttributes field to given value.

### HasInitialGlobalAttributes

`func (o *GlobalAttributeConfig) HasInitialGlobalAttributes() bool`

HasInitialGlobalAttributes returns a boolean if a field has been set.

### GetInitialGlobalAttributeWriteMode

`func (o *GlobalAttributeConfig) GetInitialGlobalAttributeWriteMode() AttributeWriteConflictMode`

GetInitialGlobalAttributeWriteMode returns the InitialGlobalAttributeWriteMode field if non-nil, zero value otherwise.

### GetInitialGlobalAttributeWriteModeOk

`func (o *GlobalAttributeConfig) GetInitialGlobalAttributeWriteModeOk() (*AttributeWriteConflictMode, bool)`

GetInitialGlobalAttributeWriteModeOk returns a tuple with the InitialGlobalAttributeWriteMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialGlobalAttributeWriteMode

`func (o *GlobalAttributeConfig) SetInitialGlobalAttributeWriteMode(v AttributeWriteConflictMode)`

SetInitialGlobalAttributeWriteMode sets InitialGlobalAttributeWriteMode field to given value.

### HasInitialGlobalAttributeWriteMode

`func (o *GlobalAttributeConfig) HasInitialGlobalAttributeWriteMode() bool`

HasInitialGlobalAttributeWriteMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


