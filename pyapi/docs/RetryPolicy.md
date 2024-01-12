# RetryPolicy


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**initial_interval_seconds** | **int** | the initial interval for the first retry, default to 1 second | [optional] 
**backoff_coefficient** | **float** | the backoff coefficient for the next retry, default to 2 | [optional] 
**maximum_interval_seconds** | **int** | the maximum interval for the next retry, default to 100x of initial interval | [optional] 
**maximum_attempts** | **int** | the maximum number of attempts, default to 0, means unlimited | [optional] 
**maximum_attempts_duration_seconds** | **int** | the maximum duration of all attempts, default to 0, means unlimited | [optional] 

## Example

```python
from xcherryapi.models.retry_policy import RetryPolicy

# TODO update the JSON string below
json = "{}"
# create an instance of RetryPolicy from a JSON string
retry_policy_instance = RetryPolicy.from_json(json)
# print the JSON string representation of the object
print RetryPolicy.to_json()

# convert the object into a dict
retry_policy_dict = retry_policy_instance.to_dict()
# create an instance of RetryPolicy from a dict
retry_policy_form_dict = retry_policy.from_dict(retry_policy_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


