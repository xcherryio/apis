# AppDatabaseTableWrite


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**table_name** | **str** |  | 
**rows** | [**List[AppDatabaseRowWrite]**](AppDatabaseRowWrite.md) |  | [optional] 

## Example

```python
from xcherryapi.models.app_database_table_write import AppDatabaseTableWrite

# TODO update the JSON string below
json = "{}"
# create an instance of AppDatabaseTableWrite from a JSON string
app_database_table_write_instance = AppDatabaseTableWrite.from_json(json)
# print the JSON string representation of the object
print AppDatabaseTableWrite.to_json()

# convert the object into a dict
app_database_table_write_dict = app_database_table_write_instance.to_dict()
# create an instance of AppDatabaseTableWrite from a dict
app_database_table_write_form_dict = app_database_table_write.from_dict(app_database_table_write_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


