package grpc

import (
	eventApp "nfxid/modules/audit/application/events"
	resourceApp "nfxid/modules/audit/application/resource"
	grpcHandler "nfxid/modules/audit/interfaces/grpc/handler"
	"nfxid/pkgs/postgresqlx"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"
	eventpb "nfxid/protos/gen/audit/event"
	healthpb "nfxid/protos/gen/common/health"
	schemapb "nfxid/protos/gen/common/schema"

	"google.golang.org/grpc"
)

type Deps interface {
	EventAppSvc() *eventApp.Service
	ResourceSvc() *resourceApp.Service
	ServerTokenVerifier() token.Verifier
	Postgres() *postgresqlx.Connection
}

func NewServer(d Deps) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier())),
	}

	s := grpc.NewServer(opts...)

	eventpb.RegisterEventServiceServer(s, grpcHandler.NewEventHandler(d.EventAppSvc()))

	// Register health check service
	healthpb.RegisterHealthServiceServer(s, grpcHandler.NewHealthHandler(d.ResourceSvc(), "audit"))
	
	// Register schema service
	schemapb.RegisterSchemaServiceServer(s, grpcHandler.NewSchemaHandler(d.Postgres().DB(), "audit"))

	return s
}
