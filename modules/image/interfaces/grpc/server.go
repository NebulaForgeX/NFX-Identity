package grpc

import (
	grpcHandler "nfxid/modules/image/interfaces/grpc/handler"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"

	"google.golang.org/grpc"
)

type Deps interface {
	// TODO: Add application services when application layer is created
	ServerTokenVerifier() token.Verifier
}

func NewServer(d Deps) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier())),
	}

	s := grpc.NewServer(opts...)

	// TODO: Register protobuf services when protos are available and application layer is created
	_ = grpcHandler.NewImageHandler()

	return s
}
