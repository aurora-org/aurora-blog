package registry

import (
	"aurora/blog/api/pkg/lifecycle"
	"context"
	"fmt"
)

var _ lifecycle.Registrar = (*HttpRegistry)(nil)

type HttpRegistry struct {
}

func (h HttpRegistry) Register(ctx context.Context, instance *lifecycle.ServiceInstance) error {
	fmt.Println("register")
	return nil
}

func (h HttpRegistry) Deregister(ctx context.Context, instance *lifecycle.ServiceInstance) error {
	fmt.Println("deregister")
	return nil
}

func New() lifecycle.Registrar {
	return &HttpRegistry{}
}
