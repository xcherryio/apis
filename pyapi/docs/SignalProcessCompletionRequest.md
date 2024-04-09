# SignalProcessCompletionRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**shard_id** | **int** |  | 
**process_execution_id** | **str** |  | 
**status** | **str** |  | 

## Example

```python
from xcherryapi.models.signal_process_completion_request import SignalProcessCompletionRequest

# TODO update the JSON string below
json = "{}"
# create an instance of SignalProcessCompletionRequest from a JSON string
signal_process_completion_request_instance = SignalProcessCompletionRequest.from_json(json)
# print the JSON string representation of the object
print SignalProcessCompletionRequest.to_json()

# convert the object into a dict
signal_process_completion_request_dict = signal_process_completion_request_instance.to_dict()
# create an instance of SignalProcessCompletionRequest from a dict
signal_process_completion_request_form_dict = signal_process_completion_request.from_dict(signal_process_completion_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


