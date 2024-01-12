# AppDatabaseTableReadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableName** | Pointer to **string** |  | [optional] 
**LockType** | Pointer to [**LockingType**](LockingType.md) |  | [optional] 
**Columns** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAppDatabaseTableReadRequest

`func NewAppDatabaseTableReadRequest() *AppDatabaseTableReadRequest`

NewAppDatabaseTableReadRequest instantiates a new AppDatabaseTableReadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDatabaseTableReadRequestWithDefaults

`func NewAppDatabaseTableReadRequestWithDefaults() *AppDatabaseTableReadRequest`

NewAppDatabaseTableReadRequestWithDefaults instantiates a new AppDatabaseTableReadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableName

`func (o *AppDatabaseTableReadRequest) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *AppDatabaseTableReadRequest) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *AppDatabaseTableReadRequest) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *AppDatabaseTableReadRequest) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetLockType

`func (o *AppDatabaseTableReadRequest) GetLockType() LockingType`

GetLockType returns the LockType field if non-nil, zero value otherwise.

### GetLockTypeOk

`func (o *AppDatabaseTableReadRequest) GetLockTypeOk() (*LockingType, bool)`

GetLockTypeOk returns a tuple with the LockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockType

`func (o *AppDatabaseTableReadRequest) SetLockType(v LockingType)`

SetLockType sets LockType field to given value.

### HasLockType

`func (o *AppDatabaseTableReadRequest) HasLockType() bool`

HasLockType returns a boolean if a field has been set.

### GetColumns

`func (o *AppDatabaseTableReadRequest) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *AppDatabaseTableReadRequest) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *AppDatabaseTableReadRequest) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *AppDatabaseTableReadRequest) HasColumns() bool`

HasColumns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


