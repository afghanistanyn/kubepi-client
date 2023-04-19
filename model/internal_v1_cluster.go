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

type InternalV1Cluster struct {
	ApiVersion string `json:"apiVersion,omitempty"`
	BuiltIn bool `json:"builtIn,omitempty"`
	CaCertificate *InternalModelV1ClusterCertificate `json:"caCertificate,omitempty"`
	CreateAt string `json:"createAt,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	Description string `json:"description,omitempty"`
	Kind string `json:"kind,omitempty"`
	Labels []string `json:"labels,omitempty"`
	Name string `json:"name,omitempty"`
	PrivateKey []int32 `json:"privateKey,omitempty"`
	Spec *InternalModelV1ClusterSpec `json:"spec,omitempty"`
	Status *InternalModelV1ClusterStatus `json:"status,omitempty"`
	UpdateAt string `json:"updateAt,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}
