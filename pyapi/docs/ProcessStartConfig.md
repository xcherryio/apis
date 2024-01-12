# ProcessStartConfig


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**timeout_seconds** | **int** |  | [optional] 
**id_reuse_policy** | [**ProcessIdReusePolicy**](ProcessIdReusePolicy.md) |  | [optional] 
**app_database_config** | [**AppDatabaseConfig**](AppDatabaseConfig.md) |  | [optional] 
**local_attribute_config** | [**LocalAttributeConfig**](LocalAttributeConfig.md) |  | [optional] 

## Example

```python
from xcherryapi.models.process_start_config import ProcessStartConfig

# TODO update the JSON string below
json = "{}"
# create an instance of ProcessStartConfig from a JSON string
process_start_config_instance = ProcessStartConfig.from_json(json)
# print the JSON string representation of the object
print ProcessStartConfig.to_json()

# convert the object into a dict
process_start_config_dict = process_start_config_instance.to_dict()
# create an instance of ProcessStartConfig from a dict
process_start_config_form_dict = process_start_config.from_dict(process_start_config_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


