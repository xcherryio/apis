# AppDatabaseTableWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableName** | **string** |  | 
**Rows** | Pointer to [**[]AppDatabaseRowWrite**](AppDatabaseRowWrite.md) |  | [optional] 

## Methods

### NewAppDatabaseTableWrite

`func NewAppDatabaseTableWrite(tableName string, ) *AppDatabaseTableWrite`

NewAppDatabaseTableWrite instantiates a new AppDatabaseTableWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDatabaseTableWriteWithDefaults

`func NewAppDatabaseTableWriteWithDefaults() *AppDatabaseTableWrite`

NewAppDatabaseTableWriteWithDefaults instantiates a new AppDatabaseTableWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableName

`func (o *AppDatabaseTableWrite) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *AppDatabaseTableWrite) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *AppDatabaseTableWrite) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetRows

`func (o *AppDatabaseTableWrite) GetRows() []AppDatabaseRowWrite`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *AppDatabaseTableWrite) GetRowsOk() (*[]AppDatabaseRowWrite, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *AppDatabaseTableWrite) SetRows(v []AppDatabaseRowWrite)`

SetRows sets Rows field to given value.

### HasRows

`func (o *AppDatabaseTableWrite) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


