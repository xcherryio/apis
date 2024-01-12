# AppDatabaseReadRequest

the request to read the selected rows of configured app database tables

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**tables** | [**List[AppDatabaseTableReadRequest]**](AppDatabaseTableReadRequest.md) |  | [optional] 

## Example

```python
from xcherryapi.models.app_database_read_request import AppDatabaseReadRequest

# TODO update the JSON string below
json = "{}"
# create an instance of AppDatabaseReadRequest from a JSON string
app_database_read_request_instance = AppDatabaseReadRequest.from_json(json)
# print the JSON string representation of the object
print AppDatabaseReadRequest.to_json()

# convert the object into a dict
app_database_read_request_dict = app_database_read_request_instance.to_dict()
# create an instance of AppDatabaseReadRequest from a dict
app_database_read_request_form_dict = app_database_read_request.from_dict(app_database_read_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


