# ProcessExecutionStartRequest

the request for starting a process execution

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**namespace** | **str** |  | 
**process_id** | **str** | the user business identifier for the process, which can be used for multiple ProcessExecution based on ProcessIdReusePolicy | 
**process_type** | **str** | the process type for SDK to lookup the process definition class | 
**worker_url** | **str** | the URL for xcherry async service to make callback to worker | 
**start_state_id** | **str** | StateId of the first AsyncState to start | [optional] 
**start_state_input** | [**EncodedObject**](EncodedObject.md) |  | [optional] 
**start_state_config** | [**AsyncStateConfig**](AsyncStateConfig.md) |  | [optional] 
**process_start_config** | [**ProcessStartConfig**](ProcessStartConfig.md) |  | [optional] 

## Example

```python
from xcherryapi.models.process_execution_start_request import ProcessExecutionStartRequest

# TODO update the JSON string below
json = "{}"
# create an instance of ProcessExecutionStartRequest from a JSON string
process_execution_start_request_instance = ProcessExecutionStartRequest.from_json(json)
# print the JSON string representation of the object
print ProcessExecutionStartRequest.to_json()

# convert the object into a dict
process_execution_start_request_dict = process_execution_start_request_instance.to_dict()
# create an instance of ProcessExecutionStartRequest from a dict
process_execution_start_request_form_dict = process_execution_start_request.from_dict(process_execution_start_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


