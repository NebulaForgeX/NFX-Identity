package grpc

import (
	systemStateApp "nfxid/modules/system/application/system_state"
	grpcHandler "nfxid/modules/system/interfaces/grpc/handler"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"
	systemstatepb "nfxid/protos/gen/system/system_state"

	"google.golang.org/grpc"
)

type Deps interface {
	SystemStateAppSvc() *systemStateApp.Service
	ServerTokenVerifier() token.Verifier
}

func NewServer(d Deps) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier())),
	}

	s := grpc.NewServer(opts...)

	// Register gRPC services
	systemstatepb.RegisterSystemStateServiceServer(s, grpcHandler.NewSystemStateHandler(d.SystemStateAppSvc()))

	return s
}
