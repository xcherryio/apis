/*
XDB APIs

This APIs between xdb service and SDKs

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xdbapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

type DefaultAPI interface {

	/*
		ApiV1XdbServiceProcessExecutionDescribePost describe a process execution

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiApiV1XdbServiceProcessExecutionDescribePostRequest
	*/
	ApiV1XdbServiceProcessExecutionDescribePost(ctx context.Context) ApiApiV1XdbServiceProcessExecutionDescribePostRequest

	// ApiV1XdbServiceProcessExecutionDescribePostExecute executes the request
	//  @return ProcessExecutionDescribeResponse
	ApiV1XdbServiceProcessExecutionDescribePostExecute(r ApiApiV1XdbServiceProcessExecutionDescribePostRequest) (*ProcessExecutionDescribeResponse, *http.Response, error)

	/*
		ApiV1XdbServiceProcessExecutionPublishToLocalQueuePost send message(s) to be consumed within a single process execution

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiApiV1XdbServiceProcessExecutionPublishToLocalQueuePostRequest
	*/
	ApiV1XdbServiceProcessExecutionPublishToLocalQueuePost(ctx context.Context) ApiApiV1XdbServiceProcessExecutionPublishToLocalQueuePostRequest

	// ApiV1XdbServiceProcessExecutionPublishToLocalQueuePostExecute executes the request
	ApiV1XdbServiceProcessExecutionPublishToLocalQueuePostExecute(r ApiApiV1XdbServiceProcessExecutionPublishToLocalQueuePostRequest) (*http.Response, error)

	/*
		ApiV1XdbServiceProcessExecutionStartPost start a process execution

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiApiV1XdbServiceProcessExecutionStartPostRequest
	*/
	ApiV1XdbServiceProcessExecutionStartPost(ctx context.Context) ApiApiV1XdbServiceProcessExecutionStartPostRequest

	// ApiV1XdbServiceProcessExecutionStartPostExecute executes the request
	//  @return ProcessExecutionStartResponse
	ApiV1XdbServiceProcessExecutionStartPostExecute(r ApiApiV1XdbServiceProcessExecutionStartPostRequest) (*ProcessExecutionStartResponse, *http.Response, error)

	/*
		ApiV1XdbServiceProcessExecutionStopPost stop a process execution

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiApiV1XdbServiceProcessExecutionStopPostRequest
	*/
	ApiV1XdbServiceProcessExecutionStopPost(ctx context.Context) ApiApiV1XdbServiceProcessExecutionStopPostRequest

	// ApiV1XdbServiceProcessExecutionStopPostExecute executes the request
	ApiV1XdbServiceProcessExecutionStopPostExecute(r ApiApiV1XdbServiceProcessExecutionStopPostRequest) (*http.Response, error)

	/*
		ApiV1XdbWorkerAsyncStateExecutePost invoking AsyncState.execute API

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiApiV1XdbWorkerAsyncStateExecutePostRequest
	*/
	ApiV1XdbWorkerAsyncStateExecutePost(ctx context.Context) ApiApiV1XdbWorkerAsyncStateExecutePostRequest

	// ApiV1XdbWorkerAsyncStateExecutePostExecute executes the request
	//  @return AsyncStateExecuteResponse
	ApiV1XdbWorkerAsyncStateExecutePostExecute(r ApiApiV1XdbWorkerAsyncStateExecutePostRequest) (*AsyncStateExecuteResponse, *http.Response, error)

	/*
		ApiV1XdbWorkerAsyncStateWaitUntilPost invoking AsyncState.waitUntil API

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiApiV1XdbWorkerAsyncStateWaitUntilPostRequest
	*/
	ApiV1XdbWorkerAsyncStateWaitUntilPost(ctx context.Context) ApiApiV1XdbWorkerAsyncStateWaitUntilPostRequest

	// ApiV1XdbWorkerAsyncStateWaitUntilPostExecute executes the request
	//  @return AsyncStateWaitUntilResponse
	ApiV1XdbWorkerAsyncStateWaitUntilPostExecute(r ApiApiV1XdbWorkerAsyncStateWaitUntilPostRequest) (*AsyncStateWaitUntilResponse, *http.Response, error)

	/*
		InternalApiV1XdbNotifyImmediateTasksPost for api service to tell async service that there are new immediate tasks added to the queue

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiInternalApiV1XdbNotifyImmediateTasksPostRequest
	*/
	InternalApiV1XdbNotifyImmediateTasksPost(ctx context.Context) ApiInternalApiV1XdbNotifyImmediateTasksPostRequest

	// InternalApiV1XdbNotifyImmediateTasksPostExecute executes the request
	InternalApiV1XdbNotifyImmediateTasksPostExecute(r ApiInternalApiV1XdbNotifyImmediateTasksPostRequest) (*http.Response, error)

	/*
		InternalApiV1XdbNotifyTimerTasksPost for api service to tell async service that there are new timer tasks added to the queue

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiInternalApiV1XdbNotifyTimerTasksPostRequest
	*/
	InternalApiV1XdbNotifyTimerTasksPost(ctx context.Context) ApiInternalApiV1XdbNotifyTimerTasksPostRequest

	// InternalApiV1XdbNotifyTimerTasksPostExecute executes the request
	InternalApiV1XdbNotifyTimerTasksPostExecute(r ApiInternalApiV1XdbNotifyTimerTasksPostRequest) (*http.Response, error)
}

