# AppDatabaseTableRowSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryKey** | [**[]AppDatabaseColumnValue**](AppDatabaseColumnValue.md) |  | 
**InitialWrite** | Pointer to [**[]AppDatabaseColumnValue**](AppDatabaseColumnValue.md) |  | [optional] 
**ConflictMode** | Pointer to [**WriteConflictMode**](WriteConflictMode.md) |  | [optional] 

## Methods

### NewAppDatabaseTableRowSelector

`func NewAppDatabaseTableRowSelector(primaryKey []AppDatabaseColumnValue, ) *AppDatabaseTableRowSelector`

NewAppDatabaseTableRowSelector instantiates a new AppDatabaseTableRowSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDatabaseTableRowSelectorWithDefaults

`func NewAppDatabaseTableRowSelectorWithDefaults() *AppDatabaseTableRowSelector`

NewAppDatabaseTableRowSelectorWithDefaults instantiates a new AppDatabaseTableRowSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryKey

`func (o *AppDatabaseTableRowSelector) GetPrimaryKey() []AppDatabaseColumnValue`

GetPrimaryKey returns the PrimaryKey field if non-nil, zero value otherwise.

### GetPrimaryKeyOk

`func (o *AppDatabaseTableRowSelector) GetPrimaryKeyOk() (*[]AppDatabaseColumnValue, bool)`

GetPrimaryKeyOk returns a tuple with the PrimaryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKey

`func (o *AppDatabaseTableRowSelector) SetPrimaryKey(v []AppDatabaseColumnValue)`

SetPrimaryKey sets PrimaryKey field to given value.


### GetInitialWrite

`func (o *AppDatabaseTableRowSelector) GetInitialWrite() []AppDatabaseColumnValue`

GetInitialWrite returns the InitialWrite field if non-nil, zero value otherwise.

### GetInitialWriteOk

`func (o *AppDatabaseTableRowSelector) GetInitialWriteOk() (*[]AppDatabaseColumnValue, bool)`

GetInitialWriteOk returns a tuple with the InitialWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialWrite

`func (o *AppDatabaseTableRowSelector) SetInitialWrite(v []AppDatabaseColumnValue)`

SetInitialWrite sets InitialWrite field to given value.

### HasInitialWrite

`func (o *AppDatabaseTableRowSelector) HasInitialWrite() bool`

HasInitialWrite returns a boolean if a field has been set.

### GetConflictMode

`func (o *AppDatabaseTableRowSelector) GetConflictMode() WriteConflictMode`

GetConflictMode returns the ConflictMode field if non-nil, zero value otherwise.

### GetConflictModeOk

`func (o *AppDatabaseTableRowSelector) GetConflictModeOk() (*WriteConflictMode, bool)`

GetConflictModeOk returns a tuple with the ConflictMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictMode

`func (o *AppDatabaseTableRowSelector) SetConflictMode(v WriteConflictMode)`

SetConflictMode sets ConflictMode field to given value.

### HasConflictMode

`func (o *AppDatabaseTableRowSelector) HasConflictMode() bool`

HasConflictMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


