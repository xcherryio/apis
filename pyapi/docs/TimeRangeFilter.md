# TimeRangeFilter


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**earliest_time** | **int** |  | [optional] 
**latest_time** | **int** |  | [optional] 

## Example

```python
from xcherryapi.models.time_range_filter import TimeRangeFilter

# TODO update the JSON string below
json = "{}"
# create an instance of TimeRangeFilter from a JSON string
time_range_filter_instance = TimeRangeFilter.from_json(json)
# print the JSON string representation of the object
print TimeRangeFilter.to_json()

# convert the object into a dict
time_range_filter_dict = time_range_filter_instance.to_dict()
# create an instance of TimeRangeFilter from a dict
time_range_filter_form_dict = time_range_filter.from_dict(time_range_filter_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


