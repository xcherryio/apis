# TimerResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**CommandStatus**](CommandStatus.md) |  | 

## Methods

### NewTimerResult

`func NewTimerResult(status CommandStatus, ) *TimerResult`

NewTimerResult instantiates a new TimerResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimerResultWithDefaults

`func NewTimerResultWithDefaults() *TimerResult`

NewTimerResultWithDefaults instantiates a new TimerResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TimerResult) GetStatus() CommandStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TimerResult) GetStatusOk() (*CommandStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TimerResult) SetStatus(v CommandStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


