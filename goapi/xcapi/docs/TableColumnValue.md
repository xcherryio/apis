# TableColumnValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbColumn** | **string** | the column name that can be used in the database query, see below for example | 
**DbQueryValue** | **string** | the plain string value that can be used in the database query(e.g. for SQL SELECT ... WHERE $Column&#x3D;$dbQueryValue or UPDATE/INSERT) | 

## Methods

### NewTableColumnValue

`func NewTableColumnValue(dbColumn string, dbQueryValue string, ) *TableColumnValue`

NewTableColumnValue instantiates a new TableColumnValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableColumnValueWithDefaults

`func NewTableColumnValueWithDefaults() *TableColumnValue`

NewTableColumnValueWithDefaults instantiates a new TableColumnValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbColumn

`func (o *TableColumnValue) GetDbColumn() string`

GetDbColumn returns the DbColumn field if non-nil, zero value otherwise.

### GetDbColumnOk

`func (o *TableColumnValue) GetDbColumnOk() (*string, bool)`

GetDbColumnOk returns a tuple with the DbColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbColumn

`func (o *TableColumnValue) SetDbColumn(v string)`

SetDbColumn sets DbColumn field to given value.


### GetDbQueryValue

`func (o *TableColumnValue) GetDbQueryValue() string`

GetDbQueryValue returns the DbQueryValue field if non-nil, zero value otherwise.

### GetDbQueryValueOk

`func (o *TableColumnValue) GetDbQueryValueOk() (*string, bool)`

GetDbQueryValueOk returns a tuple with the DbQueryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbQueryValue

`func (o *TableColumnValue) SetDbQueryValue(v string)`

SetDbQueryValue sets DbQueryValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


