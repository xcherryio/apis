# ProcessExecutionDescribeRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**namespace** | **str** |  | 
**process_id** | **str** |  | 

## Example

```python
from xcherryapi.models.process_execution_describe_request import ProcessExecutionDescribeRequest

# TODO update the JSON string below
json = "{}"
# create an instance of ProcessExecutionDescribeRequest from a JSON string
process_execution_describe_request_instance = ProcessExecutionDescribeRequest.from_json(json)
# print the JSON string representation of the object
print ProcessExecutionDescribeRequest.to_json()

# convert the object into a dict
process_execution_describe_request_dict = process_execution_describe_request_instance.to_dict()
# create an instance of ProcessExecutionDescribeRequest from a dict
process_execution_describe_request_form_dict = process_execution_describe_request.from_dict(process_execution_describe_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