// DefaultAPIService DefaultAPI service
type DefaultAPIService service

type ApiApiV1XdbServiceProcessExecutionDescribePostRequest struct {
	ctx                             context.Context
	ApiService                      DefaultAPI
	processExecutionDescribeRequest *ProcessExecutionDescribeRequest
}

func (r ApiApiV1XdbServiceProcessExecutionDescribePostRequest) ProcessExecutionDescribeRequest(processExecutionDescribeRequest ProcessExecutionDescribeRequest) ApiApiV1XdbServiceProcessExecutionDescribePostRequest {
	r.processExecutionDescribeRequest = &processExecutionDescribeRequest
	return r
}

func (r ApiApiV1XdbServiceProcessExecutionDescribePostRequest) Execute() (*ProcessExecutionDescribeResponse, *http.Response, error) {
	return r.ApiService.ApiV1XdbServiceProcessExecutionDescribePostExecute(r)
}

/*
ApiV1XdbServiceProcessExecutionDescribePost describe a process execution

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1XdbServiceProcessExecutionDescribePostRequest
*/
func (a *DefaultAPIService) ApiV1XdbServiceProcessExecutionDescribePost(ctx context.Context) ApiApiV1XdbServiceProcessExecutionDescribePostRequest {
	return ApiApiV1XdbServiceProcessExecutionDescribePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProcessExecutionDescribeResponse
func (a *DefaultAPIService) ApiV1XdbServiceProcessExecutionDescribePostExecute(r ApiApiV1XdbServiceProcessExecutionDescribePostRequest) (*ProcessExecutionDescribeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProcessExecutionDescribeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ApiV1XdbServiceProcessExecutionDescribePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/xdb/service/process-execution/describe"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.processExecutionDescribeRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiV1XdbServiceProcessExecutionPublishToLocalQueuePostRequest struct {
	ctx                        context.Context
	ApiService                 DefaultAPI
	publishToLocalQueueRequest *PublishToLocalQueueRequest
}

func (r ApiApiV1XdbServiceProcessExecutionPublishToLocalQueuePostRequest) PublishToLocalQueueRequest(publishToLocalQueueRequest PublishToLocalQueueRequest) ApiApiV1XdbServiceProcessExecutionPublishToLocalQueuePostRequest {
	r.publishToLocalQueueRequest = &publishToLocalQueueRequest
	return r
}

func (r ApiApiV1XdbServiceProcessExecutionPublishToLocalQueuePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1XdbServiceProcessExecutionPublishToLocalQueuePostExecute(r)
}

/*
ApiV1XdbServiceProcessExecutionPublishToLocalQueuePost send message(s) to be consumed within a single process execution

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1XdbServiceProcessExecutionPublishToLocalQueuePostRequest
*/
func (a *DefaultAPIService) ApiV1XdbServiceProcessExecutionPublishToLocalQueuePost(ctx context.Context) ApiApiV1XdbServiceProcessExecutionPublishToLocalQueuePostRequest {
	return ApiApiV1XdbServiceProcessExecutionPublishToLocalQueuePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *DefaultAPIService) ApiV1XdbServiceProcessExecutionPublishToLocalQueuePostExecute(r ApiApiV1XdbServiceProcessExecutionPublishToLocalQueuePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ApiV1XdbServiceProcessExecutionPublishToLocalQueuePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/xdb/service/process-execution/publish-to-local-queue"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.publishToLocalQueueRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiV1XdbServiceProcessExecutionStartPostRequest struct {
	ctx                          context.Context
	ApiService                   DefaultAPI
	processExecutionStartRequest *ProcessExecutionStartRequest
}

func (r ApiApiV1XdbServiceProcessExecutionStartPostRequest) ProcessExecutionStartRequest(processExecutionStartRequest ProcessExecutionStartRequest) ApiApiV1XdbServiceProcessExecutionStartPostRequest {
	r.processExecutionStartRequest = &processExecutionStartRequest
	return r
}

func (r ApiApiV1XdbServiceProcessExecutionStartPostRequest) Execute() (*ProcessExecutionStartResponse, *http.Response, error) {
	return r.ApiService.ApiV1XdbServiceProcessExecutionStartPostExecute(r)
}

/*
ApiV1XdbServiceProcessExecutionStartPost start a process execution

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1XdbServiceProcessExecutionStartPostRequest
*/
func (a *DefaultAPIService) ApiV1XdbServiceProcessExecutionStartPost(ctx context.Context) ApiApiV1XdbServiceProcessExecutionStartPostRequest {
	return ApiApiV1XdbServiceProcessExecutionStartPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProcessExecutionStartResponse
func (a *DefaultAPIService) ApiV1XdbServiceProcessExecutionStartPostExecute(r ApiApiV1XdbServiceProcessExecutionStartPostRequest) (*ProcessExecutionStartResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProcessExecutionStartResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ApiV1XdbServiceProcessExecutionStartPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/xdb/service/process-execution/start"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.processExecutionStartRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 424 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiV1XdbServiceProcessExecutionStopPostRequest struct {
	ctx                         context.Context
	ApiService                  DefaultAPI
	processExecutionStopRequest *ProcessExecutionStopRequest
}

func (r ApiApiV1XdbServiceProcessExecutionStopPostRequest) ProcessExecutionStopRequest(processExecutionStopRequest ProcessExecutionStopRequest) ApiApiV1XdbServiceProcessExecutionStopPostRequest {
	r.processExecutionStopRequest = &processExecutionStopRequest
	return r
}

func (r ApiApiV1XdbServiceProcessExecutionStopPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1XdbServiceProcessExecutionStopPostExecute(r)
}

/*
ApiV1XdbServiceProcessExecutionStopPost stop a process execution

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1XdbServiceProcessExecutionStopPostRequest
*/
func (a *DefaultAPIService) ApiV1XdbServiceProcessExecutionStopPost(ctx context.Context) ApiApiV1XdbServiceProcessExecutionStopPostRequest {
	return ApiApiV1XdbServiceProcessExecutionStopPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *DefaultAPIService) ApiV1XdbServiceProcessExecutionStopPostExecute(r ApiApiV1XdbServiceProcessExecutionStopPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ApiV1XdbServiceProcessExecutionStopPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/xdb/service/process-execution/stop"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.processExecutionStopRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiV1XdbWorkerAsyncStateExecutePostRequest struct {
	ctx                      context.Context
	ApiService               DefaultAPI
	asyncStateExecuteRequest *AsyncStateExecuteRequest
}

func (r ApiApiV1XdbWorkerAsyncStateExecutePostRequest) AsyncStateExecuteRequest(asyncStateExecuteRequest AsyncStateExecuteRequest) ApiApiV1XdbWorkerAsyncStateExecutePostRequest {
	r.asyncStateExecuteRequest = &asyncStateExecuteRequest
	return r
}

func (r ApiApiV1XdbWorkerAsyncStateExecutePostRequest) Execute() (*AsyncStateExecuteResponse, *http.Response, error) {
	return r.ApiService.ApiV1XdbWorkerAsyncStateExecutePostExecute(r)
}

/*
ApiV1XdbWorkerAsyncStateExecutePost invoking AsyncState.execute API

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1XdbWorkerAsyncStateExecutePostRequest
*/
func (a *DefaultAPIService) ApiV1XdbWorkerAsyncStateExecutePost(ctx context.Context) ApiApiV1XdbWorkerAsyncStateExecutePostRequest {
	return ApiApiV1XdbWorkerAsyncStateExecutePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AsyncStateExecuteResponse
func (a *DefaultAPIService) ApiV1XdbWorkerAsyncStateExecutePostExecute(r ApiApiV1XdbWorkerAsyncStateExecutePostRequest) (*AsyncStateExecuteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AsyncStateExecuteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ApiV1XdbWorkerAsyncStateExecutePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/xdb/worker/async-state/execute"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.asyncStateExecuteRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v WorkerErrorResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiV1XdbWorkerAsyncStateWaitUntilPostRequest struct {
	ctx                        context.Context
	ApiService                 DefaultAPI
	asyncStateWaitUntilRequest *AsyncStateWaitUntilRequest
}

func (r ApiApiV1XdbWorkerAsyncStateWaitUntilPostRequest) AsyncStateWaitUntilRequest(asyncStateWaitUntilRequest AsyncStateWaitUntilRequest) ApiApiV1XdbWorkerAsyncStateWaitUntilPostRequest {
	r.asyncStateWaitUntilRequest = &asyncStateWaitUntilRequest
	return r
}

func (r ApiApiV1XdbWorkerAsyncStateWaitUntilPostRequest) Execute() (*AsyncStateWaitUntilResponse, *http.Response, error) {
	return r.ApiService.ApiV1XdbWorkerAsyncStateWaitUntilPostExecute(r)
}

/*
ApiV1XdbWorkerAsyncStateWaitUntilPost invoking AsyncState.waitUntil API

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1XdbWorkerAsyncStateWaitUntilPostRequest
*/
func (a *DefaultAPIService) ApiV1XdbWorkerAsyncStateWaitUntilPost(ctx context.Context) ApiApiV1XdbWorkerAsyncStateWaitUntilPostRequest {
	return ApiApiV1XdbWorkerAsyncStateWaitUntilPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AsyncStateWaitUntilResponse
func (a *DefaultAPIService) ApiV1XdbWorkerAsyncStateWaitUntilPostExecute(r ApiApiV1XdbWorkerAsyncStateWaitUntilPostRequest) (*AsyncStateWaitUntilResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AsyncStateWaitUntilResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ApiV1XdbWorkerAsyncStateWaitUntilPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/xdb/worker/async-state/wait-until"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.asyncStateWaitUntilRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v WorkerErrorResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiInternalApiV1XdbNotifyImmediateTasksPostRequest struct {
	ctx                         context.Context
	ApiService                  DefaultAPI
	notifyImmediateTasksRequest *NotifyImmediateTasksRequest
}

func (r ApiInternalApiV1XdbNotifyImmediateTasksPostRequest) NotifyImmediateTasksRequest(notifyImmediateTasksRequest NotifyImmediateTasksRequest) ApiInternalApiV1XdbNotifyImmediateTasksPostRequest {
	r.notifyImmediateTasksRequest = &notifyImmediateTasksRequest
	return r
}

func (r ApiInternalApiV1XdbNotifyImmediateTasksPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.InternalApiV1XdbNotifyImmediateTasksPostExecute(r)
}

/*
InternalApiV1XdbNotifyImmediateTasksPost for api service to tell async service that there are new immediate tasks added to the queue

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiInternalApiV1XdbNotifyImmediateTasksPostRequest
*/
func (a *DefaultAPIService) InternalApiV1XdbNotifyImmediateTasksPost(ctx context.Context) ApiInternalApiV1XdbNotifyImmediateTasksPostRequest {
	return ApiInternalApiV1XdbNotifyImmediateTasksPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *DefaultAPIService) InternalApiV1XdbNotifyImmediateTasksPostExecute(r ApiInternalApiV1XdbNotifyImmediateTasksPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.InternalApiV1XdbNotifyImmediateTasksPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/internal/api/v1/xdb/notify-immediate-tasks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.notifyImmediateTasksRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiInternalApiV1XdbNotifyTimerTasksPostRequest struct {
	ctx                     context.Context
	ApiService              DefaultAPI
	notifyTimerTasksRequest *NotifyTimerTasksRequest
}

func (r ApiInternalApiV1XdbNotifyTimerTasksPostRequest) NotifyTimerTasksRequest(notifyTimerTasksRequest NotifyTimerTasksRequest) ApiInternalApiV1XdbNotifyTimerTasksPostRequest {
	r.notifyTimerTasksRequest = &notifyTimerTasksRequest
	return r
}

func (r ApiInternalApiV1XdbNotifyTimerTasksPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.InternalApiV1XdbNotifyTimerTasksPostExecute(r)
}

/*
InternalApiV1XdbNotifyTimerTasksPost for api service to tell async service that there are new timer tasks added to the queue

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiInternalApiV1XdbNotifyTimerTasksPostRequest
*/
func (a *DefaultAPIService) InternalApiV1XdbNotifyTimerTasksPost(ctx context.Context) ApiInternalApiV1XdbNotifyTimerTasksPostRequest {
	return ApiInternalApiV1XdbNotifyTimerTasksPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *DefaultAPIService) InternalApiV1XdbNotifyTimerTasksPostExecute(r ApiInternalApiV1XdbNotifyTimerTasksPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.InternalApiV1XdbNotifyTimerTasksPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/internal/api/v1/xdb/notify-timer-tasks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.notifyTimerTasksRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
