# ListProcessExecutionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessExecutions** | Pointer to [**[]ProcessExecutionListInfo**](ProcessExecutionListInfo.md) |  | [optional] 
**NextPageToken** | Pointer to **string** |  | [optional] 

## Methods

### NewListProcessExecutionsResponse

`func NewListProcessExecutionsResponse() *ListProcessExecutionsResponse`

NewListProcessExecutionsResponse instantiates a new ListProcessExecutionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProcessExecutionsResponseWithDefaults

`func NewListProcessExecutionsResponseWithDefaults() *ListProcessExecutionsResponse`

NewListProcessExecutionsResponseWithDefaults instantiates a new ListProcessExecutionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessExecutions

`func (o *ListProcessExecutionsResponse) GetProcessExecutions() []ProcessExecutionListInfo`

GetProcessExecutions returns the ProcessExecutions field if non-nil, zero value otherwise.

### GetProcessExecutionsOk

`func (o *ListProcessExecutionsResponse) GetProcessExecutionsOk() (*[]ProcessExecutionListInfo, bool)`

GetProcessExecutionsOk returns a tuple with the ProcessExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessExecutions

`func (o *ListProcessExecutionsResponse) SetProcessExecutions(v []ProcessExecutionListInfo)`

SetProcessExecutions sets ProcessExecutions field to given value.

### HasProcessExecutions

`func (o *ListProcessExecutionsResponse) HasProcessExecutions() bool`

HasProcessExecutions returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListProcessExecutionsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListProcessExecutionsResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListProcessExecutionsResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListProcessExecutionsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


