# AppDatabaseReadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tables** | Pointer to [**[]AppDatabaseTableReadRequest**](AppDatabaseTableReadRequest.md) |  | [optional] 

## Methods

### NewAppDatabaseReadRequest

`func NewAppDatabaseReadRequest() *AppDatabaseReadRequest`

NewAppDatabaseReadRequest instantiates a new AppDatabaseReadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDatabaseReadRequestWithDefaults

`func NewAppDatabaseReadRequestWithDefaults() *AppDatabaseReadRequest`

NewAppDatabaseReadRequestWithDefaults instantiates a new AppDatabaseReadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTables

`func (o *AppDatabaseReadRequest) GetTables() []AppDatabaseTableReadRequest`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *AppDatabaseReadRequest) GetTablesOk() (*[]AppDatabaseTableReadRequest, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *AppDatabaseReadRequest) SetTables(v []AppDatabaseTableReadRequest)`

SetTables sets Tables field to given value.

### HasTables

`func (o *AppDatabaseReadRequest) HasTables() bool`

HasTables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


