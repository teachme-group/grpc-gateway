package http

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type (
	registry interface {
		Register(ctx context.Context, mux *runtime.ServeMux) error
	}
)
