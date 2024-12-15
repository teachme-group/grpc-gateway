package config

import (
	"github.com/Markuysa/pkg/log"
	"github.com/teachme-group/grpc-gateway/pkg/registry"
)

type (
	Config struct {
		Logger   log.Config
		Services map[string]registry.ServiceCfg `yaml:"services" validate:"required,dive,required"`
	}
)
