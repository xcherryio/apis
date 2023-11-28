# AppDatabaseConfig

the configuration of what tables and rows to read/load for state/RPCs, including an optional initial write

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**tables** | [**List[AppDatabaseTableConfig]**](AppDatabaseTableConfig.md) |  | [optional] 

## Example

```python
from xcherryapi.models.app_database_config import AppDatabaseConfig

# TODO update the JSON string below
json = "{}"
# create an instance of AppDatabaseConfig from a JSON string
app_database_config_instance = AppDatabaseConfig.from_json(json)
# print the JSON string representation of the object
print AppDatabaseConfig.to_json()

# convert the object into a dict
app_database_config_dict = app_database_config_instance.to_dict()
# create an instance of AppDatabaseConfig from a dict
app_database_config_form_dict = app_database_config.from_dict(app_database_config_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


