# GlobalAttributeTableConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableName** | **string** |  | 
**PrimaryKey** | [**[]TableColumnValue**](TableColumnValue.md) |  | 
**InitialWrite** | Pointer to [**[]TableColumnValue**](TableColumnValue.md) |  | [optional] 
**InitialWriteMode** | Pointer to [**AttributeWriteConflictMode**](AttributeWriteConflictMode.md) |  | [optional] 

## Methods

### NewGlobalAttributeTableConfig

`func NewGlobalAttributeTableConfig(tableName string, primaryKey []TableColumnValue, ) *GlobalAttributeTableConfig`

NewGlobalAttributeTableConfig instantiates a new GlobalAttributeTableConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalAttributeTableConfigWithDefaults

`func NewGlobalAttributeTableConfigWithDefaults() *GlobalAttributeTableConfig`

NewGlobalAttributeTableConfigWithDefaults instantiates a new GlobalAttributeTableConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableName

`func (o *GlobalAttributeTableConfig) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *GlobalAttributeTableConfig) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *GlobalAttributeTableConfig) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetPrimaryKey

`func (o *GlobalAttributeTableConfig) GetPrimaryKey() []TableColumnValue`

GetPrimaryKey returns the PrimaryKey field if non-nil, zero value otherwise.

### GetPrimaryKeyOk

`func (o *GlobalAttributeTableConfig) GetPrimaryKeyOk() (*[]TableColumnValue, bool)`

GetPrimaryKeyOk returns a tuple with the PrimaryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKey

`func (o *GlobalAttributeTableConfig) SetPrimaryKey(v []TableColumnValue)`

SetPrimaryKey sets PrimaryKey field to given value.


### GetInitialWrite

`func (o *GlobalAttributeTableConfig) GetInitialWrite() []TableColumnValue`

GetInitialWrite returns the InitialWrite field if non-nil, zero value otherwise.

### GetInitialWriteOk

`func (o *GlobalAttributeTableConfig) GetInitialWriteOk() (*[]TableColumnValue, bool)`

GetInitialWriteOk returns a tuple with the InitialWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialWrite

`func (o *GlobalAttributeTableConfig) SetInitialWrite(v []TableColumnValue)`

SetInitialWrite sets InitialWrite field to given value.

### HasInitialWrite

`func (o *GlobalAttributeTableConfig) HasInitialWrite() bool`

HasInitialWrite returns a boolean if a field has been set.

### GetInitialWriteMode

`func (o *GlobalAttributeTableConfig) GetInitialWriteMode() AttributeWriteConflictMode`

GetInitialWriteMode returns the InitialWriteMode field if non-nil, zero value otherwise.

### GetInitialWriteModeOk

`func (o *GlobalAttributeTableConfig) GetInitialWriteModeOk() (*AttributeWriteConflictMode, bool)`

GetInitialWriteModeOk returns a tuple with the InitialWriteMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialWriteMode

`func (o *GlobalAttributeTableConfig) SetInitialWriteMode(v AttributeWriteConflictMode)`

SetInitialWriteMode sets InitialWriteMode field to given value.

### HasInitialWriteMode

`func (o *GlobalAttributeTableConfig) HasInitialWriteMode() bool`

HasInitialWriteMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


