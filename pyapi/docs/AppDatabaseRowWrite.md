# AppDatabaseRowWrite


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**primary_key** | [**List[AppDatabaseColumnValue]**](AppDatabaseColumnValue.md) | the PK to locate the rows for write | 
**write_columns** | [**List[AppDatabaseColumnValue]**](AppDatabaseColumnValue.md) |  | 

## Example

```python
from xcherryapi.models.app_database_row_write import AppDatabaseRowWrite

# TODO update the JSON string below
json = "{}"
# create an instance of AppDatabaseRowWrite from a JSON string
app_database_row_write_instance = AppDatabaseRowWrite.from_json(json)
# print the JSON string representation of the object
print AppDatabaseRowWrite.to_json()

# convert the object into a dict
app_database_row_write_dict = app_database_row_write_instance.to_dict()
# create an instance of AppDatabaseRowWrite from a dict
app_database_row_write_form_dict = app_database_row_write.from_dict(app_database_row_write_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


