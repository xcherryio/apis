# CommandResults


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**timer_results** | [**List[TimerResult]**](TimerResult.md) |  | [optional] 
**local_queue_results** | [**List[LocalQueueResult]**](LocalQueueResult.md) |  | [optional] 

## Example

```python
from xcherryapi.models.command_results import CommandResults

# TODO update the JSON string below
json = "{}"
# create an instance of CommandResults from a JSON string
command_results_instance = CommandResults.from_json(json)
# print the JSON string representation of the object
print CommandResults.to_json()

# convert the object into a dict
command_results_dict = command_results_instance.to_dict()
# create an instance of CommandResults from a dict
command_results_form_dict = command_results.from_dict(command_results_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


