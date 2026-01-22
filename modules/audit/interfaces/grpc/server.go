package grpc

import (
	eventApp "nfxid/modules/audit/application/events"
	grpcHandler "nfxid/modules/audit/interfaces/grpc/handler"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"
	eventpb "nfxid/protos/gen/audit/event"

	"google.golang.org/grpc"
)

type Deps interface {
	EventAppSvc() *eventApp.Service
	ServerTokenVerifier() token.Verifier
}

func NewServer(d Deps) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier())),
	}

	s := grpc.NewServer(opts...)

	eventpb.RegisterEventServiceServer(s, grpcHandler.NewEventHandler(d.EventAppSvc()))

	return s
}
