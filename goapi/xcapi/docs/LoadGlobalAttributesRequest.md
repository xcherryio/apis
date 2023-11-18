# LoadGlobalAttributesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableRequests** | Pointer to [**[]TableReadRequest**](TableReadRequest.md) |  | [optional] 

## Methods

### NewLoadGlobalAttributesRequest

`func NewLoadGlobalAttributesRequest() *LoadGlobalAttributesRequest`

NewLoadGlobalAttributesRequest instantiates a new LoadGlobalAttributesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadGlobalAttributesRequestWithDefaults

`func NewLoadGlobalAttributesRequestWithDefaults() *LoadGlobalAttributesRequest`

NewLoadGlobalAttributesRequestWithDefaults instantiates a new LoadGlobalAttributesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableRequests

`func (o *LoadGlobalAttributesRequest) GetTableRequests() []TableReadRequest`

GetTableRequests returns the TableRequests field if non-nil, zero value otherwise.

### GetTableRequestsOk

`func (o *LoadGlobalAttributesRequest) GetTableRequestsOk() (*[]TableReadRequest, bool)`

GetTableRequestsOk returns a tuple with the TableRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableRequests

`func (o *LoadGlobalAttributesRequest) SetTableRequests(v []TableReadRequest)`

SetTableRequests sets TableRequests field to given value.

### HasTableRequests

`func (o *LoadGlobalAttributesRequest) HasTableRequests() bool`

HasTableRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


