# AppDatabaseConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tables** | Pointer to [**[]AppDatabaseTableConfig**](AppDatabaseTableConfig.md) |  | [optional] 

## Methods

### NewAppDatabaseConfig

`func NewAppDatabaseConfig() *AppDatabaseConfig`

NewAppDatabaseConfig instantiates a new AppDatabaseConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDatabaseConfigWithDefaults

`func NewAppDatabaseConfigWithDefaults() *AppDatabaseConfig`

NewAppDatabaseConfigWithDefaults instantiates a new AppDatabaseConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTables

`func (o *AppDatabaseConfig) GetTables() []AppDatabaseTableConfig`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *AppDatabaseConfig) GetTablesOk() (*[]AppDatabaseTableConfig, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *AppDatabaseConfig) SetTables(v []AppDatabaseTableConfig)`

SetTables sets Tables field to given value.

### HasTables

`func (o *AppDatabaseConfig) HasTables() bool`

HasTables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


