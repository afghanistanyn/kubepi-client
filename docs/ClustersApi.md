# \ClustersApi

All URIs are relative to *https://localhost/kubepi/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersClusterClusterrolesClusterroleDelete**](ClustersApi.md#ClustersClusterClusterrolesClusterroleDelete) | **Delete** /clusters/{cluster}/clusterroles/{clusterrole} | Delete clusterRole by name
[**ClustersClusterClusterrolesClusterrolePut**](ClustersApi.md#ClustersClusterClusterrolesClusterrolePut) | **Put** /clusters/{cluster}/clusterroles/{clusterrole} | Update Cluster Role
[**ClustersClusterClusterrolesGet**](ClustersApi.md#ClustersClusterClusterrolesGet) | **Get** /clusters/{cluster}/clusterroles | List all clusterRoles
[**ClustersClusterClusterrolesPost**](ClustersApi.md#ClustersClusterClusterrolesPost) | **Post** /clusters/{cluster}/clusterroles | Create Cluster Role
[**ClustersClusterMembersGet**](ClustersApi.md#ClustersClusterMembersGet) | **Get** /clusters/{cluster}/members | List all ClusterMembers
[**ClustersClusterMembersMemberDelete**](ClustersApi.md#ClustersClusterMembersMemberDelete) | **Delete** /clusters/{cluster}/members/{member} | Delete clusterMember by name
[**ClustersClusterMembersMemberGet**](ClustersApi.md#ClustersClusterMembersMemberGet) | **Get** /clusters/{cluster}/members/{member} | Get Cluster Member By name
[**ClustersClusterMembersMemberPut**](ClustersApi.md#ClustersClusterMembersMemberPut) | **Put** /clusters/{cluster}/members/{member} | Update Cluster Member
[**ClustersClusterMembersPost**](ClustersApi.md#ClustersClusterMembersPost) | **Post** /clusters/{cluster}/members | Create Cluster Member
[**ClustersClusterReposGet**](ClustersApi.md#ClustersClusterReposGet) | **Get** /clusters/{cluster}/repos | Get ClusterRepo Detail
[**ClustersClusterReposPost**](ClustersApi.md#ClustersClusterReposPost) | **Post** /clusters/{cluster}/repos | Create Cluster Repo
[**ClustersClusterReposRepoDelete**](ClustersApi.md#ClustersClusterReposRepoDelete) | **Delete** /clusters/{cluster}/repos/{repo} | Delete clusterRepo by name
[**ClustersGet**](ClustersApi.md#ClustersGet) | **Get** /clusters | List all clusters
[**ClustersNameDelete**](ClustersApi.md#ClustersNameDelete) | **Delete** /clusters/{name} | Delete cluster by name
[**ClustersNameGet**](ClustersApi.md#ClustersNameGet) | **Get** /clusters/{name} | Get cluster by name
[**ClustersNamePut**](ClustersApi.md#ClustersNamePut) | **Put** /clusters/{name} | Update cluster by name
[**ClustersPost**](ClustersApi.md#ClustersPost) | **Post** /clusters | Create Cluster


# **ClustersClusterClusterrolesClusterroleDelete**
> float32 ClustersClusterClusterrolesClusterroleDelete(ctx, cluster, clusterrole)
Delete clusterRole by name

Delete clusterRole by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cluster** | **string**| 集群名称 | 
  **clusterrole** | **string**| 权限名称 | 

### Return type

**float32**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersClusterClusterrolesClusterrolePut**
> V1ClusterRole ClustersClusterClusterrolesClusterrolePut(ctx, cluster, clusterrole, request)
Update Cluster Role

Update Cluster Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cluster** | **string**| 集群名称 | 
  **clusterrole** | **string**| 权限名称 | 
  **request** | [**V1ClusterRole**](V1ClusterRole.md)| request | 

### Return type

[**V1ClusterRole**](v1.ClusterRole.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersClusterClusterrolesGet**
> []V1ClusterRole ClustersClusterClusterrolesGet(ctx, cluster)
List all clusterRoles

List all clusterRoles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cluster** | **string**| 集群名称 | 

### Return type

[**[]V1ClusterRole**](v1.ClusterRole.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersClusterClusterrolesPost**
> V1ClusterRole ClustersClusterClusterrolesPost(ctx, cluster, request)
Create Cluster Role

Create Cluster Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cluster** | **string**| 集群名称 | 
  **request** | [**V1ClusterRole**](V1ClusterRole.md)| request | 

### Return type

[**V1ClusterRole**](v1.ClusterRole.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersClusterMembersGet**
> []ClusterMember ClustersClusterMembersGet(ctx, cluster)
List all ClusterMembers

List all ClusterMembers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cluster** | **string**| 集群名称 | 

### Return type

[**[]ClusterMember**](cluster.Member.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersClusterMembersMemberDelete**
> float32 ClustersClusterMembersMemberDelete(ctx, cluster, member)
Delete clusterMember by name

Delete clusterMember by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cluster** | **string**| 集群名称 | 
  **member** | **string**| 成员名称 | 

### Return type

**float32**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersClusterMembersMemberGet**
> ClusterMember ClustersClusterMembersMemberGet(ctx, cluster, member)
Get Cluster Member By name

Get Cluster Member By name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cluster** | **string**| 集群名称 | 
  **member** | **string**| 成员名称 | 

### Return type

[**ClusterMember**](cluster.Member.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersClusterMembersMemberPut**
> ClusterMember ClustersClusterMembersMemberPut(ctx, cluster, member, request)
Update Cluster Member

Update Cluster Member

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cluster** | **string**| 集群名称 | 
  **member** | **string**| 成员名称 | 
  **request** | [**ClusterMember**](ClusterMember.md)| request | 

### Return type

[**ClusterMember**](cluster.Member.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersClusterMembersPost**
> ClusterMember ClustersClusterMembersPost(ctx, cluster, request)
Create Cluster Member

Create Cluster Member

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cluster** | **string**| 集群名称 | 
  **request** | [**ClusterMember**](ClusterMember.md)| request | 

### Return type

[**ClusterMember**](cluster.Member.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersClusterReposGet**
> []ImagerepoImageRepo ClustersClusterReposGet(ctx, cluster)
Get ClusterRepo Detail

Get ClusterRepo Detail

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

# **ClustersClusterReposPost**
> ClusterCreateRepo ClustersClusterReposPost(ctx, cluster, request)
Create Cluster Repo

Create Cluster Repo

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cluster** | **string**| 集群名称 | 
  **request** | [**ClusterCreateRepo**](ClusterCreateRepo.md)| request | 

### Return type

[**ClusterCreateRepo**](cluster.CreateRepo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersClusterReposRepoDelete**
> float32 ClustersClusterReposRepoDelete(ctx, cluster, repo)
Delete clusterRepo by name

Delete clusterRepo by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cluster** | **string**| 集群名称 | 
  **repo** | **string**| 镜像仓库名称 | 

### Return type

**float32**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersGet**
> []GithubComKubeOperatorKubepiInternalModelV1ClusterCluster ClustersGet(ctx, )
List all clusters

List all clusters

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]GithubComKubeOperatorKubepiInternalModelV1ClusterCluster**](github.com_KubeOperator_kubepi_internal_model_v1_cluster.Cluster.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersNameDelete**
> GithubComKubeOperatorKubepiInternalModelV1ClusterCluster ClustersNameDelete(ctx, name)
Delete cluster by name

Delete cluster by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| 集群名称 | 

### Return type

[**GithubComKubeOperatorKubepiInternalModelV1ClusterCluster**](github.com_KubeOperator_kubepi_internal_model_v1_cluster.Cluster.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersNameGet**
> GithubComKubeOperatorKubepiInternalModelV1ClusterCluster ClustersNameGet(ctx, name)
Get cluster by name

Get cluster by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| 集群名称 | 

### Return type

[**GithubComKubeOperatorKubepiInternalModelV1ClusterCluster**](github.com_KubeOperator_kubepi_internal_model_v1_cluster.Cluster.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersNamePut**
> ClusterUpdateCluster ClustersNamePut(ctx, request, name)
Update cluster by name

Update cluster by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**ClusterUpdateCluster**](ClusterUpdateCluster.md)| request | 
  **name** | **string**| 集群名称 | 

### Return type

[**ClusterUpdateCluster**](cluster.UpdateCluster.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersPost**
> InternalModelV1ClusterCluster ClustersPost(ctx, request)
Create Cluster

Create Cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**InternalModelV1ClusterCluster**](InternalModelV1ClusterCluster.md)| request | 

### Return type

[**InternalModelV1ClusterCluster**](internal_model_v1_cluster.Cluster.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

