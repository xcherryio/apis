# AppDatabaseTableReadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableName** | Pointer to **string** |  | [optional] 
**Rows** | Pointer to [**[]AppDatabaseRowReadResponse**](AppDatabaseRowReadResponse.md) |  | [optional] 

## Methods

### NewAppDatabaseTableReadResponse

`func NewAppDatabaseTableReadResponse() *AppDatabaseTableReadResponse`

NewAppDatabaseTableReadResponse instantiates a new AppDatabaseTableReadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDatabaseTableReadResponseWithDefaults

`func NewAppDatabaseTableReadResponseWithDefaults() *AppDatabaseTableReadResponse`

NewAppDatabaseTableReadResponseWithDefaults instantiates a new AppDatabaseTableReadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableName

`func (o *AppDatabaseTableReadResponse) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *AppDatabaseTableReadResponse) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *AppDatabaseTableReadResponse) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *AppDatabaseTableReadResponse) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetRows

`func (o *AppDatabaseTableReadResponse) GetRows() []AppDatabaseRowReadResponse`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *AppDatabaseTableReadResponse) GetRowsOk() (*[]AppDatabaseRowReadResponse, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *AppDatabaseTableReadResponse) SetRows(v []AppDatabaseRowReadResponse)`

SetRows sets Rows field to given value.

### HasRows

`func (o *AppDatabaseTableReadResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


