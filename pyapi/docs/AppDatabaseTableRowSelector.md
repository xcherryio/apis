# AppDatabaseTableRowSelector


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**primary_key** | [**List[AppDatabaseColumnValue]**](AppDatabaseColumnValue.md) |  | 
**initial_write** | [**List[AppDatabaseColumnValue]**](AppDatabaseColumnValue.md) |  | [optional] 
**conflict_mode** | [**WriteConflictMode**](WriteConflictMode.md) |  | [optional] 

## Example

```python
from xcherryapi.models.app_database_table_row_selector import AppDatabaseTableRowSelector

# TODO update the JSON string below
json = "{}"
# create an instance of AppDatabaseTableRowSelector from a JSON string
app_database_table_row_selector_instance = AppDatabaseTableRowSelector.from_json(json)
# print the JSON string representation of the object
print AppDatabaseTableRowSelector.to_json()

# convert the object into a dict
app_database_table_row_selector_dict = app_database_table_row_selector_instance.to_dict()
# create an instance of AppDatabaseTableRowSelector from a dict
app_database_table_row_selector_form_dict = app_database_table_row_selector.from_dict(app_database_table_row_selector_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


