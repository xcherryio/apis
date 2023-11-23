# AsyncStateWaitUntilResponse

the output of the waitUntil API

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**command_request** | [**CommandRequest**](CommandRequest.md) |  | 
**publish_to_local_queue** | [**List[LocalQueueMessage]**](LocalQueueMessage.md) |  | [optional] 

## Example

```python
from xcherryapi.models.async_state_wait_until_response import AsyncStateWaitUntilResponse

# TODO update the JSON string below
json = "{}"
# create an instance of AsyncStateWaitUntilResponse from a JSON string
async_state_wait_until_response_instance = AsyncStateWaitUntilResponse.from_json(json)
# print the JSON string representation of the object
print AsyncStateWaitUntilResponse.to_json()

# convert the object into a dict
async_state_wait_until_response_dict = async_state_wait_until_response_instance.to_dict()
# create an instance of AsyncStateWaitUntilResponse from a dict
async_state_wait_until_response_form_dict = async_state_wait_until_response.from_dict(async_state_wait_until_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


