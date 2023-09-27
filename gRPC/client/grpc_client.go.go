package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ClientConn *grpc.ClientConn

func InitGRPCClient() error {
	creds := insecure.NewCredentials()
	// Using insecure credentials allows for insecure (non-TLS) connections.
	// It's generally used for development/testing purposes and not recommended for production.


	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(creds))
	if err != nil {
		return err
	}

	ClientConn = conn

	return nil
}
