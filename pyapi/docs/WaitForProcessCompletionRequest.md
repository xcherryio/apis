# WaitForProcessCompletionRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**shard_id** | **int** |  | 
**process_execution_id** | **str** |  | 

## Example

```python
from xcherryapi.models.wait_for_process_completion_request import WaitForProcessCompletionRequest

# TODO update the JSON string below
json = "{}"
# create an instance of WaitForProcessCompletionRequest from a JSON string
wait_for_process_completion_request_instance = WaitForProcessCompletionRequest.from_json(json)
# print the JSON string representation of the object
print WaitForProcessCompletionRequest.to_json()

# convert the object into a dict
wait_for_process_completion_request_dict = wait_for_process_completion_request_instance.to_dict()
# create an instance of WaitForProcessCompletionRequest from a dict
wait_for_process_completion_request_form_dict = wait_for_process_completion_request.from_dict(wait_for_process_completion_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


