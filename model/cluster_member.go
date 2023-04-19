/*
 * KubePi Restful API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Contact: support@fit2cloud.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ClusterMember struct {
	BindingName string `json:"bindingName,omitempty"`
	ClusterRoles []string `json:"clusterRoles,omitempty"`
	CreateAt string `json:"createAt,omitempty"`
	Name string `json:"name,omitempty"`
	NamespaceRoles []ClusterNamespaceRoles `json:"namespaceRoles,omitempty"`
}
