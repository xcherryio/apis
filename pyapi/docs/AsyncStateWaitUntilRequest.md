# AsyncStateWaitUntilRequest

the input of the waitUntil API

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**context** | [**Context**](Context.md) |  | 
**process_type** | **str** |  | 
**state_id** | **str** |  | 
**state_input** | [**EncodedObject**](EncodedObject.md) |  | [optional] 

## Example

```python
from xcherryapi.models.async_state_wait_until_request import AsyncStateWaitUntilRequest

# TODO update the JSON string below
json = "{}"
# create an instance of AsyncStateWaitUntilRequest from a JSON string
async_state_wait_until_request_instance = AsyncStateWaitUntilRequest.from_json(json)
# print the JSON string representation of the object
print AsyncStateWaitUntilRequest.to_json()

# convert the object into a dict
async_state_wait_until_request_dict = async_state_wait_until_request_instance.to_dict()
# create an instance of AsyncStateWaitUntilRequest from a dict
async_state_wait_until_request_form_dict = async_state_wait_until_request.from_dict(async_state_wait_until_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


