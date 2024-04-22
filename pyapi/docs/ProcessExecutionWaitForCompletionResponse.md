# ProcessExecutionWaitForCompletionResponse

the response for waiting for a process completion

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**timeout** | **bool** |  | [optional] 
**stop_by_system** | **bool** |  | [optional] 
**status** | [**ProcessStatus**](ProcessStatus.md) |  | [optional] 

## Example

```python
from xcherryapi.models.process_execution_wait_for_completion_response import ProcessExecutionWaitForCompletionResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ProcessExecutionWaitForCompletionResponse from a JSON string
process_execution_wait_for_completion_response_instance = ProcessExecutionWaitForCompletionResponse.from_json(json)
# print the JSON string representation of the object
print ProcessExecutionWaitForCompletionResponse.to_json()

# convert the object into a dict
process_execution_wait_for_completion_response_dict = process_execution_wait_for_completion_response_instance.to_dict()
# create an instance of ProcessExecutionWaitForCompletionResponse from a dict
process_execution_wait_for_completion_response_form_dict = process_execution_wait_for_completion_response.from_dict(process_execution_wait_for_completion_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


