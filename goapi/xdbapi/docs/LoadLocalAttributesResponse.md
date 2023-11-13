# LoadLocalAttributesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**[]KeyValue**](KeyValue.md) |  | [optional] 

## Methods

### NewLoadLocalAttributesResponse

`func NewLoadLocalAttributesResponse() *LoadLocalAttributesResponse`

NewLoadLocalAttributesResponse instantiates a new LoadLocalAttributesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadLocalAttributesResponseWithDefaults

`func NewLoadLocalAttributesResponseWithDefaults() *LoadLocalAttributesResponse`

NewLoadLocalAttributesResponseWithDefaults instantiates a new LoadLocalAttributesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *LoadLocalAttributesResponse) GetAttributes() []KeyValue`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LoadLocalAttributesResponse) GetAttributesOk() (*[]KeyValue, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LoadLocalAttributesResponse) SetAttributes(v []KeyValue)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *LoadLocalAttributesResponse) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


