package grpc

import (
	"nfxidentity/modules/auth/application/platform"
	"nfxidentity/modules/auth/application/resource"
	grpcHandler "nfxidentity/modules/auth/interfaces/grpc/handler"
	"nfxidentity/pkgs/grpcx/interceptor"
	"nfxidentity/pkgs/postgresqlx"
	"nfxidentity/pkgs/security/token"
	"nfxidentity/pkgs/security/token/servertoken"
	accountpb "nfxidentity/protos/gen/auth/account"
	authorityprofilepb "nfxidentity/protos/gen/auth/authority_profile"
	forgerprofilepb "nfxidentity/protos/gen/auth/forger_profile"
	healthpb "nfxidentity/protos/gen/common/health"
	schemapb "nfxidentity/protos/gen/common/schema"

	"google.golang.org/grpc"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

type Deps interface {
	PlatformSvc() *platform.Service
	ResourceSvc() *resource.Service
	ServerTokenVerifier() token.Verifier
	Postgres() *postgresqlx.Connection
}

func NewServer(d Deps) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(
			interceptor.UnaryErrorHandler(),
			servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier()),
		),
	}
	s := grpc.NewServer(opts...)
	accountpb.RegisterAccountServiceServer(s, grpcHandler.NewAccountHandler(d.PlatformSvc()))
	forgerprofilepb.RegisterForgerProfileServiceServer(s, grpcHandler.NewForgerHandler(d.PlatformSvc()))
	authorityprofilepb.RegisterAuthorityProfileServiceServer(s, grpcHandler.NewAuthorityHandler(d.PlatformSvc()))
	healthpb.RegisterHealthServiceServer(s, grpcHandler.NewHealthHandler(d.ResourceSvc(), "auth"))
	schemapb.RegisterSchemaServiceServer(s, grpcHandler.NewSchemaHandler(d.Postgres().DB(), "auth"))
	return s
}
