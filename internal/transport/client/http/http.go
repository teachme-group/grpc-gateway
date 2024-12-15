package http

import (
	"context"
	"net/http"

	"github.com/Markuysa/pkg/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type transport struct {
	mux      *runtime.ServeMux
	registry registry
}

func New(registry registry) *transport {
	return &transport{
		mux:      runtime.NewServeMux(),
		registry: registry,
	}
}

func (t *transport) RegisterServicesAndRun(ctx context.Context) error {
	log.Info("registering services")

	if err := t.registry.Register(ctx, t.mux); err != nil {
		return err
	}

	log.Info("http server started on :8000")

	return http.ListenAndServe(":8000", t.mux)
}
