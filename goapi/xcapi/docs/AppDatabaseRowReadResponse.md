# AppDatabaseRowReadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to [**[]AppDatabaseColumnValue**](AppDatabaseColumnValue.md) |  | [optional] 

## Methods

### NewAppDatabaseRowReadResponse

`func NewAppDatabaseRowReadResponse() *AppDatabaseRowReadResponse`

NewAppDatabaseRowReadResponse instantiates a new AppDatabaseRowReadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDatabaseRowReadResponseWithDefaults

`func NewAppDatabaseRowReadResponseWithDefaults() *AppDatabaseRowReadResponse`

NewAppDatabaseRowReadResponseWithDefaults instantiates a new AppDatabaseRowReadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *AppDatabaseRowReadResponse) GetColumns() []AppDatabaseColumnValue`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *AppDatabaseRowReadResponse) GetColumnsOk() (*[]AppDatabaseColumnValue, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *AppDatabaseRowReadResponse) SetColumns(v []AppDatabaseColumnValue)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *AppDatabaseRowReadResponse) HasColumns() bool`

HasColumns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


