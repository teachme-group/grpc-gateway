package main

import (
	"log"
	"os"

	cfgLoader "github.com/Markuysa/pkg/config"
	"github.com/teachme-group/web-bff/internal/app"
	"github.com/teachme-group/web-bff/internal/config"
)

const (
	cfgPathKey = "CONFIG_PATH"
)

func main() {
	onBuild()

	cfgPath := os.Getenv(cfgPathKey)
	cfg := &config.Config{}

	err := cfgLoader.LoadFromYAML(cfg, cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	if err = app.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
