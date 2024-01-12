# EncodedObject


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**encoding** | **str** |  | 
**data** | **str** |  | 

## Example

```python
from xcherryapi.models.encoded_object import EncodedObject

# TODO update the JSON string below
json = "{}"
# create an instance of EncodedObject from a JSON string
encoded_object_instance = EncodedObject.from_json(json)
# print the JSON string representation of the object
print EncodedObject.to_json()

# convert the object into a dict
encoded_object_dict = encoded_object_instance.to_dict()
# create an instance of EncodedObject from a dict
encoded_object_form_dict = encoded_object.from_dict(encoded_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


