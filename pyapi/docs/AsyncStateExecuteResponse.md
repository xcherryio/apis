# AsyncStateExecuteResponse

the output of the execute API

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**state_decision** | [**StateDecision**](StateDecision.md) |  | 
**publish_to_local_queue** | [**List[LocalQueueMessage]**](LocalQueueMessage.md) |  | [optional] 
**write_to_app_database** | [**AppDatabaseWrite**](AppDatabaseWrite.md) |  | [optional] 
**write_to_local_attributes** | [**List[KeyValue]**](KeyValue.md) |  | [optional] 

## Example

```python
from xcherryapi.models.async_state_execute_response import AsyncStateExecuteResponse

# TODO update the JSON string below
json = "{}"
# create an instance of AsyncStateExecuteResponse from a JSON string
async_state_execute_response_instance = AsyncStateExecuteResponse.from_json(json)
# print the JSON string representation of the object
print AsyncStateExecuteResponse.to_json()

# convert the object into a dict
async_state_execute_response_dict = async_state_execute_response_instance.to_dict()
# create an instance of AsyncStateExecuteResponse from a dict
async_state_execute_response_form_dict = async_state_execute_response.from_dict(async_state_execute_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


