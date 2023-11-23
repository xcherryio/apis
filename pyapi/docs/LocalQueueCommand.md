# LocalQueueCommand


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**queue_name** | **str** |  | 
**count** | **int** | the number of identical messages to await | [optional] [default to 1]

## Example

```python
from xcherryapi.models.local_queue_command import LocalQueueCommand

# TODO update the JSON string below
json = "{}"
# create an instance of LocalQueueCommand from a JSON string
local_queue_command_instance = LocalQueueCommand.from_json(json)
# print the JSON string representation of the object
print LocalQueueCommand.to_json()

# convert the object into a dict
local_queue_command_dict = local_queue_command_instance.to_dict()
# create an instance of LocalQueueCommand from a dict
local_queue_command_form_dict = local_queue_command.from_dict(local_queue_command_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


