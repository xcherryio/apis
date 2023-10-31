# CommandResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimerResults** | Pointer to [**[]TimerResult**](TimerResult.md) |  | [optional] 
**LocalQueueResults** | Pointer to [**[]LocalQueueResult**](LocalQueueResult.md) |  | [optional] 

## Methods

### NewCommandResults

`func NewCommandResults() *CommandResults`

NewCommandResults instantiates a new CommandResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandResultsWithDefaults

`func NewCommandResultsWithDefaults() *CommandResults`

NewCommandResultsWithDefaults instantiates a new CommandResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimerResults

`func (o *CommandResults) GetTimerResults() []TimerResult`

GetTimerResults returns the TimerResults field if non-nil, zero value otherwise.

### GetTimerResultsOk

`func (o *CommandResults) GetTimerResultsOk() (*[]TimerResult, bool)`

GetTimerResultsOk returns a tuple with the TimerResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerResults

`func (o *CommandResults) SetTimerResults(v []TimerResult)`

SetTimerResults sets TimerResults field to given value.

### HasTimerResults

`func (o *CommandResults) HasTimerResults() bool`

HasTimerResults returns a boolean if a field has been set.

### GetLocalQueueResults

`func (o *CommandResults) GetLocalQueueResults() []LocalQueueResult`

GetLocalQueueResults returns the LocalQueueResults field if non-nil, zero value otherwise.

### GetLocalQueueResultsOk

`func (o *CommandResults) GetLocalQueueResultsOk() (*[]LocalQueueResult, bool)`

GetLocalQueueResultsOk returns a tuple with the LocalQueueResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalQueueResults

`func (o *CommandResults) SetLocalQueueResults(v []LocalQueueResult)`

SetLocalQueueResults sets LocalQueueResults field to given value.

### HasLocalQueueResults

`func (o *CommandResults) HasLocalQueueResults() bool`

HasLocalQueueResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


