# ThreadCloseDecision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloseType** | Pointer to [**ThreadCloseType**](ThreadCloseType.md) |  | [optional] 
**CloseInput** | Pointer to [**EncodedObject**](EncodedObject.md) |  | [optional] 

## Methods

### NewThreadCloseDecision

`func NewThreadCloseDecision() *ThreadCloseDecision`

NewThreadCloseDecision instantiates a new ThreadCloseDecision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadCloseDecisionWithDefaults

`func NewThreadCloseDecisionWithDefaults() *ThreadCloseDecision`

NewThreadCloseDecisionWithDefaults instantiates a new ThreadCloseDecision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloseType

`func (o *ThreadCloseDecision) GetCloseType() ThreadCloseType`

GetCloseType returns the CloseType field if non-nil, zero value otherwise.

### GetCloseTypeOk

`func (o *ThreadCloseDecision) GetCloseTypeOk() (*ThreadCloseType, bool)`

GetCloseTypeOk returns a tuple with the CloseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseType

`func (o *ThreadCloseDecision) SetCloseType(v ThreadCloseType)`

SetCloseType sets CloseType field to given value.

### HasCloseType

`func (o *ThreadCloseDecision) HasCloseType() bool`

HasCloseType returns a boolean if a field has been set.

### GetCloseInput

`func (o *ThreadCloseDecision) GetCloseInput() EncodedObject`

GetCloseInput returns the CloseInput field if non-nil, zero value otherwise.

### GetCloseInputOk

`func (o *ThreadCloseDecision) GetCloseInputOk() (*EncodedObject, bool)`

GetCloseInputOk returns a tuple with the CloseInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseInput

`func (o *ThreadCloseDecision) SetCloseInput(v EncodedObject)`

SetCloseInput sets CloseInput field to given value.

### HasCloseInput

`func (o *ThreadCloseDecision) HasCloseInput() bool`

HasCloseInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


