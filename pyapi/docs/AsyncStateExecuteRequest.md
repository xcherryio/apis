# AsyncStateExecuteRequest

the input of the execute API

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**context** | [**Context**](Context.md) |  | 
**process_type** | **str** |  | 
**state_id** | **str** |  | 
**state_input** | [**EncodedObject**](EncodedObject.md) |  | [optional] 
**command_results** | [**CommandResults**](CommandResults.md) |  | [optional] 
**app_database_read_response** | [**AppDatabaseReadResponse**](AppDatabaseReadResponse.md) |  | [optional] 
**app_database_error** | [**AppDatabaseError**](AppDatabaseError.md) |  | [optional] 
**loaded_local_attributes** | [**LoadLocalAttributesResponse**](LoadLocalAttributesResponse.md) |  | [optional] 

## Example

```python
from xcherryapi.models.async_state_execute_request import AsyncStateExecuteRequest

# TODO update the JSON string below
json = "{}"
# create an instance of AsyncStateExecuteRequest from a JSON string
async_state_execute_request_instance = AsyncStateExecuteRequest.from_json(json)
# print the JSON string representation of the object
print AsyncStateExecuteRequest.to_json()

# convert the object into a dict
async_state_execute_request_dict = async_state_execute_request_instance.to_dict()
# create an instance of AsyncStateExecuteRequest from a dict
async_state_execute_request_form_dict = async_state_execute_request.from_dict(async_state_execute_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


