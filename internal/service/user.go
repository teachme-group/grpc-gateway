package service

import (
	"context"
	"errors"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/teachme-group/grpc-gateway/pkg/registry"
	"github.com/teachme-group/grpc-gateway/pkg/tls"
	userV1 "github.com/teachme-group/user/pkg/api/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type userService struct {
	cfg registry.ServiceCfg
}

func NewUserService(cfg map[string]registry.ServiceCfg) *userService {
	srv := &userService{}
	if conf, ok := cfg[srv.Name()]; !ok {
		log.Fatalf("service %s not found in config", srv.Name())
	} else {
		srv.cfg = conf
	}

	return srv
}

func (s *userService) Register(ctx context.Context, mux *runtime.ServeMux) error {
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

	return errors.Join(
		userV1.RegisterSignInServiceHandlerFromEndpoint(ctx, mux, s.cfg.Address, dialOpts),
		userV1.RegisterSignUpServiceHandlerFromEndpoint(ctx, mux, s.cfg.Address, dialOpts),
	)
}

func (s *userService) Name() string {
	return "user"
}
