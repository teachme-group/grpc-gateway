package registry

import (
	"context"

	"github.com/Markuysa/pkg/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type service interface {
	Register(ctx context.Context, mux *runtime.ServeMux) error
	Name() string
}

type services map[string]service

func (s services) Register(ctx context.Context, mux *runtime.ServeMux) error {
	for _, service := range s {
		err := service.Register(ctx, mux)
		if err != nil {
			return err
		}

		log.Infof("service %s registered", service.Name())
	}

	return nil
}
