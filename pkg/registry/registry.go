package registry

import (
	"context"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type registry struct {
	services services
}

func NewRegistry() *registry {
	return &registry{
		services: services{},
	}
}

func (r *registry) MustRegister(service service) {
	if r.services[service.Name()] != nil {
		log.Fatalf("service %s already registered", service.Name())
	}

	r.services[service.Name()] = service
}

func (r *registry) MustRegisterMany(services ...service) {
	for _, service := range services {
		r.MustRegister(service)
	}
}

func (r *registry) Register(ctx context.Context, mux *runtime.ServeMux) error {
	return r.services.Register(ctx, mux)
}
