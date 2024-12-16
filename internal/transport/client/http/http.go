package http

import (
	"context"
	"net/http"

	"github.com/Markuysa/pkg/log"
	"github.com/Markuysa/pkg/prober"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type transport struct {
	cfg      prober.Config
	mux      *runtime.ServeMux
	registry registry
}

func New(cfg prober.Config, registry registry) *transport {
	return &transport{
		cfg:      cfg,
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
