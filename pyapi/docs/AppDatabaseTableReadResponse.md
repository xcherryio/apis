# AppDatabaseTableReadResponse


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**table_name** | **str** |  | [optional] 
**rows** | [**List[AppDatabaseRowReadResponse]**](AppDatabaseRowReadResponse.md) |  | [optional] 

## Example

```python
from xcherryapi.models.app_database_table_read_response import AppDatabaseTableReadResponse

# TODO update the JSON string below
json = "{}"
# create an instance of AppDatabaseTableReadResponse from a JSON string
app_database_table_read_response_instance = AppDatabaseTableReadResponse.from_json(json)
# print the JSON string representation of the object
print AppDatabaseTableReadResponse.to_json()

# convert the object into a dict
app_database_table_read_response_dict = app_database_table_read_response_instance.to_dict()
# create an instance of AppDatabaseTableReadResponse from a dict
app_database_table_read_response_form_dict = app_database_table_read_response.from_dict(app_database_table_read_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


