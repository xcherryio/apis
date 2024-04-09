# WaitForProcessCompletionResponse


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**timeout** | **bool** |  | [optional] 
**status** | [**ProcessStatus**](ProcessStatus.md) |  | [optional] 

## Example

```python
from xcherryapi.models.wait_for_process_completion_response import WaitForProcessCompletionResponse

# TODO update the JSON string below
json = "{}"
# create an instance of WaitForProcessCompletionResponse from a JSON string
wait_for_process_completion_response_instance = WaitForProcessCompletionResponse.from_json(json)
# print the JSON string representation of the object
print WaitForProcessCompletionResponse.to_json()

# convert the object into a dict
wait_for_process_completion_response_dict = wait_for_process_completion_response_instance.to_dict()
# create an instance of WaitForProcessCompletionResponse from a dict
wait_for_process_completion_response_form_dict = wait_for_process_completion_response.from_dict(wait_for_process_completion_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


