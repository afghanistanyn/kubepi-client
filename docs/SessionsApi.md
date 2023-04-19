# \SessionsApi

All URIs are relative to *https://localhost/kubepi/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SessionsPost**](SessionsApi.md#SessionsPost) | **Post** /sessions | User Login


# **SessionsPost**
> SessionsPost(ctx, request)
User Login

User Login

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**SessionLoginCredential**](SessionLoginCredential.md)| request | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

