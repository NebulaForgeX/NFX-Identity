package grpc

import (
	"context"

	"nfxidentity/modules/auth/application/account"
	"nfxidentity/modules/auth/application/resource"
	grpcHandler "nfxidentity/modules/auth/interface/grpc/handler"
	"nfxidentity/pkgs/grpcx/interceptor"
	"nfxidentity/pkgs/postgresqlx"
	"nfxidentity/pkgs/security/token"
	"nfxidentity/pkgs/security/token/servertoken"
	accountpb "nfxidentity/protos/gen/auth/account"
	authorityprofilepb "nfxidentity/protos/gen/auth/authority_profile"
	forgerprofilepb "nfxidentity/protos/gen/auth/forger_profile"
	healthpb "nfxidentity/protos/gen/common/health"
	schemapb "nfxidentity/protos/gen/common/schema"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type Deps interface {
	AccountService() *account.Service
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
	accountpb.RegisterAccountServiceServer(s, grpcHandler.NewAccountHandler(d.AccountService()))
	forgerprofilepb.RegisterForgerProfileServiceServer(s, grpcHandler.NewForgerHandler(d.AccountService()))
	authorityprofilepb.RegisterAuthorityProfileServiceServer(s, grpcHandler.NewAuthorityHandler(d.AccountService()))
	healthpb.RegisterHealthServiceServer(s, grpcHandler.NewHealthHandler(d.ResourceSvc(), "auth"))
	schemapb.RegisterSchemaServiceServer(s, grpcHandler.NewSchemaHandler(func(ctx context.Context) (int32, error) {
		n, err := postgresqlx.ClearSchema(ctx, d.Postgres().DB(), "auth", nil)
		return int32(n), err
	}))
	return s
}
