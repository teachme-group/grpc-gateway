package config

import (
	"github.com/Markuysa/pkg/consul"
	"github.com/Markuysa/pkg/log"
	"github.com/Markuysa/pkg/prober"
	promLoager "github.com/Markuysa/pkg/prometheus"
	"github.com/teachme-group/grpc-gateway/pkg/registry"
)

type (
	Config struct {
		Logger     log.Config
		Services   map[string]registry.ServiceCfg `yaml:"services" validate:"required,dive,required"`
		Consul     consul.Config                  `validate:"required" yaml:"consul"`
		Probes     prober.Config                  `validate:"required" yaml:"probes"`
		Prometheus promLoager.Config              `validate:"required" yaml:"prometheus"`
	}
)
