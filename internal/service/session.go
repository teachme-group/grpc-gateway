package service

import (
	"context"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/teachme-group/grpc-gateway/pkg/registry"
	"github.com/teachme-group/grpc-gateway/pkg/tls"
	sessionV1 "github.com/teachme-group/session/pkg/api/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type sessionService struct {
	cfg registry.ServiceCfg
}

func NewSessionService(cfg map[string]registry.ServiceCfg) *sessionService {
	srv := &sessionService{}
	if conf, ok := cfg[srv.Name()]; !ok {
		log.Fatalf("service %s not found in config", srv.Name())
	} else {
		srv.cfg = conf
	}

	return srv
}

func (s *sessionService) Register(ctx context.Context, mux *runtime.ServeMux) error {
	var dialOpts []grpc.DialOption
	if s.cfg.TLS != nil {
		tlsOpts, err := tls.GrpcTLS(s.cfg.TLS.CertificatePath)
		if err != nil {
			return err
		}

		dialOpts = append(dialOpts, tlsOpts)
	} else {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return sessionV1.RegisterSessionServiceHandlerFromEndpoint(ctx, mux, s.cfg.Address, dialOpts)
}

func (s *sessionService) Name() string {
	return "session"
}
