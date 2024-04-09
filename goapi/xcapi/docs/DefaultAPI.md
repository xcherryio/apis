# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1XcherryServiceProcessExecutionDescribePost**](DefaultAPI.md#ApiV1XcherryServiceProcessExecutionDescribePost) | **Post** /api/v1/xcherry/service/process-execution/describe | describe a process execution
[**ApiV1XcherryServiceProcessExecutionListPost**](DefaultAPI.md#ApiV1XcherryServiceProcessExecutionListPost) | **Post** /api/v1/xcherry/service/process-execution/list | list process executions
[**ApiV1XcherryServiceProcessExecutionPublishToLocalQueuePost**](DefaultAPI.md#ApiV1XcherryServiceProcessExecutionPublishToLocalQueuePost) | **Post** /api/v1/xcherry/service/process-execution/publish-to-local-queue | send message(s) to be consumed within a single process execution
[**ApiV1XcherryServiceProcessExecutionRpcPost**](DefaultAPI.md#ApiV1XcherryServiceProcessExecutionRpcPost) | **Post** /api/v1/xcherry/service/process-execution/rpc | execute a RPC method of a process execution
[**ApiV1XcherryServiceProcessExecutionStartPost**](DefaultAPI.md#ApiV1XcherryServiceProcessExecutionStartPost) | **Post** /api/v1/xcherry/service/process-execution/start | start a process execution
[**ApiV1XcherryServiceProcessExecutionStopPost**](DefaultAPI.md#ApiV1XcherryServiceProcessExecutionStopPost) | **Post** /api/v1/xcherry/service/process-execution/stop | stop a process execution
[**ApiV1XcherryServiceProcessExecutionWaitForProcessCompletionPost**](DefaultAPI.md#ApiV1XcherryServiceProcessExecutionWaitForProcessCompletionPost) | **Post** /api/v1/xcherry/service/process-execution/wait-for-process-completion | wait for a process completion
[**ApiV1XcherryWorkerAsyncStateExecutePost**](DefaultAPI.md#ApiV1XcherryWorkerAsyncStateExecutePost) | **Post** /api/v1/xcherry/worker/async-state/execute | invoking AsyncState.execute API
[**ApiV1XcherryWorkerAsyncStateWaitUntilPost**](DefaultAPI.md#ApiV1XcherryWorkerAsyncStateWaitUntilPost) | **Post** /api/v1/xcherry/worker/async-state/wait-until | invoking AsyncState.waitUntil API
[**ApiV1XcherryWorkerProcessRpcPost**](DefaultAPI.md#ApiV1XcherryWorkerProcessRpcPost) | **Post** /api/v1/xcherry/worker/process/rpc | execute a RPC method of a process execution in the worker
[**InternalApiV1XcherryNotifyImmediateTasksPost**](DefaultAPI.md#InternalApiV1XcherryNotifyImmediateTasksPost) | **Post** /internal/api/v1/xcherry/notify-immediate-tasks | for api service to tell async service that there are new immediate tasks added to the queue
[**InternalApiV1XcherryNotifyTimerTasksPost**](DefaultAPI.md#InternalApiV1XcherryNotifyTimerTasksPost) | **Post** /internal/api/v1/xcherry/notify-timer-tasks | for api service to tell async service that there are new timer tasks added to the queue
[**InternalApiV1XcherrySignalProcessCompletionPost**](DefaultAPI.md#InternalApiV1XcherrySignalProcessCompletionPost) | **Post** /internal/api/v1/xcherry/signal-process-completion | for async service to signal for process completion
[**InternalApiV1XcherryWaitForProcessCompletionPost**](DefaultAPI.md#InternalApiV1XcherryWaitForProcessCompletionPost) | **Post** /internal/api/v1/xcherry/wait-for-process-completion | for api service to ask async service to wait for process completion



## ApiV1XcherryServiceProcessExecutionDescribePost

> ProcessExecutionDescribeResponse ApiV1XcherryServiceProcessExecutionDescribePost(ctx).ProcessExecutionDescribeRequest(processExecutionDescribeRequest).Execute()

describe a process execution

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xcherryio/apis"
)

func main() {
    processExecutionDescribeRequest := *openapiclient.NewProcessExecutionDescribeRequest("Namespace_example", "ProcessId_example") // ProcessExecutionDescribeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ApiV1XcherryServiceProcessExecutionDescribePost(context.Background()).ProcessExecutionDescribeRequest(processExecutionDescribeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1XcherryServiceProcessExecutionDescribePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1XcherryServiceProcessExecutionDescribePost`: ProcessExecutionDescribeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1XcherryServiceProcessExecutionDescribePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1XcherryServiceProcessExecutionDescribePostRequest struct via the builder pattern


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


## ApiV1XcherryServiceProcessExecutionListPost

> ListProcessExecutionsResponse ApiV1XcherryServiceProcessExecutionListPost(ctx).ListProcessExecutionsRequest(listProcessExecutionsRequest).Execute()

list process executions

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xcherryio/apis"
)

func main() {
    listProcessExecutionsRequest := *openapiclient.NewListProcessExecutionsRequest("Namespace_example", int32(123)) // ListProcessExecutionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ApiV1XcherryServiceProcessExecutionListPost(context.Background()).ListProcessExecutionsRequest(listProcessExecutionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1XcherryServiceProcessExecutionListPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1XcherryServiceProcessExecutionListPost`: ListProcessExecutionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1XcherryServiceProcessExecutionListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1XcherryServiceProcessExecutionListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listProcessExecutionsRequest** | [**ListProcessExecutionsRequest**](ListProcessExecutionsRequest.md) |  | 

### Return type

[**ListProcessExecutionsResponse**](ListProcessExecutionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1XcherryServiceProcessExecutionPublishToLocalQueuePost

> ApiV1XcherryServiceProcessExecutionPublishToLocalQueuePost(ctx).PublishToLocalQueueRequest(publishToLocalQueueRequest).Execute()

send message(s) to be consumed within a single process execution

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xcherryio/apis"
)

func main() {
    publishToLocalQueueRequest := *openapiclient.NewPublishToLocalQueueRequest("Namespace_example", "ProcessId_example") // PublishToLocalQueueRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.ApiV1XcherryServiceProcessExecutionPublishToLocalQueuePost(context.Background()).PublishToLocalQueueRequest(publishToLocalQueueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1XcherryServiceProcessExecutionPublishToLocalQueuePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1XcherryServiceProcessExecutionPublishToLocalQueuePostRequest struct via the builder pattern


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


## ApiV1XcherryServiceProcessExecutionRpcPost

> ProcessExecutionRpcResponse ApiV1XcherryServiceProcessExecutionRpcPost(ctx).ProcessExecutionRpcRequest(processExecutionRpcRequest).Execute()

execute a RPC method of a process execution

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xcherryio/apis"
)

func main() {
    processExecutionRpcRequest := *openapiclient.NewProcessExecutionRpcRequest("Namespace_example", "ProcessId_example", "RpcName_example") // ProcessExecutionRpcRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ApiV1XcherryServiceProcessExecutionRpcPost(context.Background()).ProcessExecutionRpcRequest(processExecutionRpcRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1XcherryServiceProcessExecutionRpcPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1XcherryServiceProcessExecutionRpcPost`: ProcessExecutionRpcResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1XcherryServiceProcessExecutionRpcPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1XcherryServiceProcessExecutionRpcPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processExecutionRpcRequest** | [**ProcessExecutionRpcRequest**](ProcessExecutionRpcRequest.md) |  | 

### Return type

[**ProcessExecutionRpcResponse**](ProcessExecutionRpcResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1XcherryServiceProcessExecutionStartPost

> ProcessExecutionStartResponse ApiV1XcherryServiceProcessExecutionStartPost(ctx).ProcessExecutionStartRequest(processExecutionStartRequest).Execute()

start a process execution

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xcherryio/apis"
)

func main() {
    processExecutionStartRequest := *openapiclient.NewProcessExecutionStartRequest("Namespace_example", "ProcessId_example", "ProcessType_example", "WorkerUrl_example") // ProcessExecutionStartRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ApiV1XcherryServiceProcessExecutionStartPost(context.Background()).ProcessExecutionStartRequest(processExecutionStartRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1XcherryServiceProcessExecutionStartPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1XcherryServiceProcessExecutionStartPost`: ProcessExecutionStartResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1XcherryServiceProcessExecutionStartPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1XcherryServiceProcessExecutionStartPostRequest struct via the builder pattern


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


## ApiV1XcherryServiceProcessExecutionStopPost

> ApiV1XcherryServiceProcessExecutionStopPost(ctx).ProcessExecutionStopRequest(processExecutionStopRequest).Execute()

stop a process execution

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xcherryio/apis"
)

func main() {
    processExecutionStopRequest := *openapiclient.NewProcessExecutionStopRequest("Namespace_example", "ProcessId_example") // ProcessExecutionStopRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.ApiV1XcherryServiceProcessExecutionStopPost(context.Background()).ProcessExecutionStopRequest(processExecutionStopRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1XcherryServiceProcessExecutionStopPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1XcherryServiceProcessExecutionStopPostRequest struct via the builder pattern


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


## ApiV1XcherryServiceProcessExecutionWaitForProcessCompletionPost

> ProcessExecutionWaitForCompletionResponse ApiV1XcherryServiceProcessExecutionWaitForProcessCompletionPost(ctx).ProcessExecutionWaitForCompletionRequest(processExecutionWaitForCompletionRequest).Execute()

wait for a process completion

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xcherryio/apis"
)

func main() {
    processExecutionWaitForCompletionRequest := *openapiclient.NewProcessExecutionWaitForCompletionRequest("Namespace_example", "ProcessId_example") // ProcessExecutionWaitForCompletionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ApiV1XcherryServiceProcessExecutionWaitForProcessCompletionPost(context.Background()).ProcessExecutionWaitForCompletionRequest(processExecutionWaitForCompletionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1XcherryServiceProcessExecutionWaitForProcessCompletionPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1XcherryServiceProcessExecutionWaitForProcessCompletionPost`: ProcessExecutionWaitForCompletionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1XcherryServiceProcessExecutionWaitForProcessCompletionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1XcherryServiceProcessExecutionWaitForProcessCompletionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processExecutionWaitForCompletionRequest** | [**ProcessExecutionWaitForCompletionRequest**](ProcessExecutionWaitForCompletionRequest.md) |  | 

### Return type

[**ProcessExecutionWaitForCompletionResponse**](ProcessExecutionWaitForCompletionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1XcherryWorkerAsyncStateExecutePost

> AsyncStateExecuteResponse ApiV1XcherryWorkerAsyncStateExecutePost(ctx).AsyncStateExecuteRequest(asyncStateExecuteRequest).Execute()

invoking AsyncState.execute API

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xcherryio/apis"
)

func main() {
    asyncStateExecuteRequest := *openapiclient.NewAsyncStateExecuteRequest(*openapiclient.NewContext("ProcessId_example", "ProcessExecutionId_example", int64(123)), "ProcessType_example", "StateId_example") // AsyncStateExecuteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ApiV1XcherryWorkerAsyncStateExecutePost(context.Background()).AsyncStateExecuteRequest(asyncStateExecuteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1XcherryWorkerAsyncStateExecutePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1XcherryWorkerAsyncStateExecutePost`: AsyncStateExecuteResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1XcherryWorkerAsyncStateExecutePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1XcherryWorkerAsyncStateExecutePostRequest struct via the builder pattern


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


## ApiV1XcherryWorkerAsyncStateWaitUntilPost

> AsyncStateWaitUntilResponse ApiV1XcherryWorkerAsyncStateWaitUntilPost(ctx).AsyncStateWaitUntilRequest(asyncStateWaitUntilRequest).Execute()

invoking AsyncState.waitUntil API

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xcherryio/apis"
)

func main() {
    asyncStateWaitUntilRequest := *openapiclient.NewAsyncStateWaitUntilRequest(*openapiclient.NewContext("ProcessId_example", "ProcessExecutionId_example", int64(123)), "ProcessType_example", "StateId_example") // AsyncStateWaitUntilRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ApiV1XcherryWorkerAsyncStateWaitUntilPost(context.Background()).AsyncStateWaitUntilRequest(asyncStateWaitUntilRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1XcherryWorkerAsyncStateWaitUntilPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1XcherryWorkerAsyncStateWaitUntilPost`: AsyncStateWaitUntilResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1XcherryWorkerAsyncStateWaitUntilPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1XcherryWorkerAsyncStateWaitUntilPostRequest struct via the builder pattern


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


## ApiV1XcherryWorkerProcessRpcPost

> ProcessRpcWorkerResponse ApiV1XcherryWorkerProcessRpcPost(ctx).ProcessRpcWorkerRequest(processRpcWorkerRequest).Execute()

execute a RPC method of a process execution in the worker

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xcherryio/apis"
)

func main() {
    processRpcWorkerRequest := *openapiclient.NewProcessRpcWorkerRequest(*openapiclient.NewContext("ProcessId_example", "ProcessExecutionId_example", int64(123)), "ProcessType_example", "RpcName_example") // ProcessRpcWorkerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ApiV1XcherryWorkerProcessRpcPost(context.Background()).ProcessRpcWorkerRequest(processRpcWorkerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1XcherryWorkerProcessRpcPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1XcherryWorkerProcessRpcPost`: ProcessRpcWorkerResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1XcherryWorkerProcessRpcPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1XcherryWorkerProcessRpcPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processRpcWorkerRequest** | [**ProcessRpcWorkerRequest**](ProcessRpcWorkerRequest.md) |  | 

### Return type

[**ProcessRpcWorkerResponse**](ProcessRpcWorkerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InternalApiV1XcherryNotifyImmediateTasksPost

> InternalApiV1XcherryNotifyImmediateTasksPost(ctx).NotifyImmediateTasksRequest(notifyImmediateTasksRequest).Execute()

for api service to tell async service that there are new immediate tasks added to the queue

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xcherryio/apis"
)

func main() {
    notifyImmediateTasksRequest := *openapiclient.NewNotifyImmediateTasksRequest(int32(123)) // NotifyImmediateTasksRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.InternalApiV1XcherryNotifyImmediateTasksPost(context.Background()).NotifyImmediateTasksRequest(notifyImmediateTasksRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.InternalApiV1XcherryNotifyImmediateTasksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInternalApiV1XcherryNotifyImmediateTasksPostRequest struct via the builder pattern


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


## InternalApiV1XcherryNotifyTimerTasksPost

> InternalApiV1XcherryNotifyTimerTasksPost(ctx).NotifyTimerTasksRequest(notifyTimerTasksRequest).Execute()

for api service to tell async service that there are new timer tasks added to the queue

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xcherryio/apis"
)

func main() {
    notifyTimerTasksRequest := *openapiclient.NewNotifyTimerTasksRequest(int32(123), []int64{int64(123)}) // NotifyTimerTasksRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.InternalApiV1XcherryNotifyTimerTasksPost(context.Background()).NotifyTimerTasksRequest(notifyTimerTasksRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.InternalApiV1XcherryNotifyTimerTasksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInternalApiV1XcherryNotifyTimerTasksPostRequest struct via the builder pattern


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


## InternalApiV1XcherrySignalProcessCompletionPost

> InternalApiV1XcherrySignalProcessCompletionPost(ctx).SignalProcessCompletionRequest(signalProcessCompletionRequest).Execute()

for async service to signal for process completion

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xcherryio/apis"
)

func main() {
    signalProcessCompletionRequest := *openapiclient.NewSignalProcessCompletionRequest(int32(123), "ProcessExecutionId_example", "Status_example") // SignalProcessCompletionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.InternalApiV1XcherrySignalProcessCompletionPost(context.Background()).SignalProcessCompletionRequest(signalProcessCompletionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.InternalApiV1XcherrySignalProcessCompletionPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInternalApiV1XcherrySignalProcessCompletionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signalProcessCompletionRequest** | [**SignalProcessCompletionRequest**](SignalProcessCompletionRequest.md) |  | 

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


## InternalApiV1XcherryWaitForProcessCompletionPost

> WaitForProcessCompletionResponse InternalApiV1XcherryWaitForProcessCompletionPost(ctx).WaitForProcessCompletionRequest(waitForProcessCompletionRequest).Execute()

for api service to ask async service to wait for process completion

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/xcherryio/apis"
)

func main() {
    waitForProcessCompletionRequest := *openapiclient.NewWaitForProcessCompletionRequest(int32(123), "ProcessExecutionId_example") // WaitForProcessCompletionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.InternalApiV1XcherryWaitForProcessCompletionPost(context.Background()).WaitForProcessCompletionRequest(waitForProcessCompletionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.InternalApiV1XcherryWaitForProcessCompletionPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InternalApiV1XcherryWaitForProcessCompletionPost`: WaitForProcessCompletionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.InternalApiV1XcherryWaitForProcessCompletionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInternalApiV1XcherryWaitForProcessCompletionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **waitForProcessCompletionRequest** | [**WaitForProcessCompletionRequest**](WaitForProcessCompletionRequest.md) |  | 

### Return type

[**WaitForProcessCompletionResponse**](WaitForProcessCompletionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

