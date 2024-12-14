package app

import (
	"log"
	"net"

	"github.com/teachme-group/web-bff/internal/config"
	"google.golang.org/grpc"
)

func Run(cfg *config.Config) error {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		return err
	}

	go func() {
		if err := grpc.NewServer().Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}
