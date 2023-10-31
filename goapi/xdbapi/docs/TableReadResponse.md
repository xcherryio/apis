# TableReadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableName** | Pointer to **string** |  | [optional] 
**Columns** | Pointer to [**[]TableColumnValue**](TableColumnValue.md) |  | [optional] 

## Methods

### NewTableReadResponse

`func NewTableReadResponse() *TableReadResponse`

NewTableReadResponse instantiates a new TableReadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableReadResponseWithDefaults

`func NewTableReadResponseWithDefaults() *TableReadResponse`

NewTableReadResponseWithDefaults instantiates a new TableReadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableName

`func (o *TableReadResponse) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *TableReadResponse) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *TableReadResponse) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *TableReadResponse) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetColumns

`func (o *TableReadResponse) GetColumns() []TableColumnValue`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TableReadResponse) GetColumnsOk() (*[]TableColumnValue, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TableReadResponse) SetColumns(v []TableColumnValue)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *TableReadResponse) HasColumns() bool`

HasColumns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


