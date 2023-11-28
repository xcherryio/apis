# ProcessExecutionDescribeResponse


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**process_execution_id** | **str** |  | [optional] 
**process_type** | **str** | the process type for SDK to lookup the process definition class | [optional] 
**worker_url** | **str** | the URL for xcherry async service to make callback to worker | [optional] 
**start_timestamp** | **int** | start time of the process execution | [optional] 
**status** | [**ProcessStatus**](ProcessStatus.md) |  | [optional] 

## Example

```python
from xcherryapi.models.process_execution_describe_response import ProcessExecutionDescribeResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ProcessExecutionDescribeResponse from a JSON string
process_execution_describe_response_instance = ProcessExecutionDescribeResponse.from_json(json)
# print the JSON string representation of the object
print ProcessExecutionDescribeResponse.to_json()

# convert the object into a dict
process_execution_describe_response_dict = process_execution_describe_response_instance.to_dict()
# create an instance of ProcessExecutionDescribeResponse from a dict
process_execution_describe_response_form_dict = process_execution_describe_response.from_dict(process_execution_describe_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


