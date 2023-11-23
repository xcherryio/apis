# LoadLocalAttributesResponse


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**attributes** | [**List[KeyValue]**](KeyValue.md) |  | [optional] 

## Example

```python
from xcherryapi.models.load_local_attributes_response import LoadLocalAttributesResponse

# TODO update the JSON string below
json = "{}"
# create an instance of LoadLocalAttributesResponse from a JSON string
load_local_attributes_response_instance = LoadLocalAttributesResponse.from_json(json)
# print the JSON string representation of the object
print LoadLocalAttributesResponse.to_json()

# convert the object into a dict
load_local_attributes_response_dict = load_local_attributes_response_instance.to_dict()
# create an instance of LoadLocalAttributesResponse from a dict
load_local_attributes_response_form_dict = load_local_attributes_response.from_dict(load_local_attributes_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


