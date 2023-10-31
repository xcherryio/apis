# TableReadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableName** | Pointer to **string** |  | [optional] 
**LockingPolicy** | Pointer to [**TableReadLockingPolicy**](TableReadLockingPolicy.md) |  | [optional] 
**Columns** | Pointer to [**[]TableColumnDef**](TableColumnDef.md) |  | [optional] 

## Methods

### NewTableReadRequest

`func NewTableReadRequest() *TableReadRequest`

NewTableReadRequest instantiates a new TableReadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableReadRequestWithDefaults

`func NewTableReadRequestWithDefaults() *TableReadRequest`

NewTableReadRequestWithDefaults instantiates a new TableReadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableName

`func (o *TableReadRequest) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *TableReadRequest) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *TableReadRequest) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *TableReadRequest) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetLockingPolicy

`func (o *TableReadRequest) GetLockingPolicy() TableReadLockingPolicy`

GetLockingPolicy returns the LockingPolicy field if non-nil, zero value otherwise.

### GetLockingPolicyOk

`func (o *TableReadRequest) GetLockingPolicyOk() (*TableReadLockingPolicy, bool)`

GetLockingPolicyOk returns a tuple with the LockingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockingPolicy

`func (o *TableReadRequest) SetLockingPolicy(v TableReadLockingPolicy)`

SetLockingPolicy sets LockingPolicy field to given value.

### HasLockingPolicy

`func (o *TableReadRequest) HasLockingPolicy() bool`

HasLockingPolicy returns a boolean if a field has been set.

### GetColumns

`func (o *TableReadRequest) GetColumns() []TableColumnDef`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TableReadRequest) GetColumnsOk() (*[]TableColumnDef, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TableReadRequest) SetColumns(v []TableColumnDef)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *TableReadRequest) HasColumns() bool`

HasColumns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


