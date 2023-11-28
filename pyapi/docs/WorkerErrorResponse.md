# WorkerErrorResponse


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**detail** | **str** | an optional field to let application set some detailed information. Default to the error message + stacktrace of the error | [optional] 
**error_type** | **str** | an optional field for error handling. Default to the class/error Name | 

## Example

```python
from xcherryapi.models.worker_error_response import WorkerErrorResponse

# TODO update the JSON string below
json = "{}"
# create an instance of WorkerErrorResponse from a JSON string
worker_error_response_instance = WorkerErrorResponse.from_json(json)
# print the JSON string representation of the object
print WorkerErrorResponse.to_json()

# convert the object into a dict
worker_error_response_dict = worker_error_response_instance.to_dict()
# create an instance of WorkerErrorResponse from a dict
worker_error_response_form_dict = worker_error_response.from_dict(worker_error_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


