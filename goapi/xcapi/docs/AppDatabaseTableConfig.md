# AppDatabaseTableConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableName** | **string** |  | 
**Rows** | [**[]AppDatabaseTableRowSelector**](AppDatabaseTableRowSelector.md) |  | 

## Methods

### NewAppDatabaseTableConfig

`func NewAppDatabaseTableConfig(tableName string, rows []AppDatabaseTableRowSelector, ) *AppDatabaseTableConfig`

NewAppDatabaseTableConfig instantiates a new AppDatabaseTableConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDatabaseTableConfigWithDefaults

`func NewAppDatabaseTableConfigWithDefaults() *AppDatabaseTableConfig`

NewAppDatabaseTableConfigWithDefaults instantiates a new AppDatabaseTableConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableName

`func (o *AppDatabaseTableConfig) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *AppDatabaseTableConfig) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *AppDatabaseTableConfig) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetRows

`func (o *AppDatabaseTableConfig) GetRows() []AppDatabaseTableRowSelector`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *AppDatabaseTableConfig) GetRowsOk() (*[]AppDatabaseTableRowSelector, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *AppDatabaseTableConfig) SetRows(v []AppDatabaseTableRowSelector)`

SetRows sets Rows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


