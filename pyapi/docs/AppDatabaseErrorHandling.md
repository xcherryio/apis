# AppDatabaseErrorHandling

handling the AppDatabase error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**latest_app_db_req_for_revise** | [**AppDatabaseReadRequest**](AppDatabaseReadRequest.md) |  | [optional] 

## Example

```python
from xcherryapi.models.app_database_error_handling import AppDatabaseErrorHandling

# TODO update the JSON string below
json = "{}"
# create an instance of AppDatabaseErrorHandling from a JSON string
app_database_error_handling_instance = AppDatabaseErrorHandling.from_json(json)
# print the JSON string representation of the object
print AppDatabaseErrorHandling.to_json()

# convert the object into a dict
app_database_error_handling_dict = app_database_error_handling_instance.to_dict()
# create an instance of AppDatabaseErrorHandling from a dict
app_database_error_handling_form_dict = app_database_error_handling.from_dict(app_database_error_handling_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


