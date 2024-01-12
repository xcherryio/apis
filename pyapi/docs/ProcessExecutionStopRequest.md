# ProcessExecutionStopRequest

the request for stopping a process execution

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**namespace** | **str** |  | 
**process_id** | **str** |  | 
**stop_type** | [**ProcessExecutionStopType**](ProcessExecutionStopType.md) |  | [optional] 

## Example

```python
from xcherryapi.models.process_execution_stop_request import ProcessExecutionStopRequest

# TODO update the JSON string below
json = "{}"
# create an instance of ProcessExecutionStopRequest from a JSON string
process_execution_stop_request_instance = ProcessExecutionStopRequest.from_json(json)
# print the JSON string representation of the object
print ProcessExecutionStopRequest.to_json()

# convert the object into a dict
process_execution_stop_request_dict = process_execution_stop_request_instance.to_dict()
# create an instance of ProcessExecutionStopRequest from a dict
process_execution_stop_request_form_dict = process_execution_stop_request.from_dict(process_execution_stop_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


