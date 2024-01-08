# ListProcessExecutionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** |  | 
**PageSize** | **int32** |  | [default to 100]
**NextPageToken** | Pointer to **string** |  | [optional] 
**StatusFilter** | Pointer to [**ProcessStatus**](ProcessStatus.md) |  | [optional] 
**StartTimeFilter** | Pointer to [**TimeRangeFilter**](TimeRangeFilter.md) |  | [optional] 
**ProcessTypeFilter** | Pointer to [**ProcessTypeFilter**](ProcessTypeFilter.md) |  | [optional] 
**ProcessIdFilter** | Pointer to [**ProcessIdFilter**](ProcessIdFilter.md) |  | [optional] 

## Methods

### NewListProcessExecutionsRequest

`func NewListProcessExecutionsRequest(namespace string, pageSize int32, ) *ListProcessExecutionsRequest`

NewListProcessExecutionsRequest instantiates a new ListProcessExecutionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProcessExecutionsRequestWithDefaults

`func NewListProcessExecutionsRequestWithDefaults() *ListProcessExecutionsRequest`

NewListProcessExecutionsRequestWithDefaults instantiates a new ListProcessExecutionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *ListProcessExecutionsRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ListProcessExecutionsRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ListProcessExecutionsRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetPageSize

`func (o *ListProcessExecutionsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListProcessExecutionsRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListProcessExecutionsRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetNextPageToken

`func (o *ListProcessExecutionsRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListProcessExecutionsRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListProcessExecutionsRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListProcessExecutionsRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetStatusFilter

`func (o *ListProcessExecutionsRequest) GetStatusFilter() ProcessStatus`

GetStatusFilter returns the StatusFilter field if non-nil, zero value otherwise.

### GetStatusFilterOk

`func (o *ListProcessExecutionsRequest) GetStatusFilterOk() (*ProcessStatus, bool)`

GetStatusFilterOk returns a tuple with the StatusFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusFilter

`func (o *ListProcessExecutionsRequest) SetStatusFilter(v ProcessStatus)`

SetStatusFilter sets StatusFilter field to given value.

### HasStatusFilter

`func (o *ListProcessExecutionsRequest) HasStatusFilter() bool`

HasStatusFilter returns a boolean if a field has been set.

### GetStartTimeFilter

`func (o *ListProcessExecutionsRequest) GetStartTimeFilter() TimeRangeFilter`

GetStartTimeFilter returns the StartTimeFilter field if non-nil, zero value otherwise.

### GetStartTimeFilterOk

`func (o *ListProcessExecutionsRequest) GetStartTimeFilterOk() (*TimeRangeFilter, bool)`

GetStartTimeFilterOk returns a tuple with the StartTimeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeFilter

`func (o *ListProcessExecutionsRequest) SetStartTimeFilter(v TimeRangeFilter)`

SetStartTimeFilter sets StartTimeFilter field to given value.

### HasStartTimeFilter

`func (o *ListProcessExecutionsRequest) HasStartTimeFilter() bool`

HasStartTimeFilter returns a boolean if a field has been set.

### GetProcessTypeFilter

`func (o *ListProcessExecutionsRequest) GetProcessTypeFilter() ProcessTypeFilter`

GetProcessTypeFilter returns the ProcessTypeFilter field if non-nil, zero value otherwise.

### GetProcessTypeFilterOk

`func (o *ListProcessExecutionsRequest) GetProcessTypeFilterOk() (*ProcessTypeFilter, bool)`

GetProcessTypeFilterOk returns a tuple with the ProcessTypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessTypeFilter

`func (o *ListProcessExecutionsRequest) SetProcessTypeFilter(v ProcessTypeFilter)`

SetProcessTypeFilter sets ProcessTypeFilter field to given value.

### HasProcessTypeFilter

`func (o *ListProcessExecutionsRequest) HasProcessTypeFilter() bool`

HasProcessTypeFilter returns a boolean if a field has been set.

### GetProcessIdFilter

`func (o *ListProcessExecutionsRequest) GetProcessIdFilter() ProcessIdFilter`

GetProcessIdFilter returns the ProcessIdFilter field if non-nil, zero value otherwise.

### GetProcessIdFilterOk

`func (o *ListProcessExecutionsRequest) GetProcessIdFilterOk() (*ProcessIdFilter, bool)`

GetProcessIdFilterOk returns a tuple with the ProcessIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessIdFilter

`func (o *ListProcessExecutionsRequest) SetProcessIdFilter(v ProcessIdFilter)`

SetProcessIdFilter sets ProcessIdFilter field to given value.

### HasProcessIdFilter

`func (o *ListProcessExecutionsRequest) HasProcessIdFilter() bool`

HasProcessIdFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


