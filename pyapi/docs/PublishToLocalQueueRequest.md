# PublishToLocalQueueRequest

the request for sending messages to be consumed within a single process execution

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**namespace** | **str** |  | 
**process_id** | **str** |  | 
**messages** | [**List[LocalQueueMessage]**](LocalQueueMessage.md) |  | [optional] 

## Example

```python
from xcherryapi.models.publish_to_local_queue_request import PublishToLocalQueueRequest

# TODO update the JSON string below
json = "{}"
# create an instance of PublishToLocalQueueRequest from a JSON string
publish_to_local_queue_request_instance = PublishToLocalQueueRequest.from_json(json)
# print the JSON string representation of the object
print PublishToLocalQueueRequest.to_json()

# convert the object into a dict
publish_to_local_queue_request_dict = publish_to_local_queue_request_instance.to_dict()
# create an instance of PublishToLocalQueueRequest from a dict
publish_to_local_queue_request_form_dict = publish_to_local_queue_request.from_dict(publish_to_local_queue_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


