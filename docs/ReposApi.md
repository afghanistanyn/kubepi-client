# \ReposApi

All URIs are relative to *https://localhost/kubepi/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImagereposClusterClusterGet**](ReposApi.md#ImagereposClusterClusterGet) | **Get** /imagerepos/cluster/{cluster} | List repo for cluster
[**ImagereposClusterRepoPut**](ReposApi.md#ImagereposClusterRepoPut) | **Put** /imagerepos/{cluster}/{repo} | List images for cluster
[**ImagereposNameDelete**](ReposApi.md#ImagereposNameDelete) | **Delete** /imagerepos/{name} | Delete repo by name
[**ImagereposNameGet**](ReposApi.md#ImagereposNameGet) | **Get** /imagerepos/{name} | Get repo by name
[**ImagereposNamePut**](ReposApi.md#ImagereposNamePut) | **Put** /imagerepos/{name} | Update repo by name
[**ImagereposPost**](ReposApi.md#ImagereposPost) | **Post** /imagerepos | Create repo
[**ImagereposRepoSearchGet**](ReposApi.md#ImagereposRepoSearchGet) | **Get** /imagerepos/{repo}/search | List images by repo
[**ImagereposRepositoriesSearchPost**](ReposApi.md#ImagereposRepositoriesSearchPost) | **Post** /imagerepos/repositories/search | List Internal repos


# **ImagereposClusterClusterGet**
> []ImagerepoImageRepo ImagereposClusterClusterGet(ctx, cluster)
List repo for cluster

Get repo for cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cluster** | **string**| 集群名称 | 

### Return type

[**[]ImagerepoImageRepo**](imagerepo.ImageRepo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImagereposClusterRepoPut**
> []string ImagereposClusterRepoPut(ctx, cluster, repo)
List images for cluster

Get images for cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cluster** | **string**| 集群名称 | 
  **repo** | **string**| 镜像仓库名称 | 

### Return type

**[]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImagereposNameDelete**
> ImagerepoImageRepo ImagereposNameDelete(ctx, name)
Delete repo by name

Delete repo by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| 镜像仓库名称 | 

### Return type

[**ImagerepoImageRepo**](imagerepo.ImageRepo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImagereposNameGet**
> ImagerepoImageRepo ImagereposNameGet(ctx, name)
Get repo by name

Get repo by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| 镜像仓库名称 | 

### Return type

[**ImagerepoImageRepo**](imagerepo.ImageRepo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImagereposNamePut**
> ImagerepoImageRepo ImagereposNamePut(ctx, request, name)
Update repo by name

Update repo by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**ImagerepoRepoConfig**](ImagerepoRepoConfig.md)| request | 
  **name** | **string**| 镜像仓库名称 | 

### Return type

[**ImagerepoImageRepo**](imagerepo.ImageRepo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImagereposPost**
> ImagerepoRepoConfig ImagereposPost(ctx, request)
Create repo

Create repo

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**ImagerepoRepoConfig**](ImagerepoRepoConfig.md)| request | 

### Return type

[**ImagerepoRepoConfig**](imagerepo.RepoConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImagereposRepoSearchGet**
> ImagerepoRepoResponse ImagereposRepoSearchGet(ctx, repo)
List images by repo

Get images by repo

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| 镜像仓库名称 | 

### Return type

[**ImagerepoRepoResponse**](imagerepo.RepoResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImagereposRepositoriesSearchPost**
> []string ImagereposRepositoriesSearchPost(ctx, request)
List Internal repos

List Internal repos

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**ImagerepoRepoConfig**](ImagerepoRepoConfig.md)| request | 

### Return type

**[]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

