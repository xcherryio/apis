# ProcessExecutionListInfo


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**namespace** | **str** |  | [optional] 
**process_id** | **str** |  | [optional] 
**process_execution_id** | **str** |  | [optional] 
**process_type** | **str** |  | [optional] 
**start_timestamp** | **int** |  | [optional] 
**close_timestamp** | **int** |  | [optional] 
**status** | [**ProcessStatus**](ProcessStatus.md) |  | [optional] 

## Example

```python
from xcherryapi.models.process_execution_list_info import ProcessExecutionListInfo

# TODO update the JSON string below
json = "{}"
# create an instance of ProcessExecutionListInfo from a JSON string
process_execution_list_info_instance = ProcessExecutionListInfo.from_json(json)
# print the JSON string representation of the object
print ProcessExecutionListInfo.to_json()

# convert the object into a dict
process_execution_list_info_dict = process_execution_list_info_instance.to_dict()
# create an instance of ProcessExecutionListInfo from a dict
process_execution_list_info_form_dict = process_execution_list_info.from_dict(process_execution_list_info_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


