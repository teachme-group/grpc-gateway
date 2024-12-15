package tls

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func GrpcTLS(caPath string) (grpc.DialOption, error) {
	caCert, err := os.ReadFile(caPath)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return nil, err
	}

	creds := credentials.NewTLS(&tls.Config{
		RootCAs: certPool,
	})

	return grpc.WithTransportCredentials(creds), nil
}
