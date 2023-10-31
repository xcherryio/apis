# TableColumnDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbColumn** | **string** | the column name that can be used in the database query, see below for example | 

## Methods

### NewTableColumnDef

`func NewTableColumnDef(dbColumn string, ) *TableColumnDef`

NewTableColumnDef instantiates a new TableColumnDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableColumnDefWithDefaults

`func NewTableColumnDefWithDefaults() *TableColumnDef`

NewTableColumnDefWithDefaults instantiates a new TableColumnDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbColumn

`func (o *TableColumnDef) GetDbColumn() string`

GetDbColumn returns the DbColumn field if non-nil, zero value otherwise.

### GetDbColumnOk

`func (o *TableColumnDef) GetDbColumnOk() (*string, bool)`

GetDbColumnOk returns a tuple with the DbColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbColumn

`func (o *TableColumnDef) SetDbColumn(v string)`

SetDbColumn sets DbColumn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


