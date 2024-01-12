# AppDatabaseTableConfig


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**table_name** | **str** |  | 
**rows** | [**List[AppDatabaseTableRowSelector]**](AppDatabaseTableRowSelector.md) |  | 

## Example

```python
from xcherryapi.models.app_database_table_config import AppDatabaseTableConfig

# TODO update the JSON string below
json = "{}"
# create an instance of AppDatabaseTableConfig from a JSON string
app_database_table_config_instance = AppDatabaseTableConfig.from_json(json)
# print the JSON string representation of the object
print AppDatabaseTableConfig.to_json()

# convert the object into a dict
app_database_table_config_dict = app_database_table_config_instance.to_dict()
# create an instance of AppDatabaseTableConfig from a dict
app_database_table_config_form_dict = app_database_table_config.from_dict(app_database_table_config_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


