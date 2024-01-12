# AppDatabaseWrite

the write operation for state/RPCs to write to the app database values

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**tables** | [**List[AppDatabaseTableWrite]**](AppDatabaseTableWrite.md) |  | [optional] 

## Example

```python
from xcherryapi.models.app_database_write import AppDatabaseWrite

# TODO update the JSON string below
json = "{}"
# create an instance of AppDatabaseWrite from a JSON string
app_database_write_instance = AppDatabaseWrite.from_json(json)
# print the JSON string representation of the object
print AppDatabaseWrite.to_json()

# convert the object into a dict
app_database_write_dict = app_database_write_instance.to_dict()
# create an instance of AppDatabaseWrite from a dict
app_database_write_form_dict = app_database_write.from_dict(app_database_write_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


