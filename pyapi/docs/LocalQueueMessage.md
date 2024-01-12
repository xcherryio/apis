# LocalQueueMessage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**queue_name** | **str** |  | 
**dedup_id** | **str** | UUID to uniquely distinguish different messages. If not specified, the server will generate a UUID instead. | [optional] 
**payload** | [**EncodedObject**](EncodedObject.md) |  | [optional] 

## Example

```python
from xcherryapi.models.local_queue_message import LocalQueueMessage

# TODO update the JSON string below
json = "{}"
# create an instance of LocalQueueMessage from a JSON string
local_queue_message_instance = LocalQueueMessage.from_json(json)
# print the JSON string representation of the object
print LocalQueueMessage.to_json()

# convert the object into a dict
local_queue_message_dict = local_queue_message_instance.to_dict()
# create an instance of LocalQueueMessage from a dict
local_queue_message_form_dict = local_queue_message.from_dict(local_queue_message_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


