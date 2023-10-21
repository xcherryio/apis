# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1XdbServiceProcessExecutionDescribePost**](DefaultAPI.md#ApiV1XdbServiceProcessExecutionDescribePost) | **Post** /api/v1/xdb/service/process-execution/describe | describe a process execution
[**ApiV1XdbServiceProcessExecutionPublishToLocalQueuePost**](DefaultAPI.md#ApiV1XdbServiceProcessExecutionPublishToLocalQueuePost) | **Post** /api/v1/xdb/service/process-execution/publish-to-local-queue | send message(s) to be consumed within a single process execution
[**ApiV1XdbServiceProcessExecutionStartPost**](DefaultAPI.md#ApiV1XdbServiceProcessExecutionStartPost) | **Post** /api/v1/xdb/service/process-execution/start | start a process execution
[**ApiV1XdbServiceProcessExecutionStopPost**](DefaultAPI.md#ApiV1XdbServiceProcessExecutionStopPost) | **Post** /api/v1/xdb/service/process-execution/stop | stop a process execution
[**ApiV1XdbWorkerAsyncStateExecutePost**](DefaultAPI.md#ApiV1XdbWorkerAsyncStateExecutePost) | **Post** /api/v1/xdb/worker/async-state/execute | invoking AsyncState.execute API
[**ApiV1XdbWorkerAsyncStateWaitUntilPost**](DefaultAPI.md#ApiV1XdbWorkerAsyncStateWaitUntilPost) | **Post** /api/v1/xdb/worker/async-state/wait-until | invoking AsyncState.waitUntil API
[**InternalApiV1XdbNotifyImmediateTasksPost**](DefaultAPI.md#InternalApiV1XdbNotifyImmediateTasksPost) | **Post** /internal/api/v1/xdb/notify-immediate-tasks | for api service to tell async service that there are new immediate tasks added to the queue
[**InternalApiV1XdbNotifyTimerTasksPost**](DefaultAPI.md#InternalApiV1XdbNotifyTimerTasksPost) | **Post** /internal/api/v1/xdb/notify-timer-tasks | for api service to tell async service that there are new timer tasks added to the queue



## ApiV1XdbServiceProcessExecutionDescribePost

> ProcessExecutionDescribeResponse ApiV1XdbServiceProcessExecutionDescribePost(ctx).ProcessExecutionDescribeRequest(processExecutionDescribeRequest).Execute()

describe a process execution

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xdblab/xdb-apis"
)

func main() {
    processExecutionDescribeRequest := *openapiclient.NewProcessExecutionDescribeRequest("Namespace_example", "ProcessId_example") // ProcessExecutionDescribeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ApiV1XdbServiceProcessExecutionDescribePost(context.Background()).ProcessExecutionDescribeRequest(processExecutionDescribeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1XdbServiceProcessExecutionDescribePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1XdbServiceProcessExecutionDescribePost`: ProcessExecutionDescribeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1XdbServiceProcessExecutionDescribePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1XdbServiceProcessExecutionDescribePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processExecutionDescribeRequest** | [**ProcessExecutionDescribeRequest**](ProcessExecutionDescribeRequest.md) |  | 

### Return type

[**ProcessExecutionDescribeResponse**](ProcessExecutionDescribeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1XdbServiceProcessExecutionPublishToLocalQueuePost

> ApiV1XdbServiceProcessExecutionPublishToLocalQueuePost(ctx).PublishToLocalQueueRequest(publishToLocalQueueRequest).Execute()

send message(s) to be consumed within a single process execution

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xdblab/xdb-apis"
)

func main() {
    publishToLocalQueueRequest := *openapiclient.NewPublishToLocalQueueRequest("Namespace_example", "ProcessId_example") // PublishToLocalQueueRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.ApiV1XdbServiceProcessExecutionPublishToLocalQueuePost(context.Background()).PublishToLocalQueueRequest(publishToLocalQueueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1XdbServiceProcessExecutionPublishToLocalQueuePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1XdbServiceProcessExecutionPublishToLocalQueuePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publishToLocalQueueRequest** | [**PublishToLocalQueueRequest**](PublishToLocalQueueRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1XdbServiceProcessExecutionStartPost

> ProcessExecutionStartResponse ApiV1XdbServiceProcessExecutionStartPost(ctx).ProcessExecutionStartRequest(processExecutionStartRequest).Execute()

start a process execution

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xdblab/xdb-apis"
)

func main() {
    processExecutionStartRequest := *openapiclient.NewProcessExecutionStartRequest("Namespace_example", "ProcessId_example", "ProcessType_example", "WorkerUrl_example") // ProcessExecutionStartRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ApiV1XdbServiceProcessExecutionStartPost(context.Background()).ProcessExecutionStartRequest(processExecutionStartRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1XdbServiceProcessExecutionStartPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1XdbServiceProcessExecutionStartPost`: ProcessExecutionStartResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1XdbServiceProcessExecutionStartPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1XdbServiceProcessExecutionStartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processExecutionStartRequest** | [**ProcessExecutionStartRequest**](ProcessExecutionStartRequest.md) |  | 

### Return type

[**ProcessExecutionStartResponse**](ProcessExecutionStartResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1XdbServiceProcessExecutionStopPost

> ApiV1XdbServiceProcessExecutionStopPost(ctx).ProcessExecutionStopRequest(processExecutionStopRequest).Execute()

stop a process execution

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xdblab/xdb-apis"
)

func main() {
    processExecutionStopRequest := *openapiclient.NewProcessExecutionStopRequest("Namespace_example", "ProcessId_example") // ProcessExecutionStopRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.ApiV1XdbServiceProcessExecutionStopPost(context.Background()).ProcessExecutionStopRequest(processExecutionStopRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1XdbServiceProcessExecutionStopPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1XdbServiceProcessExecutionStopPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processExecutionStopRequest** | [**ProcessExecutionStopRequest**](ProcessExecutionStopRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1XdbWorkerAsyncStateExecutePost

> AsyncStateExecuteResponse ApiV1XdbWorkerAsyncStateExecutePost(ctx).AsyncStateExecuteRequest(asyncStateExecuteRequest).Execute()

invoking AsyncState.execute API

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xdblab/xdb-apis"
)

func main() {
    asyncStateExecuteRequest := *openapiclient.NewAsyncStateExecuteRequest(*openapiclient.NewContext("ProcessId_example", "ProcessExecutionId_example", int64(123)), "ProcessType_example", "StateId_example") // AsyncStateExecuteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ApiV1XdbWorkerAsyncStateExecutePost(context.Background()).AsyncStateExecuteRequest(asyncStateExecuteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1XdbWorkerAsyncStateExecutePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1XdbWorkerAsyncStateExecutePost`: AsyncStateExecuteResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1XdbWorkerAsyncStateExecutePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1XdbWorkerAsyncStateExecutePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asyncStateExecuteRequest** | [**AsyncStateExecuteRequest**](AsyncStateExecuteRequest.md) |  | 

### Return type

[**AsyncStateExecuteResponse**](AsyncStateExecuteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1XdbWorkerAsyncStateWaitUntilPost

> AsyncStateWaitUntilResponse ApiV1XdbWorkerAsyncStateWaitUntilPost(ctx).AsyncStateWaitUntilRequest(asyncStateWaitUntilRequest).Execute()

invoking AsyncState.waitUntil API

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xdblab/xdb-apis"
)

func main() {
    asyncStateWaitUntilRequest := *openapiclient.NewAsyncStateWaitUntilRequest(*openapiclient.NewContext("ProcessId_example", "ProcessExecutionId_example", int64(123)), "ProcessType_example", "StateId_example") // AsyncStateWaitUntilRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ApiV1XdbWorkerAsyncStateWaitUntilPost(context.Background()).AsyncStateWaitUntilRequest(asyncStateWaitUntilRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1XdbWorkerAsyncStateWaitUntilPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1XdbWorkerAsyncStateWaitUntilPost`: AsyncStateWaitUntilResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1XdbWorkerAsyncStateWaitUntilPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1XdbWorkerAsyncStateWaitUntilPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asyncStateWaitUntilRequest** | [**AsyncStateWaitUntilRequest**](AsyncStateWaitUntilRequest.md) |  | 

### Return type

[**AsyncStateWaitUntilResponse**](AsyncStateWaitUntilResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InternalApiV1XdbNotifyImmediateTasksPost

> InternalApiV1XdbNotifyImmediateTasksPost(ctx).NotifyImmediateTasksRequest(notifyImmediateTasksRequest).Execute()

for api service to tell async service that there are new immediate tasks added to the queue

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xdblab/xdb-apis"
)

func main() {
    notifyImmediateTasksRequest := *openapiclient.NewNotifyImmediateTasksRequest(int32(123)) // NotifyImmediateTasksRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.InternalApiV1XdbNotifyImmediateTasksPost(context.Background()).NotifyImmediateTasksRequest(notifyImmediateTasksRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.InternalApiV1XdbNotifyImmediateTasksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInternalApiV1XdbNotifyImmediateTasksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notifyImmediateTasksRequest** | [**NotifyImmediateTasksRequest**](NotifyImmediateTasksRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InternalApiV1XdbNotifyTimerTasksPost

> InternalApiV1XdbNotifyTimerTasksPost(ctx).NotifyTimerTasksRequest(notifyTimerTasksRequest).Execute()

for api service to tell async service that there are new timer tasks added to the queue

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xdblab/xdb-apis"
)

func main() {
    notifyTimerTasksRequest := *openapiclient.NewNotifyTimerTasksRequest(int32(123), []int64{int64(123)}) // NotifyTimerTasksRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.InternalApiV1XdbNotifyTimerTasksPost(context.Background()).NotifyTimerTasksRequest(notifyTimerTasksRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.InternalApiV1XdbNotifyTimerTasksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInternalApiV1XdbNotifyTimerTasksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notifyTimerTasksRequest** | [**NotifyTimerTasksRequest**](NotifyTimerTasksRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

