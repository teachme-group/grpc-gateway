package app

import (
	"context"

	"github.com/Markuysa/pkg/log"
	"github.com/teachme-group/grpc-gateway/internal/config"
	"github.com/teachme-group/grpc-gateway/internal/service"
	"github.com/teachme-group/grpc-gateway/internal/transport/client/http"
	"github.com/teachme-group/grpc-gateway/pkg/registry"
)

func Run(ctx context.Context, cfg *config.Config) error {
	err := log.InitLogger(cfg.Logger)
	if err != nil {
		return err
	}

	registry := registry.NewRegistry()

	registry.MustRegisterMany(
		service.NewSessionService(cfg.Services),
		service.NewUserService(cfg.Services),
	)

	transport := http.New(registry)

	return transport.RegisterServicesAndRun(ctx)
}
