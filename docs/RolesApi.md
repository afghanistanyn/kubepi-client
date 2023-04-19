# \RolesApi

All URIs are relative to *https://localhost/kubepi/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RolesGet**](RolesApi.md#RolesGet) | **Get** /roles | List all roles
[**RolesNameDelete**](RolesApi.md#RolesNameDelete) | **Delete** /roles/{name} | Delete role by name
[**RolesNameGet**](RolesApi.md#RolesNameGet) | **Get** /roles/{name} | Get role by name
[**RolesNamePut**](RolesApi.md#RolesNamePut) | **Put** /roles/{name} | Update role by name
[**RolesPost**](RolesApi.md#RolesPost) | **Post** /roles | Create role


# **RolesGet**
> []RoleRole RolesGet(ctx, )
List all roles

List all roles

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]RoleRole**](role.Role.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RolesNameDelete**
> RoleRole RolesNameDelete(ctx, name)
Delete role by name

Delete role by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| 角色名称 | 

### Return type

[**RoleRole**](role.Role.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RolesNameGet**
> RoleRole RolesNameGet(ctx, name)
Get role by name

Get role by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| 权限名称 | 

### Return type

[**RoleRole**](role.Role.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RolesNamePut**
> RoleRole RolesNamePut(ctx, name)
Update role by name

Update role by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| 角色名称 | 

### Return type

[**RoleRole**](role.Role.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RolesPost**
> RoleRole RolesPost(ctx, request)
Create role

Create role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**RoleRole**](RoleRole.md)| request | 

### Return type

[**RoleRole**](role.Role.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

