# AppDatabaseColumnValue

the value of a table column (from SDK to server or from server to SDK)

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**column** | **str** | the column name that can be used in the database query, see below for example | 
**query_value** | **str** | the plain string value that can be used in the database query(e.g. for SQL SELECT ... WHERE $Column&#x3D;$dbQueryValue or UPDATE/INSERT) | 

## Example

```python
from xcherryapi.models.app_database_column_value import AppDatabaseColumnValue

# TODO update the JSON string below
json = "{}"
# create an instance of AppDatabaseColumnValue from a JSON string
app_database_column_value_instance = AppDatabaseColumnValue.from_json(json)
# print the JSON string representation of the object
print AppDatabaseColumnValue.to_json()

# convert the object into a dict
app_database_column_value_dict = app_database_column_value_instance.to_dict()
# create an instance of AppDatabaseColumnValue from a dict
app_database_column_value_form_dict = app_database_column_value.from_dict(app_database_column_value_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


