# GlobalAttributeTableRowUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableName** | **string** |  | 
**UpdateColumns** | Pointer to [**[]TableColumnValue**](TableColumnValue.md) |  | [optional] 

## Methods

### NewGlobalAttributeTableRowUpdate

`func NewGlobalAttributeTableRowUpdate(tableName string, ) *GlobalAttributeTableRowUpdate`

NewGlobalAttributeTableRowUpdate instantiates a new GlobalAttributeTableRowUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalAttributeTableRowUpdateWithDefaults

`func NewGlobalAttributeTableRowUpdateWithDefaults() *GlobalAttributeTableRowUpdate`

NewGlobalAttributeTableRowUpdateWithDefaults instantiates a new GlobalAttributeTableRowUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableName

`func (o *GlobalAttributeTableRowUpdate) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *GlobalAttributeTableRowUpdate) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *GlobalAttributeTableRowUpdate) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetUpdateColumns

`func (o *GlobalAttributeTableRowUpdate) GetUpdateColumns() []TableColumnValue`

GetUpdateColumns returns the UpdateColumns field if non-nil, zero value otherwise.

### GetUpdateColumnsOk

`func (o *GlobalAttributeTableRowUpdate) GetUpdateColumnsOk() (*[]TableColumnValue, bool)`

GetUpdateColumnsOk returns a tuple with the UpdateColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateColumns

`func (o *GlobalAttributeTableRowUpdate) SetUpdateColumns(v []TableColumnValue)`

SetUpdateColumns sets UpdateColumns field to given value.

### HasUpdateColumns

`func (o *GlobalAttributeTableRowUpdate) HasUpdateColumns() bool`

HasUpdateColumns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


