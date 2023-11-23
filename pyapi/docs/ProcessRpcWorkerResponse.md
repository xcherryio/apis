# ProcessRpcWorkerResponse

the response of the worker RPC API

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**output** | [**EncodedObject**](EncodedObject.md) |  | [optional] 
**state_decision** | [**StateDecision**](StateDecision.md) |  | 
**publish_to_local_queue** | [**List[LocalQueueMessage]**](LocalQueueMessage.md) |  | [optional] 
**write_to_app_database** | [**AppDatabaseWrite**](AppDatabaseWrite.md) |  | [optional] 

## Example

```python
from xcherryapi.models.process_rpc_worker_response import ProcessRpcWorkerResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ProcessRpcWorkerResponse from a JSON string
process_rpc_worker_response_instance = ProcessRpcWorkerResponse.from_json(json)
# print the JSON string representation of the object
print ProcessRpcWorkerResponse.to_json()

# convert the object into a dict
process_rpc_worker_response_dict = process_rpc_worker_response_instance.to_dict()
# create an instance of ProcessRpcWorkerResponse from a dict
process_rpc_worker_response_form_dict = process_rpc_worker_response.from_dict(process_rpc_worker_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


