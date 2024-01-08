# ProcessIdFilter


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**process_id** | **str** |  | [optional] 

## Example

```python
from xcherryapi.models.process_id_filter import ProcessIdFilter

# TODO update the JSON string below
json = "{}"
# create an instance of ProcessIdFilter from a JSON string
process_id_filter_instance = ProcessIdFilter.from_json(json)
# print the JSON string representation of the object
print ProcessIdFilter.to_json()

# convert the object into a dict
process_id_filter_dict = process_id_filter_instance.to_dict()
# create an instance of ProcessIdFilter from a dict
process_id_filter_form_dict = process_id_filter.from_dict(process_id_filter_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


