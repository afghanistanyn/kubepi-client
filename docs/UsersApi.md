# \UsersApi

All URIs are relative to *https://localhost/kubepi/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersGet**](UsersApi.md#UsersGet) | **Get** /users | List all users
[**UsersNameDelete**](UsersApi.md#UsersNameDelete) | **Delete** /users/{name} | Delete user by name
[**UsersNameGet**](UsersApi.md#UsersNameGet) | **Get** /users/{name} | Get user by name
[**UsersNamePut**](UsersApi.md#UsersNamePut) | **Put** /users/{name} | Update user by name
[**UsersPost**](UsersApi.md#UsersPost) | **Post** /users | Create user
[**UsersSearchPost**](UsersApi.md#UsersSearchPost) | **Post** /users/search | Search users


# **UsersGet**
> []KubepiUser UsersGet(ctx, )
List all users

List all users

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]KubepiUser**](user.User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNameDelete**
> KubepiUser UsersNameDelete(ctx, name)
Delete user by name

Delete user by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| 用户名称 | 

### Return type

[**KubepiUser**](user.User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNameGet**
> KubepiUser UsersNameGet(ctx, name)
Get user by name

Get user by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| 用户名称 | 

### Return type

[**KubepiUser**](user.User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNamePut**
> KubepiUser UsersNamePut(ctx, name)
Update user by name

Update user by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| 用户名称 | 

### Return type

[**KubepiUser**](user.User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPost**
> KubepiUser UsersPost(ctx, request)
Create user

Create user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**DocsUserCreate**](DocsUserCreate.md)| request | 

### Return type

[**KubepiUser**](user.User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersSearchPost**
> ApiPage UsersSearchPost(ctx, )
Search users

Search users by Condition

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiPage**](api.Page.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

