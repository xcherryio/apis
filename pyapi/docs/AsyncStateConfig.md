# AsyncStateConfig


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**skip_wait_until** | **bool** |  | [optional] 
**wait_until_api_timeout_seconds** | **int** | the timeout for the single attempt of AsyncState.waitUntil API | [optional] 
**execute_api_timeout_seconds** | **int** | the timeout for the single attempt of AsyncState.execute API | [optional] 
**wait_until_api_retry_policy** | [**RetryPolicy**](RetryPolicy.md) |  | [optional] 
**execute_api_retry_policy** | [**RetryPolicy**](RetryPolicy.md) |  | [optional] 
**state_failure_recovery_options** | [**StateFailureRecoveryOptions**](StateFailureRecoveryOptions.md) |  | [optional] 
**app_database_read_request** | [**AppDatabaseReadRequest**](AppDatabaseReadRequest.md) |  | [optional] 
**load_local_attributes_request** | [**LoadLocalAttributesRequest**](LoadLocalAttributesRequest.md) |  | [optional] 

## Example

```python
from xcherryapi.models.async_state_config import AsyncStateConfig

# TODO update the JSON string below
json = "{}"
# create an instance of AsyncStateConfig from a JSON string
async_state_config_instance = AsyncStateConfig.from_json(json)
# print the JSON string representation of the object
print AsyncStateConfig.to_json()

# convert the object into a dict
async_state_config_dict = async_state_config_instance.to_dict()
# create an instance of AsyncStateConfig from a dict
async_state_config_form_dict = async_state_config.from_dict(async_state_config_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


