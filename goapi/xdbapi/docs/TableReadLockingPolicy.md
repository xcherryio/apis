# TableReadLockingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadLockingType** | [**AttributeReadLockingType**](AttributeReadLockingType.md) |  | 
**TableName** | **string** |  | 

## Methods

### NewTableReadLockingPolicy

`func NewTableReadLockingPolicy(readLockingType AttributeReadLockingType, tableName string, ) *TableReadLockingPolicy`

NewTableReadLockingPolicy instantiates a new TableReadLockingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableReadLockingPolicyWithDefaults

`func NewTableReadLockingPolicyWithDefaults() *TableReadLockingPolicy`

NewTableReadLockingPolicyWithDefaults instantiates a new TableReadLockingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadLockingType

`func (o *TableReadLockingPolicy) GetReadLockingType() AttributeReadLockingType`

GetReadLockingType returns the ReadLockingType field if non-nil, zero value otherwise.

### GetReadLockingTypeOk

`func (o *TableReadLockingPolicy) GetReadLockingTypeOk() (*AttributeReadLockingType, bool)`

GetReadLockingTypeOk returns a tuple with the ReadLockingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadLockingType

`func (o *TableReadLockingPolicy) SetReadLockingType(v AttributeReadLockingType)`

SetReadLockingType sets ReadLockingType field to given value.


### GetTableName

`func (o *TableReadLockingPolicy) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *TableReadLockingPolicy) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *TableReadLockingPolicy) SetTableName(v string)`

SetTableName sets TableName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


