# CommandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WaitingType** | Pointer to [**CommandWaitingType**](CommandWaitingType.md) |  | [optional] 

## Methods

### NewCommandRequest

`func NewCommandRequest() *CommandRequest`

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

### HasWaitingType

`func (o *CommandRequest) HasWaitingType() bool`

HasWaitingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


