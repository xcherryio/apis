# LoadLocalAttributesRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**keys_to_load_no_lock** | **List[str]** |  | [optional] 
**keys_to_load_with_lock** | **List[str]** |  | [optional] 
**lock_type** | [**DatabaseLockingType**](DatabaseLockingType.md) |  | [optional] 

## Example

```python
from xcherryapi.models.load_local_attributes_request import LoadLocalAttributesRequest

# TODO update the JSON string below
json = "{}"
# create an instance of LoadLocalAttributesRequest from a JSON string
load_local_attributes_request_instance = LoadLocalAttributesRequest.from_json(json)
# print the JSON string representation of the object
print LoadLocalAttributesRequest.to_json()

# convert the object into a dict
load_local_attributes_request_dict = load_local_attributes_request_instance.to_dict()
# create an instance of LoadLocalAttributesRequest from a dict
load_local_attributes_request_form_dict = load_local_attributes_request.from_dict(load_local_attributes_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


