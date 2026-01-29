package grpc

import (
	"fmt"

	"nfxid/pkgs/grpcx"
	"nfxid/pkgs/security/token/servertoken"
	"nfxid/pkgs/tokenx"

	"google.golang.org/grpc"
)

func createConnection(addr string, tokenProvider servertoken.TokenProvider) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, grpcx.DefaultClientOptions(tokenProvider)...)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC client: %w", err)
	}
	return conn, nil
}

func createTokenProvider(tokenCfg *tokenx.Config) servertoken.TokenProvider {
	return servertoken.NewProvider(
		&servertoken.HMACSigner{Key: []byte(tokenCfg.SecretKey)},
		tokenCfg.Issuer,
		"directory-service",
	)
}
