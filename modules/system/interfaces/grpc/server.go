package grpc

import (
	resourceApp "nfxidentity/modules/system/application/resource"
	systemStateApp "nfxidentity/modules/system/application/system_state"
	grpcHandler "nfxidentity/modules/system/interfaces/grpc/handler"
	"nfxidentity/pkgs/grpcx/interceptor"
	"nfxidentity/pkgs/postgresqlx"
	"nfxidentity/pkgs/security/token"
	"nfxidentity/pkgs/security/token/servertoken"
	healthpb "nfxidentity/protos/gen/common/health"
	schemapb "nfxidentity/protos/gen/common/schema"
	systemstatepb "nfxidentity/protos/gen/system/system_state"

	"google.golang.org/grpc"
)

type Deps interface {
	SystemStateAppSvc() *systemStateApp.Service
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

	// Register gRPC services
	systemstatepb.RegisterSystemStateServiceServer(s, grpcHandler.NewSystemStateHandler(d.SystemStateAppSvc()))

	// Register health check service
	healthpb.RegisterHealthServiceServer(s, grpcHandler.NewHealthHandler(d.ResourceSvc(), "system"))

	// Register schema service (special handler for system to ensure system_state has only one record)
	schemapb.RegisterSchemaServiceServer(s, grpcHandler.NewSystemSchemaHandler(d.Postgres().DB(), "system"))

	return s
}
