# GlobalAttributeValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbColumn** | **string** | the column name that can be used in the database query, see below for example | 
**DbQueryValue** | **string** | the plain string value that can be used in the database query(e.g. for SQL SELECT ... WHERE $Column&#x3D;$dbQueryValue or UPDATE/INSERT) | 
**AlternativeTable** | Pointer to **string** | if not empty, the global attribute is mapped to a different table name than the default table. Must be used with alternativeTableForeignKeyColumn | [optional] 
**AlternativeTableForeignKeyColumn** | Pointer to **string** | Must present when alternativeTable is not empty. It means the foreign key column to join with the primary key of the default table | [optional] 

## Methods

### NewGlobalAttributeValue

`func NewGlobalAttributeValue(dbColumn string, dbQueryValue string, ) *GlobalAttributeValue`

NewGlobalAttributeValue instantiates a new GlobalAttributeValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalAttributeValueWithDefaults

`func NewGlobalAttributeValueWithDefaults() *GlobalAttributeValue`

NewGlobalAttributeValueWithDefaults instantiates a new GlobalAttributeValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbColumn

`func (o *GlobalAttributeValue) GetDbColumn() string`

GetDbColumn returns the DbColumn field if non-nil, zero value otherwise.

### GetDbColumnOk

`func (o *GlobalAttributeValue) GetDbColumnOk() (*string, bool)`

GetDbColumnOk returns a tuple with the DbColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbColumn

`func (o *GlobalAttributeValue) SetDbColumn(v string)`

SetDbColumn sets DbColumn field to given value.


### GetDbQueryValue

`func (o *GlobalAttributeValue) GetDbQueryValue() string`

GetDbQueryValue returns the DbQueryValue field if non-nil, zero value otherwise.

### GetDbQueryValueOk

`func (o *GlobalAttributeValue) GetDbQueryValueOk() (*string, bool)`

GetDbQueryValueOk returns a tuple with the DbQueryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbQueryValue

`func (o *GlobalAttributeValue) SetDbQueryValue(v string)`

SetDbQueryValue sets DbQueryValue field to given value.


### GetAlternativeTable

`func (o *GlobalAttributeValue) GetAlternativeTable() string`

GetAlternativeTable returns the AlternativeTable field if non-nil, zero value otherwise.

### GetAlternativeTableOk

`func (o *GlobalAttributeValue) GetAlternativeTableOk() (*string, bool)`

GetAlternativeTableOk returns a tuple with the AlternativeTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeTable

`func (o *GlobalAttributeValue) SetAlternativeTable(v string)`

SetAlternativeTable sets AlternativeTable field to given value.

### HasAlternativeTable

`func (o *GlobalAttributeValue) HasAlternativeTable() bool`

HasAlternativeTable returns a boolean if a field has been set.

### GetAlternativeTableForeignKeyColumn

`func (o *GlobalAttributeValue) GetAlternativeTableForeignKeyColumn() string`

GetAlternativeTableForeignKeyColumn returns the AlternativeTableForeignKeyColumn field if non-nil, zero value otherwise.

### GetAlternativeTableForeignKeyColumnOk

`func (o *GlobalAttributeValue) GetAlternativeTableForeignKeyColumnOk() (*string, bool)`

GetAlternativeTableForeignKeyColumnOk returns a tuple with the AlternativeTableForeignKeyColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeTableForeignKeyColumn

`func (o *GlobalAttributeValue) SetAlternativeTableForeignKeyColumn(v string)`

SetAlternativeTableForeignKeyColumn sets AlternativeTableForeignKeyColumn field to given value.

### HasAlternativeTableForeignKeyColumn

`func (o *GlobalAttributeValue) HasAlternativeTableForeignKeyColumn() bool`

HasAlternativeTableForeignKeyColumn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


