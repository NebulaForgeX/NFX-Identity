package grpc

import (
	systemStateApp "nfxid/modules/system/application/system_state"
	resourceApp "nfxid/modules/system/application/resource"
	grpcHandler "nfxid/modules/system/interfaces/grpc/handler"
	"nfxid/pkgs/postgresqlx"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"
	healthpb "nfxid/protos/gen/common/health"
	schemapb "nfxid/protos/gen/common/schema"
	systemstatepb "nfxid/protos/gen/system/system_state"

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
		grpc.UnaryInterceptor(servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier())),
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
