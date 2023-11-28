# AppDatabaseColumnValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Column** | **string** | the column name that can be used in the database query, see below for example | 
**QueryValue** | **string** | the plain string value that can be used in the database query(e.g. for SQL SELECT ... WHERE $Column&#x3D;$dbQueryValue or UPDATE/INSERT) | 

## Methods

### NewAppDatabaseColumnValue

`func NewAppDatabaseColumnValue(column string, queryValue string, ) *AppDatabaseColumnValue`

NewAppDatabaseColumnValue instantiates a new AppDatabaseColumnValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDatabaseColumnValueWithDefaults

`func NewAppDatabaseColumnValueWithDefaults() *AppDatabaseColumnValue`

NewAppDatabaseColumnValueWithDefaults instantiates a new AppDatabaseColumnValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumn

`func (o *AppDatabaseColumnValue) GetColumn() string`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *AppDatabaseColumnValue) GetColumnOk() (*string, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *AppDatabaseColumnValue) SetColumn(v string)`

SetColumn sets Column field to given value.


### GetQueryValue

`func (o *AppDatabaseColumnValue) GetQueryValue() string`

GetQueryValue returns the QueryValue field if non-nil, zero value otherwise.

### GetQueryValueOk

`func (o *AppDatabaseColumnValue) GetQueryValueOk() (*string, bool)`

GetQueryValueOk returns a tuple with the QueryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryValue

`func (o *AppDatabaseColumnValue) SetQueryValue(v string)`

SetQueryValue sets QueryValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


