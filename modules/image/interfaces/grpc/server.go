package grpc

import (
	resourceApp "nfxid/modules/image/application/resource"
	grpcHandler "nfxid/modules/image/interfaces/grpc/handler"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"
	healthpb "nfxid/protos/gen/common/health"

	"google.golang.org/grpc"
)

type Deps interface {
	// TODO: Add application services when application layer is created
	ResourceSvc() *resourceApp.Service
	ServerTokenVerifier() token.Verifier
}

func NewServer(d Deps) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier())),
	}

	s := grpc.NewServer(opts...)

	// TODO: Register protobuf services when protos are available and application layer is created
	_ = grpcHandler.NewImageHandler()

	// Register health check service
	healthpb.RegisterHealthServiceServer(s, grpcHandler.NewHealthHandler(d.ResourceSvc(), "image"))

	return s
}
