# CommandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WaitingType** | [**CommandWaitingType**](CommandWaitingType.md) |  | 
**TimerCommands** | Pointer to [**[]TimerCommand**](TimerCommand.md) |  | [optional] 
**LocalQueueCommands** | Pointer to [**[]LocalQueueCommand**](LocalQueueCommand.md) |  | [optional] 

## Methods

### NewCommandRequest

`func NewCommandRequest(waitingType CommandWaitingType, ) *CommandRequest`

NewCommandRequest instantiates a new CommandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandRequestWithDefaults

`func NewCommandRequestWithDefaults() *CommandRequest`

NewCommandRequestWithDefaults instantiates a new CommandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWaitingType

`func (o *CommandRequest) GetWaitingType() CommandWaitingType`

GetWaitingType returns the WaitingType field if non-nil, zero value otherwise.

### GetWaitingTypeOk

`func (o *CommandRequest) GetWaitingTypeOk() (*CommandWaitingType, bool)`

GetWaitingTypeOk returns a tuple with the WaitingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingType

`func (o *CommandRequest) SetWaitingType(v CommandWaitingType)`

SetWaitingType sets WaitingType field to given value.


### GetTimerCommands

`func (o *CommandRequest) GetTimerCommands() []TimerCommand`

GetTimerCommands returns the TimerCommands field if non-nil, zero value otherwise.

### GetTimerCommandsOk

`func (o *CommandRequest) GetTimerCommandsOk() (*[]TimerCommand, bool)`

GetTimerCommandsOk returns a tuple with the TimerCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerCommands

`func (o *CommandRequest) SetTimerCommands(v []TimerCommand)`

SetTimerCommands sets TimerCommands field to given value.

### HasTimerCommands

`func (o *CommandRequest) HasTimerCommands() bool`

HasTimerCommands returns a boolean if a field has been set.

### GetLocalQueueCommands

`func (o *CommandRequest) GetLocalQueueCommands() []LocalQueueCommand`

GetLocalQueueCommands returns the LocalQueueCommands field if non-nil, zero value otherwise.

### GetLocalQueueCommandsOk

`func (o *CommandRequest) GetLocalQueueCommandsOk() (*[]LocalQueueCommand, bool)`

GetLocalQueueCommandsOk returns a tuple with the LocalQueueCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalQueueCommands

`func (o *CommandRequest) SetLocalQueueCommands(v []LocalQueueCommand)`

SetLocalQueueCommands sets LocalQueueCommands field to given value.

### HasLocalQueueCommands

`func (o *CommandRequest) HasLocalQueueCommands() bool`

HasLocalQueueCommands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


