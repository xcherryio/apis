# ListProcessExecutionsRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**namespace** | **str** |  | 
**page_size** | **int** |  | [default to 100]
**next_page_token** | **str** |  | [optional] 
**status_filter** | [**ProcessStatus**](ProcessStatus.md) |  | [optional] 
**start_time_filter** | [**TimeRangeFilter**](TimeRangeFilter.md) |  | [optional] 
**process_type_filter** | [**ProcessTypeFilter**](ProcessTypeFilter.md) |  | [optional] 
**process_id_filter** | [**ProcessIdFilter**](ProcessIdFilter.md) |  | [optional] 

## Example

```python
from xcherryapi.models.list_process_executions_request import ListProcessExecutionsRequest

# TODO update the JSON string below
json = "{}"
# create an instance of ListProcessExecutionsRequest from a JSON string
list_process_executions_request_instance = ListProcessExecutionsRequest.from_json(json)
# print the JSON string representation of the object
print ListProcessExecutionsRequest.to_json()

# convert the object into a dict
list_process_executions_request_dict = list_process_executions_request_instance.to_dict()
# create an instance of ListProcessExecutionsRequest from a dict
list_process_executions_request_form_dict = list_process_executions_request.from_dict(list_process_executions_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


