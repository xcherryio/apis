# ProcessExecutionRpcResponse

the response for executing a RPC method of a process execution

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**output** | [**EncodedObject**](EncodedObject.md) |  | [optional] 

## Example

```python
from xcherryapi.models.process_execution_rpc_response import ProcessExecutionRpcResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ProcessExecutionRpcResponse from a JSON string
process_execution_rpc_response_instance = ProcessExecutionRpcResponse.from_json(json)
# print the JSON string representation of the object
print ProcessExecutionRpcResponse.to_json()

# convert the object into a dict
process_execution_rpc_response_dict = process_execution_rpc_response_instance.to_dict()
# create an instance of ProcessExecutionRpcResponse from a dict
process_execution_rpc_response_form_dict = process_execution_rpc_response.from_dict(process_execution_rpc_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


