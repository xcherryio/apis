# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1XdbServiceProcessExecutionDesribePost**](DefaultAPI.md#ApiV1XdbServiceProcessExecutionDesribePost) | **Post** /api/v1/xdb/service/process-execution/desribe | describe a process execution
[**ApiV1XdbServiceProcessExecutionStartPost**](DefaultAPI.md#ApiV1XdbServiceProcessExecutionStartPost) | **Post** /api/v1/xdb/service/process-execution/start | start a process execution
[**ApiV1XdbWorkerAsyncStateExecutePost**](DefaultAPI.md#ApiV1XdbWorkerAsyncStateExecutePost) | **Post** /api/v1/xdb/worker/async-state/execute | invoking AsyncState.execute API
[**ApiV1XdbWorkerAsyncStateWaitUntilPost**](DefaultAPI.md#ApiV1XdbWorkerAsyncStateWaitUntilPost) | **Post** /api/v1/xdb/worker/async-state/wait-until | invoking AsyncState.waitUntil API



## ApiV1XdbServiceProcessExecutionDesribePost

> ProcessExecutionDescribeResponse ApiV1XdbServiceProcessExecutionDesribePost(ctx).ProcessExecutionDescribeRequest(processExecutionDescribeRequest).Execute()

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
    processExecutionDescribeRequest := *openapiclient.NewProcessExecutionDescribeRequest() // ProcessExecutionDescribeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ApiV1XdbServiceProcessExecutionDesribePost(context.Background()).ProcessExecutionDescribeRequest(processExecutionDescribeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1XdbServiceProcessExecutionDesribePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1XdbServiceProcessExecutionDesribePost`: ProcessExecutionDescribeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1XdbServiceProcessExecutionDesribePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1XdbServiceProcessExecutionDesribePostRequest struct via the builder pattern


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
    processExecutionStartRequest := *openapiclient.NewProcessExecutionStartRequest("ProcessId_example", "ProcessType_example", "WorkerUrl_example") // ProcessExecutionStartRequest |  (optional)

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
    asyncStateExecuteRequest := *openapiclient.NewAsyncStateExecuteRequest() // AsyncStateExecuteRequest |  (optional)

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
    asyncStateWaitUntilRequest := *openapiclient.NewAsyncStateWaitUntilRequest() // AsyncStateWaitUntilRequest |  (optional)

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

