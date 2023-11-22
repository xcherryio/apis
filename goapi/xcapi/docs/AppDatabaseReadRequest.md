# AppDatabaseReadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableRequests** | Pointer to [**[]AppDatabaseTableReadRequest**](AppDatabaseTableReadRequest.md) |  | [optional] 

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

### GetTableRequests

`func (o *AppDatabaseReadRequest) GetTableRequests() []AppDatabaseTableReadRequest`

GetTableRequests returns the TableRequests field if non-nil, zero value otherwise.

### GetTableRequestsOk

`func (o *AppDatabaseReadRequest) GetTableRequestsOk() (*[]AppDatabaseTableReadRequest, bool)`

GetTableRequestsOk returns a tuple with the TableRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableRequests

`func (o *AppDatabaseReadRequest) SetTableRequests(v []AppDatabaseTableReadRequest)`

SetTableRequests sets TableRequests field to given value.

### HasTableRequests

`func (o *AppDatabaseReadRequest) HasTableRequests() bool`

HasTableRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


