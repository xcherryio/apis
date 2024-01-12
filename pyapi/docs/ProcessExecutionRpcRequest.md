# ProcessExecutionRpcRequest

the request for executing a RPC method of a process execution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**namespace** | **str** |  | 
**process_id** | **str** |  | 
**rpc_name** | **str** |  | 
**input** | [**EncodedObject**](EncodedObject.md) |  | [optional] 
**timeout_seconds** | **int** | the timeout for the single attempt of the Process RPC API | [optional] 
**app_database_read_request** | [**AppDatabaseReadRequest**](AppDatabaseReadRequest.md) |  | [optional] 

## Example

```python
from xcherryapi.models.process_execution_rpc_request import ProcessExecutionRpcRequest

# TODO update the JSON string below
json = "{}"
# create an instance of ProcessExecutionRpcRequest from a JSON string
process_execution_rpc_request_instance = ProcessExecutionRpcRequest.from_json(json)
# print the JSON string representation of the object
print ProcessExecutionRpcRequest.to_json()

# convert the object into a dict
process_execution_rpc_request_dict = process_execution_rpc_request_instance.to_dict()
# create an instance of ProcessExecutionRpcRequest from a dict
process_execution_rpc_request_form_dict = process_execution_rpc_request.from_dict(process_execution_rpc_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


