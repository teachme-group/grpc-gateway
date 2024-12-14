package config

import "github.com/Markuysa/pkg/postgres"

type Config struct {
	Postgres postgres.PgxPoolCfg
}
