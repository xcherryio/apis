# TimeRangeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EarliestTime** | Pointer to **int64** |  | [optional] 
**LatestTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewTimeRangeFilter

`func NewTimeRangeFilter() *TimeRangeFilter`

NewTimeRangeFilter instantiates a new TimeRangeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeRangeFilterWithDefaults

`func NewTimeRangeFilterWithDefaults() *TimeRangeFilter`

NewTimeRangeFilterWithDefaults instantiates a new TimeRangeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEarliestTime

`func (o *TimeRangeFilter) GetEarliestTime() int64`

GetEarliestTime returns the EarliestTime field if non-nil, zero value otherwise.

### GetEarliestTimeOk

`func (o *TimeRangeFilter) GetEarliestTimeOk() (*int64, bool)`

GetEarliestTimeOk returns a tuple with the EarliestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestTime

`func (o *TimeRangeFilter) SetEarliestTime(v int64)`

SetEarliestTime sets EarliestTime field to given value.

### HasEarliestTime

`func (o *TimeRangeFilter) HasEarliestTime() bool`

HasEarliestTime returns a boolean if a field has been set.

### GetLatestTime

`func (o *TimeRangeFilter) GetLatestTime() int64`

GetLatestTime returns the LatestTime field if non-nil, zero value otherwise.

### GetLatestTimeOk

`func (o *TimeRangeFilter) GetLatestTimeOk() (*int64, bool)`

GetLatestTimeOk returns a tuple with the LatestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTime

`func (o *TimeRangeFilter) SetLatestTime(v int64)`

SetLatestTime sets LatestTime field to given value.

### HasLatestTime

`func (o *TimeRangeFilter) HasLatestTime() bool`

HasLatestTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


