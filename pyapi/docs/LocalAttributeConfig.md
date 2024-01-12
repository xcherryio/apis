# LocalAttributeConfig


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**initial_write** | [**List[KeyValue]**](KeyValue.md) |  | [optional] 

## Example

```python
from xcherryapi.models.local_attribute_config import LocalAttributeConfig

# TODO update the JSON string below
json = "{}"
# create an instance of LocalAttributeConfig from a JSON string
local_attribute_config_instance = LocalAttributeConfig.from_json(json)
# print the JSON string representation of the object
print LocalAttributeConfig.to_json()

# convert the object into a dict
local_attribute_config_dict = local_attribute_config_instance.to_dict()
# create an instance of LocalAttributeConfig from a dict
local_attribute_config_form_dict = local_attribute_config.from_dict(local_attribute_config_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


