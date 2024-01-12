# AppDatabaseRowReadResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**columns** | [**List[AppDatabaseColumnValue]**](AppDatabaseColumnValue.md) |  | [optional] 

## Example

```python
from xcherryapi.models.app_database_row_read_response import AppDatabaseRowReadResponse

# TODO update the JSON string below
json = "{}"
# create an instance of AppDatabaseRowReadResponse from a JSON string
app_database_row_read_response_instance = AppDatabaseRowReadResponse.from_json(json)
# print the JSON string representation of the object
print AppDatabaseRowReadResponse.to_json()

# convert the object into a dict
app_database_row_read_response_dict = app_database_row_read_response_instance.to_dict()
# create an instance of AppDatabaseRowReadResponse from a dict
app_database_row_read_response_form_dict = app_database_row_read_response.from_dict(app_database_row_read_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


