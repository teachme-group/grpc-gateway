package registry

type (
	ServiceCfg struct {
		Address string `yaml:"address" validate:"required"`
		TLS     *TLS   `yaml:"tls"`
	}
	TLS struct {
		CertificatePath string `yaml:"certificate_path" validate:"required"`
		KeyPath         string `yaml:"key_path" validate:"required"`
	}
)
