# Context


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**process_id** | **str** |  | 
**process_execution_id** | **str** |  | 
**process_started_timestamp** | **int** |  | 
**state_execution_id** | **str** | stateExecutionId is for async state API only | [optional] 
**first_attempt_timestamp** | **int** | for async state API only(during backoff retry) | [optional] 
**attempt** | **int** | for async state API only(during backoff retry) | [optional] 
**recover_from_state_execution_id** | **str** | for async state API only, state id + sequence number | [optional] 
**recover_from_api** | [**WorkerApiType**](WorkerApiType.md) |  | [optional] 

## Example

```python
from xcherryapi.models.context import Context

# TODO update the JSON string below
json = "{}"
# create an instance of Context from a JSON string
context_instance = Context.from_json(json)
# print the JSON string representation of the object
print Context.to_json()

# convert the object into a dict
context_dict = context_instance.to_dict()
# create an instance of Context from a dict
context_form_dict = context.from_dict(context_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


