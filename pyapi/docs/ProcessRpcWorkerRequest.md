# ProcessRpcWorkerRequest

the request of the worker RPC API

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**context** | [**Context**](Context.md) |  | 
**process_type** | **str** |  | 
**rpc_name** | **str** |  | 
**input** | [**EncodedObject**](EncodedObject.md) |  | [optional] 
**app_database_read_response** | [**AppDatabaseReadResponse**](AppDatabaseReadResponse.md) |  | [optional] 

## Example

```python
from xcherryapi.models.process_rpc_worker_request import ProcessRpcWorkerRequest

# TODO update the JSON string below
json = "{}"
# create an instance of ProcessRpcWorkerRequest from a JSON string
process_rpc_worker_request_instance = ProcessRpcWorkerRequest.from_json(json)
# print the JSON string representation of the object
print ProcessRpcWorkerRequest.to_json()

# convert the object into a dict
process_rpc_worker_request_dict = process_rpc_worker_request_instance.to_dict()
# create an instance of ProcessRpcWorkerRequest from a dict
process_rpc_worker_request_form_dict = process_rpc_worker_request.from_dict(process_rpc_worker_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


