# xcherryapi.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_v1_xcherry_service_process_execution_describe_post**](DefaultApi.md#api_v1_xcherry_service_process_execution_describe_post) | **POST** /api/v1/xcherry/service/process-execution/describe | describe a process execution
[**api_v1_xcherry_service_process_execution_list_post**](DefaultApi.md#api_v1_xcherry_service_process_execution_list_post) | **POST** /api/v1/xcherry/service/process-execution/list | list process executions
[**api_v1_xcherry_service_process_execution_publish_to_local_queue_post**](DefaultApi.md#api_v1_xcherry_service_process_execution_publish_to_local_queue_post) | **POST** /api/v1/xcherry/service/process-execution/publish-to-local-queue | send message(s) to be consumed within a single process execution
[**api_v1_xcherry_service_process_execution_rpc_post**](DefaultApi.md#api_v1_xcherry_service_process_execution_rpc_post) | **POST** /api/v1/xcherry/service/process-execution/rpc | execute a RPC method of a process execution
[**api_v1_xcherry_service_process_execution_start_post**](DefaultApi.md#api_v1_xcherry_service_process_execution_start_post) | **POST** /api/v1/xcherry/service/process-execution/start | start a process execution
[**api_v1_xcherry_service_process_execution_stop_post**](DefaultApi.md#api_v1_xcherry_service_process_execution_stop_post) | **POST** /api/v1/xcherry/service/process-execution/stop | stop a process execution
[**api_v1_xcherry_worker_async_state_execute_post**](DefaultApi.md#api_v1_xcherry_worker_async_state_execute_post) | **POST** /api/v1/xcherry/worker/async-state/execute | invoking AsyncState.execute API
[**api_v1_xcherry_worker_async_state_wait_until_post**](DefaultApi.md#api_v1_xcherry_worker_async_state_wait_until_post) | **POST** /api/v1/xcherry/worker/async-state/wait-until | invoking AsyncState.waitUntil API
[**api_v1_xcherry_worker_process_rpc_post**](DefaultApi.md#api_v1_xcherry_worker_process_rpc_post) | **POST** /api/v1/xcherry/worker/process/rpc | execute a RPC method of a process execution in the worker
[**internal_api_v1_xcherry_notify_immediate_tasks_post**](DefaultApi.md#internal_api_v1_xcherry_notify_immediate_tasks_post) | **POST** /internal/api/v1/xcherry/notify-immediate-tasks | for api service to tell async service that there are new immediate tasks added to the queue
[**internal_api_v1_xcherry_notify_timer_tasks_post**](DefaultApi.md#internal_api_v1_xcherry_notify_timer_tasks_post) | **POST** /internal/api/v1/xcherry/notify-timer-tasks | for api service to tell async service that there are new timer tasks added to the queue


# **api_v1_xcherry_service_process_execution_describe_post**
> ProcessExecutionDescribeResponse api_v1_xcherry_service_process_execution_describe_post(process_execution_describe_request=process_execution_describe_request)

describe a process execution

### Example

```python
import time
import os
import xcherryapi
from xcherryapi.models.process_execution_describe_request import ProcessExecutionDescribeRequest
from xcherryapi.models.process_execution_describe_response import ProcessExecutionDescribeResponse
from xcherryapi.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = xcherryapi.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with xcherryapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = xcherryapi.DefaultApi(api_client)
    process_execution_describe_request = xcherryapi.ProcessExecutionDescribeRequest() # ProcessExecutionDescribeRequest |  (optional)

    try:
        # describe a process execution
        api_response = api_instance.api_v1_xcherry_service_process_execution_describe_post(process_execution_describe_request=process_execution_describe_request)
        print("The response of DefaultApi->api_v1_xcherry_service_process_execution_describe_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DefaultApi->api_v1_xcherry_service_process_execution_describe_post: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **process_execution_describe_request** | [**ProcessExecutionDescribeRequest**](ProcessExecutionDescribeRequest.md)|  | [optional] 

### Return type

[**ProcessExecutionDescribeResponse**](ProcessExecutionDescribeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | successful operation |  -  |
**0** | 400: Invalid request, 404: Process execution not exists |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_v1_xcherry_service_process_execution_list_post**
> ListProcessExecutionsResponse api_v1_xcherry_service_process_execution_list_post(list_process_executions_request=list_process_executions_request)

list process executions

### Example

```python
import time
import os
import xcherryapi
from xcherryapi.models.list_process_executions_request import ListProcessExecutionsRequest
from xcherryapi.models.list_process_executions_response import ListProcessExecutionsResponse
from xcherryapi.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = xcherryapi.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with xcherryapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = xcherryapi.DefaultApi(api_client)
    list_process_executions_request = xcherryapi.ListProcessExecutionsRequest() # ListProcessExecutionsRequest |  (optional)

    try:
        # list process executions
        api_response = api_instance.api_v1_xcherry_service_process_execution_list_post(list_process_executions_request=list_process_executions_request)
        print("The response of DefaultApi->api_v1_xcherry_service_process_execution_list_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DefaultApi->api_v1_xcherry_service_process_execution_list_post: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **list_process_executions_request** | [**ListProcessExecutionsRequest**](ListProcessExecutionsRequest.md)|  | [optional] 

### Return type

[**ListProcessExecutionsResponse**](ListProcessExecutionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | successful operation |  -  |
**0** | 400: Invalid request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_v1_xcherry_service_process_execution_publish_to_local_queue_post**
> api_v1_xcherry_service_process_execution_publish_to_local_queue_post(publish_to_local_queue_request=publish_to_local_queue_request)

send message(s) to be consumed within a single process execution

### Example

```python
import time
import os
import xcherryapi
from xcherryapi.models.publish_to_local_queue_request import PublishToLocalQueueRequest
from xcherryapi.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = xcherryapi.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with xcherryapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = xcherryapi.DefaultApi(api_client)
    publish_to_local_queue_request = xcherryapi.PublishToLocalQueueRequest() # PublishToLocalQueueRequest |  (optional)

    try:
        # send message(s) to be consumed within a single process execution
        api_instance.api_v1_xcherry_service_process_execution_publish_to_local_queue_post(publish_to_local_queue_request=publish_to_local_queue_request)
    except Exception as e:
        print("Exception when calling DefaultApi->api_v1_xcherry_service_process_execution_publish_to_local_queue_post: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publish_to_local_queue_request** | [**PublishToLocalQueueRequest**](PublishToLocalQueueRequest.md)|  | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | successful operation |  -  |
**0** | 400: Invalid request, 404: Process execution not exists or not running |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_v1_xcherry_service_process_execution_rpc_post**
> ProcessExecutionRpcResponse api_v1_xcherry_service_process_execution_rpc_post(process_execution_rpc_request=process_execution_rpc_request)

execute a RPC method of a process execution

### Example

```python
import time
import os
import xcherryapi
from xcherryapi.models.process_execution_rpc_request import ProcessExecutionRpcRequest
from xcherryapi.models.process_execution_rpc_response import ProcessExecutionRpcResponse
from xcherryapi.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = xcherryapi.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with xcherryapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = xcherryapi.DefaultApi(api_client)
    process_execution_rpc_request = xcherryapi.ProcessExecutionRpcRequest() # ProcessExecutionRpcRequest |  (optional)

    try:
        # execute a RPC method of a process execution
        api_response = api_instance.api_v1_xcherry_service_process_execution_rpc_post(process_execution_rpc_request=process_execution_rpc_request)
        print("The response of DefaultApi->api_v1_xcherry_service_process_execution_rpc_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DefaultApi->api_v1_xcherry_service_process_execution_rpc_post: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **process_execution_rpc_request** | [**ProcessExecutionRpcRequest**](ProcessExecutionRpcRequest.md)|  | [optional] 

### Return type

[**ProcessExecutionRpcResponse**](ProcessExecutionRpcResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | successful operation |  -  |
**0** | 400: Invalid request, 404: Process execution not exists, or not running to accept write operation,  424: app database write failure or worker RPC execution failure, see ErrorSubType for details  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_v1_xcherry_service_process_execution_start_post**
> ProcessExecutionStartResponse api_v1_xcherry_service_process_execution_start_post(process_execution_start_request=process_execution_start_request)

start a process execution

### Example

```python
import time
import os
import xcherryapi
from xcherryapi.models.process_execution_start_request import ProcessExecutionStartRequest
from xcherryapi.models.process_execution_start_response import ProcessExecutionStartResponse
from xcherryapi.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = xcherryapi.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with xcherryapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = xcherryapi.DefaultApi(api_client)
    process_execution_start_request = xcherryapi.ProcessExecutionStartRequest() # ProcessExecutionStartRequest |  (optional)

    try:
        # start a process execution
        api_response = api_instance.api_v1_xcherry_service_process_execution_start_post(process_execution_start_request=process_execution_start_request)
        print("The response of DefaultApi->api_v1_xcherry_service_process_execution_start_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DefaultApi->api_v1_xcherry_service_process_execution_start_post: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **process_execution_start_request** | [**ProcessExecutionStartRequest**](ProcessExecutionStartRequest.md)|  | [optional] 

### Return type

[**ProcessExecutionStartResponse**](ProcessExecutionStartResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | successful operation |  -  |
**0** | 400: Invalid request, 409: Process already started, 424: app database write failure |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_v1_xcherry_service_process_execution_stop_post**
> api_v1_xcherry_service_process_execution_stop_post(process_execution_stop_request=process_execution_stop_request)

stop a process execution

### Example

```python
import time
import os
import xcherryapi
from xcherryapi.models.process_execution_stop_request import ProcessExecutionStopRequest
from xcherryapi.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = xcherryapi.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with xcherryapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = xcherryapi.DefaultApi(api_client)
    process_execution_stop_request = xcherryapi.ProcessExecutionStopRequest() # ProcessExecutionStopRequest |  (optional)

    try:
        # stop a process execution
        api_instance.api_v1_xcherry_service_process_execution_stop_post(process_execution_stop_request=process_execution_stop_request)
    except Exception as e:
        print("Exception when calling DefaultApi->api_v1_xcherry_service_process_execution_stop_post: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **process_execution_stop_request** | [**ProcessExecutionStopRequest**](ProcessExecutionStopRequest.md)|  | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | successful operation |  -  |
**0** | 400: Invalid request, 404: Process execution not exists |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_v1_xcherry_worker_async_state_execute_post**
> AsyncStateExecuteResponse api_v1_xcherry_worker_async_state_execute_post(async_state_execute_request=async_state_execute_request)

invoking AsyncState.execute API

### Example

```python
import time
import os
import xcherryapi
from xcherryapi.models.async_state_execute_request import AsyncStateExecuteRequest
from xcherryapi.models.async_state_execute_response import AsyncStateExecuteResponse
from xcherryapi.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = xcherryapi.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with xcherryapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = xcherryapi.DefaultApi(api_client)
    async_state_execute_request = xcherryapi.AsyncStateExecuteRequest() # AsyncStateExecuteRequest |  (optional)

    try:
        # invoking AsyncState.execute API
        api_response = api_instance.api_v1_xcherry_worker_async_state_execute_post(async_state_execute_request=async_state_execute_request)
        print("The response of DefaultApi->api_v1_xcherry_worker_async_state_execute_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DefaultApi->api_v1_xcherry_worker_async_state_execute_post: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **async_state_execute_request** | [**AsyncStateExecuteRequest**](AsyncStateExecuteRequest.md)|  | [optional] 

### Return type

[**AsyncStateExecuteResponse**](AsyncStateExecuteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | successful operation |  -  |
**406** | response of handling app database error |  -  |
**424** | worker execution failure |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_v1_xcherry_worker_async_state_wait_until_post**
> AsyncStateWaitUntilResponse api_v1_xcherry_worker_async_state_wait_until_post(async_state_wait_until_request=async_state_wait_until_request)

invoking AsyncState.waitUntil API

### Example

```python
import time
import os
import xcherryapi
from xcherryapi.models.async_state_wait_until_request import AsyncStateWaitUntilRequest
from xcherryapi.models.async_state_wait_until_response import AsyncStateWaitUntilResponse
from xcherryapi.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = xcherryapi.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with xcherryapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = xcherryapi.DefaultApi(api_client)
    async_state_wait_until_request = xcherryapi.AsyncStateWaitUntilRequest() # AsyncStateWaitUntilRequest |  (optional)

    try:
        # invoking AsyncState.waitUntil API
        api_response = api_instance.api_v1_xcherry_worker_async_state_wait_until_post(async_state_wait_until_request=async_state_wait_until_request)
        print("The response of DefaultApi->api_v1_xcherry_worker_async_state_wait_until_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DefaultApi->api_v1_xcherry_worker_async_state_wait_until_post: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **async_state_wait_until_request** | [**AsyncStateWaitUntilRequest**](AsyncStateWaitUntilRequest.md)|  | [optional] 

### Return type

[**AsyncStateWaitUntilResponse**](AsyncStateWaitUntilResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | successful operation |  -  |
**424** | worker execution failure |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **api_v1_xcherry_worker_process_rpc_post**
> ProcessRpcWorkerResponse api_v1_xcherry_worker_process_rpc_post(process_rpc_worker_request=process_rpc_worker_request)

execute a RPC method of a process execution in the worker

### Example

```python
import time
import os
import xcherryapi
from xcherryapi.models.process_rpc_worker_request import ProcessRpcWorkerRequest
from xcherryapi.models.process_rpc_worker_response import ProcessRpcWorkerResponse
from xcherryapi.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = xcherryapi.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with xcherryapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = xcherryapi.DefaultApi(api_client)
    process_rpc_worker_request = xcherryapi.ProcessRpcWorkerRequest() # ProcessRpcWorkerRequest |  (optional)

    try:
        # execute a RPC method of a process execution in the worker
        api_response = api_instance.api_v1_xcherry_worker_process_rpc_post(process_rpc_worker_request=process_rpc_worker_request)
        print("The response of DefaultApi->api_v1_xcherry_worker_process_rpc_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DefaultApi->api_v1_xcherry_worker_process_rpc_post: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **process_rpc_worker_request** | [**ProcessRpcWorkerRequest**](ProcessRpcWorkerRequest.md)|  | [optional] 

### Return type

[**ProcessRpcWorkerResponse**](ProcessRpcWorkerResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | successful operation |  -  |
**424** | worker execution failure |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **internal_api_v1_xcherry_notify_immediate_tasks_post**
> internal_api_v1_xcherry_notify_immediate_tasks_post(notify_immediate_tasks_request=notify_immediate_tasks_request)

for api service to tell async service that there are new immediate tasks added to the queue

### Example

```python
import time
import os
import xcherryapi
from xcherryapi.models.notify_immediate_tasks_request import NotifyImmediateTasksRequest
from xcherryapi.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = xcherryapi.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with xcherryapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = xcherryapi.DefaultApi(api_client)
    notify_immediate_tasks_request = xcherryapi.NotifyImmediateTasksRequest() # NotifyImmediateTasksRequest |  (optional)

    try:
        # for api service to tell async service that there are new immediate tasks added to the queue
        api_instance.internal_api_v1_xcherry_notify_immediate_tasks_post(notify_immediate_tasks_request=notify_immediate_tasks_request)
    except Exception as e:
        print("Exception when calling DefaultApi->internal_api_v1_xcherry_notify_immediate_tasks_post: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notify_immediate_tasks_request** | [**NotifyImmediateTasksRequest**](NotifyImmediateTasksRequest.md)|  | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | successful operation |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **internal_api_v1_xcherry_notify_timer_tasks_post**
> internal_api_v1_xcherry_notify_timer_tasks_post(notify_timer_tasks_request=notify_timer_tasks_request)

for api service to tell async service that there are new timer tasks added to the queue

### Example

```python
import time
import os
import xcherryapi
from xcherryapi.models.notify_timer_tasks_request import NotifyTimerTasksRequest
from xcherryapi.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = xcherryapi.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with xcherryapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = xcherryapi.DefaultApi(api_client)
    notify_timer_tasks_request = xcherryapi.NotifyTimerTasksRequest() # NotifyTimerTasksRequest |  (optional)

    try:
        # for api service to tell async service that there are new timer tasks added to the queue
        api_instance.internal_api_v1_xcherry_notify_timer_tasks_post(notify_timer_tasks_request=notify_timer_tasks_request)
    except Exception as e:
        print("Exception when calling DefaultApi->internal_api_v1_xcherry_notify_timer_tasks_post: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notify_timer_tasks_request** | [**NotifyTimerTasksRequest**](NotifyTimerTasksRequest.md)|  | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | successful operation |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

