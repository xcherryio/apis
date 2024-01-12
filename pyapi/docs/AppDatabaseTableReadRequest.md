# AppDatabaseTableReadRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**table_name** | **str** |  | [optional] 
**lock_type** | [**LockType**](LockType.md) |  | [optional] 
**columns** | **List[str]** |  | [optional] 

## Example

```python
from xcherryapi.models.app_database_table_read_request import AppDatabaseTableReadRequest

# TODO update the JSON string below
json = "{}"
# create an instance of AppDatabaseTableReadRequest from a JSON string
app_database_table_read_request_instance = AppDatabaseTableReadRequest.from_json(json)
# print the JSON string representation of the object
print AppDatabaseTableReadRequest.to_json()

# convert the object into a dict
app_database_table_read_request_dict = app_database_table_read_request_instance.to_dict()
# create an instance of AppDatabaseTableReadRequest from a dict
app_database_table_read_request_form_dict = app_database_table_read_request.from_dict(app_database_table_read_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


