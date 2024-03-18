# NotifyReBalancingRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**address** | **str** |  | 
**is_join** | **bool** |  | 

## Example

```python
from xcherryapi.models.notify_re_balancing_request import NotifyReBalancingRequest

# TODO update the JSON string below
json = "{}"
# create an instance of NotifyReBalancingRequest from a JSON string
notify_re_balancing_request_instance = NotifyReBalancingRequest.from_json(json)
# print the JSON string representation of the object
print NotifyReBalancingRequest.to_json()

# convert the object into a dict
notify_re_balancing_request_dict = notify_re_balancing_request_instance.to_dict()
# create an instance of NotifyReBalancingRequest from a dict
notify_re_balancing_request_form_dict = notify_re_balancing_request.from_dict(notify_re_balancing_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


