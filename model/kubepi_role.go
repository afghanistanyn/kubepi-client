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

type KubepiRole struct {
	ApiVersion string `json:"apiVersion,omitempty"`
	BuiltIn bool `json:"builtIn,omitempty"`
	CreateAt string `json:"createAt,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	Description string `json:"description,omitempty"`
	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`
	Rules []RolePolicyRule `json:"rules,omitempty"`
	UpdateAt string `json:"updateAt,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}