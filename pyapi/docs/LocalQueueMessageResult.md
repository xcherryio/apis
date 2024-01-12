# LocalQueueMessageResult


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**dedup_id** | **str** | UUID to uniquely distinguish different messages. | 
**payload** | [**EncodedObject**](EncodedObject.md) |  | [optional] 

## Example

```python
from xcherryapi.models.local_queue_message_result import LocalQueueMessageResult

# TODO update the JSON string below
json = "{}"
# create an instance of LocalQueueMessageResult from a JSON string
local_queue_message_result_instance = LocalQueueMessageResult.from_json(json)
# print the JSON string representation of the object
print LocalQueueMessageResult.to_json()

# convert the object into a dict
local_queue_message_result_dict = local_queue_message_result_instance.to_dict()
# create an instance of LocalQueueMessageResult from a dict
local_queue_message_result_form_dict = local_queue_message_result.from_dict(local_queue_message_result_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


