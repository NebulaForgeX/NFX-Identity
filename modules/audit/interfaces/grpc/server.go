package grpc

import (
	eventApp "nfxidentity/modules/audit/application/events"
	resourceApp "nfxidentity/modules/audit/application/resource"
	grpcHandler "nfxidentity/modules/audit/interfaces/grpc/handler"
	"nfxidentity/pkgs/grpcx/interceptor"
	"nfxidentity/pkgs/postgresqlx"
	"nfxidentity/pkgs/security/token"
	"nfxidentity/pkgs/security/token/servertoken"
	eventpb "nfxidentity/protos/gen/audit/event"
	healthpb "nfxidentity/protos/gen/common/health"
	schemapb "nfxidentity/protos/gen/common/schema"

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
		grpc.ChainUnaryInterceptor(
			interceptor.UnaryErrorHandler(),
			servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier()),
		),
	}

	s := grpc.NewServer(opts...)

	eventpb.RegisterEventServiceServer(s, grpcHandler.NewEventHandler(d.EventAppSvc()))

	// Register health check service
	healthpb.RegisterHealthServiceServer(s, grpcHandler.NewHealthHandler(d.ResourceSvc(), "audit"))

	// Register schema service
	schemapb.RegisterSchemaServiceServer(s, grpcHandler.NewSchemaHandler(d.Postgres().DB(), "audit"))

	return s
}
