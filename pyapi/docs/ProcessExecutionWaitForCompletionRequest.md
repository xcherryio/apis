# ProcessExecutionWaitForCompletionRequest

the request for waiting for a process completion

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**namespace** | **str** |  | 
**process_id** | **str** |  | 
**timeout_seconds** | **int** | the timeout for the waiting operation | [optional] 

## Example

```python
from xcherryapi.models.process_execution_wait_for_completion_request import ProcessExecutionWaitForCompletionRequest

# TODO update the JSON string below
json = "{}"
# create an instance of ProcessExecutionWaitForCompletionRequest from a JSON string
process_execution_wait_for_completion_request_instance = ProcessExecutionWaitForCompletionRequest.from_json(json)
# print the JSON string representation of the object
print ProcessExecutionWaitForCompletionRequest.to_json()

# convert the object into a dict
process_execution_wait_for_completion_request_dict = process_execution_wait_for_completion_request_instance.to_dict()
# create an instance of ProcessExecutionWaitForCompletionRequest from a dict
process_execution_wait_for_completion_request_form_dict = process_execution_wait_for_completion_request.from_dict(process_execution_wait_for_completion_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


