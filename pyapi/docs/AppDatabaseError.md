# AppDatabaseError

the error for read/write the app database. For the write error, it's from last attempt.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**app_db_error_type** | [**ErrorSubType**](ErrorSubType.md) |  | 
**app_db_error_code** | **str** | the error code from database driver | 
**app_db_error_message** | **str** | the error message from database driver | [optional] 
**app_db_error_table_name** | **str** | the first table that encounters the error to help SDK to throw the error in a friendly way  | [optional] 

## Example

```python
from xcherryapi.models.app_database_error import AppDatabaseError

# TODO update the JSON string below
json = "{}"
# create an instance of AppDatabaseError from a JSON string
app_database_error_instance = AppDatabaseError.from_json(json)
# print the JSON string representation of the object
print AppDatabaseError.to_json()

# convert the object into a dict
app_database_error_dict = app_database_error_instance.to_dict()
# create an instance of AppDatabaseError from a dict
app_database_error_form_dict = app_database_error.from_dict(app_database_error_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


