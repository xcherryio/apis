# ProcessExecutionDescribeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** |  | 
**ProcessId** | **string** |  | 

## Methods

### NewProcessExecutionDescribeRequest

`func NewProcessExecutionDescribeRequest(namespace string, processId string, ) *ProcessExecutionDescribeRequest`

NewProcessExecutionDescribeRequest instantiates a new ProcessExecutionDescribeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessExecutionDescribeRequestWithDefaults

`func NewProcessExecutionDescribeRequestWithDefaults() *ProcessExecutionDescribeRequest`

NewProcessExecutionDescribeRequestWithDefaults instantiates a new ProcessExecutionDescribeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *ProcessExecutionDescribeRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ProcessExecutionDescribeRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ProcessExecutionDescribeRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetProcessId

`func (o *ProcessExecutionDescribeRequest) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *ProcessExecutionDescribeRequest) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *ProcessExecutionDescribeRequest) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


