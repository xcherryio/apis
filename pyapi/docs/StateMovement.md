# StateMovement


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**state_id** | **str** |  | 
**state_input** | [**EncodedObject**](EncodedObject.md) |  | [optional] 
**state_config** | [**AsyncStateConfig**](AsyncStateConfig.md) |  | [optional] 

## Example

```python
from xcherryapi.models.state_movement import StateMovement

# TODO update the JSON string below
json = "{}"
# create an instance of StateMovement from a JSON string
state_movement_instance = StateMovement.from_json(json)
# print the JSON string representation of the object
print StateMovement.to_json()

# convert the object into a dict
state_movement_dict = state_movement_instance.to_dict()
# create an instance of StateMovement from a dict
state_movement_form_dict = state_movement.from_dict(state_movement_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


