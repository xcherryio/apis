# LocalQueueResult


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**status** | [**CommandStatus**](CommandStatus.md) |  | 
**queue_name** | **str** |  | 
**messages** | [**List[LocalQueueMessageResult]**](LocalQueueMessageResult.md) |  | [optional] 

## Example

```python
from xcherryapi.models.local_queue_result import LocalQueueResult

# TODO update the JSON string below
json = "{}"
# create an instance of LocalQueueResult from a JSON string
local_queue_result_instance = LocalQueueResult.from_json(json)
# print the JSON string representation of the object
print LocalQueueResult.to_json()

# convert the object into a dict
local_queue_result_dict = local_queue_result_instance.to_dict()
# create an instance of LocalQueueResult from a dict
local_queue_result_form_dict = local_queue_result.from_dict(local_queue_result_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


