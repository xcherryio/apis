# ApiErrorResponse


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**error_sub_type** | [**ErrorSubType**](ErrorSubType.md) |  | [optional] 
**app_error_type** | **str** | for WORKER_EXECUTION_ERROR, it&#39;s the value from WorkerErrorResponse.errorType; for APP_DATABASE_READ/WRITE_ERROR, it&#39;s the error code from database driver  | [optional] 
**details** | **str** | for WORKER_EXECUTION_ERROR, it&#39;s the value from WorkerErrorResponse.details; for APP_DATABASE_READ/WRITE_ERROR, it&#39;s the error message from database driver; for other apiErrorType, it&#39;s the detailed message from server.  | [optional] 

## Example

```python
from xcherryapi.models.api_error_response import ApiErrorResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ApiErrorResponse from a JSON string
api_error_response_instance = ApiErrorResponse.from_json(json)
# print the JSON string representation of the object
print ApiErrorResponse.to_json()

# convert the object into a dict
api_error_response_dict = api_error_response_instance.to_dict()
# create an instance of ApiErrorResponse from a dict
api_error_response_form_dict = api_error_response.from_dict(api_error_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


