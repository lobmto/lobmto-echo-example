// Package health provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package health

// Defines values for HealthStatus.
const (
	DOWN HealthStatus = "DOWN"
	UP   HealthStatus = "UP"
)

// HealthStatus defines model for HealthStatus.
type HealthStatus string

// GetHealthResponse defines model for GetHealthResponse.
type GetHealthResponse struct {
	Status HealthStatus `json:"status"`
}
