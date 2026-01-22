package grpc

import (
	eventApp "nfxid/modules/audit/application/events"
	resourceApp "nfxid/modules/audit/application/resource"
	grpcHandler "nfxid/modules/audit/interfaces/grpc/handler"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"
	eventpb "nfxid/protos/gen/audit/event"
	healthpb "nfxid/protos/gen/common/health"

	"google.golang.org/grpc"
)

type Deps interface {
	EventAppSvc() *eventApp.Service
	ResourceSvc() *resourceApp.Service
	ServerTokenVerifier() token.Verifier
}

func NewServer(d Deps) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier())),
	}

	s := grpc.NewServer(opts...)

	eventpb.RegisterEventServiceServer(s, grpcHandler.NewEventHandler(d.EventAppSvc()))

	// Register health check service
	healthpb.RegisterHealthServiceServer(s, grpcHandler.NewHealthHandler(d.ResourceSvc(), "audit"))

	return s
}
