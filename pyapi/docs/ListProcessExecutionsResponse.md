# ListProcessExecutionsResponse


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**process_executions** | [**List[ProcessExecutionListInfo]**](ProcessExecutionListInfo.md) |  | [optional] 
**next_page_token** | **str** |  | [optional] 

## Example

```python
from xcherryapi.models.list_process_executions_response import ListProcessExecutionsResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ListProcessExecutionsResponse from a JSON string
list_process_executions_response_instance = ListProcessExecutionsResponse.from_json(json)
# print the JSON string representation of the object
print ListProcessExecutionsResponse.to_json()

# convert the object into a dict
list_process_executions_response_dict = list_process_executions_response_instance.to_dict()
# create an instance of ListProcessExecutionsResponse from a dict
list_process_executions_response_form_dict = list_process_executions_response.from_dict(list_process_executions_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


