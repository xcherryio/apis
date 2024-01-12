# NotifyTimerTasksRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**shard_id** | **int** |  | 
**fire_timestamps** | **List[int]** | the fire timestamp of all timer tasks to pull | 
**namespace** | **str** | optional field for distributed database without global secondary index, to pull for specific task rather than a page | [optional] 
**process_id** | **str** | optional field for distributed database without global secondary index, to pull for specific task rather than a page | [optional] 
**process_execution_id** | **str** | optional field for distributed database without global secondary index, to pull for specific task rather than a page | [optional] 

## Example

```python
from xcherryapi.models.notify_timer_tasks_request import NotifyTimerTasksRequest

# TODO update the JSON string below
json = "{}"
# create an instance of NotifyTimerTasksRequest from a JSON string
notify_timer_tasks_request_instance = NotifyTimerTasksRequest.from_json(json)
# print the JSON string representation of the object
print NotifyTimerTasksRequest.to_json()

# convert the object into a dict
notify_timer_tasks_request_dict = notify_timer_tasks_request_instance.to_dict()
# create an instance of NotifyTimerTasksRequest from a dict
notify_timer_tasks_request_form_dict = notify_timer_tasks_request.from_dict(notify_timer_tasks_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


