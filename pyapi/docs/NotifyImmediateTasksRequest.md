# NotifyImmediateTasksRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**shard_id** | **int** |  | 
**namespace** | **str** | optional field for distributed database without global secondary index, to pull for specific task rather than a page | [optional] 
**process_id** | **str** | optional field for distributed database without global secondary index, to pull for specific task rather than a page | [optional] 
**process_execution_id** | **str** | optional field for distributed database without global secondary index, to pull for specific task rather than a page | [optional] 

## Example

```python
from xcherryapi.models.notify_immediate_tasks_request import NotifyImmediateTasksRequest

# TODO update the JSON string below
json = "{}"
# create an instance of NotifyImmediateTasksRequest from a JSON string
notify_immediate_tasks_request_instance = NotifyImmediateTasksRequest.from_json(json)
# print the JSON string representation of the object
print NotifyImmediateTasksRequest.to_json()

# convert the object into a dict
notify_immediate_tasks_request_dict = notify_immediate_tasks_request_instance.to_dict()
# create an instance of NotifyImmediateTasksRequest from a dict
notify_immediate_tasks_request_form_dict = notify_immediate_tasks_request.from_dict(notify_immediate_tasks_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


