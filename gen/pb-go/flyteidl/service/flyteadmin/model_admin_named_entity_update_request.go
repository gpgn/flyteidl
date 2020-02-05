/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Request to set the referenced launch plan state to the configured value.
type AdminNamedEntityUpdateRequest struct {
	ResourceType *CoreResourceType `json:"resource_type,omitempty"`
	Id *AdminNamedEntityIdentifier `json:"id,omitempty"`
	Metadata *AdminNamedEntityMetadata `json:"metadata,omitempty"`
}