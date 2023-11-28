# TimerCommand


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**delay_in_seconds** | **int** |  | 

## Example

```python
from xcherryapi.models.timer_command import TimerCommand

# TODO update the JSON string below
json = "{}"
# create an instance of TimerCommand from a JSON string
timer_command_instance = TimerCommand.from_json(json)
# print the JSON string representation of the object
print TimerCommand.to_json()

# convert the object into a dict
timer_command_dict = timer_command_instance.to_dict()
# create an instance of TimerCommand from a dict
timer_command_form_dict = timer_command.from_dict(timer_command_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


