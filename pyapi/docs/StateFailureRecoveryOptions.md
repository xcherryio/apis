# StateFailureRecoveryOptions


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**policy** | [**StateFailureRecoveryPolicy**](StateFailureRecoveryPolicy.md) |  | 
**state_failure_proceed_state_id** | **str** |  | [optional] 
**state_failure_proceed_state_config** | [**AsyncStateConfig**](AsyncStateConfig.md) |  | [optional] 

## Example

```python
from xcherryapi.models.state_failure_recovery_options import StateFailureRecoveryOptions

# TODO update the JSON string below
json = "{}"
# create an instance of StateFailureRecoveryOptions from a JSON string
state_failure_recovery_options_instance = StateFailureRecoveryOptions.from_json(json)
# print the JSON string representation of the object
print StateFailureRecoveryOptions.to_json()

# convert the object into a dict
state_failure_recovery_options_dict = state_failure_recovery_options_instance.to_dict()
# create an instance of StateFailureRecoveryOptions from a dict
state_failure_recovery_options_form_dict = state_failure_recovery_options.from_dict(state_failure_recovery_options_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


