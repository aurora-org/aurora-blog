package registry

import "context"

// Registrar is service registrar.
type Registrar interface {
	Register(context.Context, *ServiceInstance) error
	Deregister(context.Context, *ServiceInstance) error
}

// ServiceInstance ...
type ServiceInstance struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Version   string            `json:"version"`
	Metadata  map[string]string `json:"metadata"`
	Endpoints []string          `json:"endpoints"`
}
