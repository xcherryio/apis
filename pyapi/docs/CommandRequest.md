# CommandRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**waiting_type** | [**CommandWaitingType**](CommandWaitingType.md) |  | 
**timer_commands** | [**List[TimerCommand]**](TimerCommand.md) |  | [optional] 
**local_queue_commands** | [**List[LocalQueueCommand]**](LocalQueueCommand.md) |  | [optional] 

## Example

```python
from xcherryapi.models.command_request import CommandRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CommandRequest from a JSON string
command_request_instance = CommandRequest.from_json(json)
# print the JSON string representation of the object
print CommandRequest.to_json()

# convert the object into a dict
command_request_dict = command_request_instance.to_dict()
# create an instance of CommandRequest from a dict
command_request_form_dict = command_request.from_dict(command_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


