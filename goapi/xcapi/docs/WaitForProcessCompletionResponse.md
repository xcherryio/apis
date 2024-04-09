# WaitForProcessCompletionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to [**ProcessStatus**](ProcessStatus.md) |  | [optional] 

## Methods

### NewWaitForProcessCompletionResponse

`func NewWaitForProcessCompletionResponse() *WaitForProcessCompletionResponse`

NewWaitForProcessCompletionResponse instantiates a new WaitForProcessCompletionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaitForProcessCompletionResponseWithDefaults

`func NewWaitForProcessCompletionResponseWithDefaults() *WaitForProcessCompletionResponse`

NewWaitForProcessCompletionResponseWithDefaults instantiates a new WaitForProcessCompletionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *WaitForProcessCompletionResponse) GetTimeout() bool`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *WaitForProcessCompletionResponse) GetTimeoutOk() (*bool, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *WaitForProcessCompletionResponse) SetTimeout(v bool)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *WaitForProcessCompletionResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetStatus

`func (o *WaitForProcessCompletionResponse) GetStatus() ProcessStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WaitForProcessCompletionResponse) GetStatusOk() (*ProcessStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WaitForProcessCompletionResponse) SetStatus(v ProcessStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WaitForProcessCompletionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


