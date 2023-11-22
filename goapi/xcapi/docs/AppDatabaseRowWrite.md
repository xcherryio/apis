# AppDatabaseRowWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryKey** | [**[]AppDatabaseColumnValue**](AppDatabaseColumnValue.md) | the PK to locate the rows for write | 
**WriteColumns** | [**[]AppDatabaseColumnValue**](AppDatabaseColumnValue.md) |  | 

## Methods

### NewAppDatabaseRowWrite

`func NewAppDatabaseRowWrite(primaryKey []AppDatabaseColumnValue, writeColumns []AppDatabaseColumnValue, ) *AppDatabaseRowWrite`

NewAppDatabaseRowWrite instantiates a new AppDatabaseRowWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDatabaseRowWriteWithDefaults

`func NewAppDatabaseRowWriteWithDefaults() *AppDatabaseRowWrite`

NewAppDatabaseRowWriteWithDefaults instantiates a new AppDatabaseRowWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryKey

`func (o *AppDatabaseRowWrite) GetPrimaryKey() []AppDatabaseColumnValue`

GetPrimaryKey returns the PrimaryKey field if non-nil, zero value otherwise.

### GetPrimaryKeyOk

`func (o *AppDatabaseRowWrite) GetPrimaryKeyOk() (*[]AppDatabaseColumnValue, bool)`

GetPrimaryKeyOk returns a tuple with the PrimaryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKey

`func (o *AppDatabaseRowWrite) SetPrimaryKey(v []AppDatabaseColumnValue)`

SetPrimaryKey sets PrimaryKey field to given value.


### GetWriteColumns

`func (o *AppDatabaseRowWrite) GetWriteColumns() []AppDatabaseColumnValue`

GetWriteColumns returns the WriteColumns field if non-nil, zero value otherwise.

### GetWriteColumnsOk

`func (o *AppDatabaseRowWrite) GetWriteColumnsOk() (*[]AppDatabaseColumnValue, bool)`

GetWriteColumnsOk returns a tuple with the WriteColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteColumns

`func (o *AppDatabaseRowWrite) SetWriteColumns(v []AppDatabaseColumnValue)`

SetWriteColumns sets WriteColumns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


