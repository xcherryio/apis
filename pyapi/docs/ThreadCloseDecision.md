# ThreadCloseDecision


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**close_type** | [**ThreadCloseType**](ThreadCloseType.md) |  | 
**close_input** | [**EncodedObject**](EncodedObject.md) |  | [optional] 

## Example

```python
from xcherryapi.models.thread_close_decision import ThreadCloseDecision

# TODO update the JSON string below
json = "{}"
# create an instance of ThreadCloseDecision from a JSON string
thread_close_decision_instance = ThreadCloseDecision.from_json(json)
# print the JSON string representation of the object
print ThreadCloseDecision.to_json()

# convert the object into a dict
thread_close_decision_dict = thread_close_decision_instance.to_dict()
# create an instance of ThreadCloseDecision from a dict
thread_close_decision_form_dict = thread_close_decision.from_dict(thread_close_decision_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


