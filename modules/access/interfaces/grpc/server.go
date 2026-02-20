package grpc

import (
	tenantrolesApp "nfxid/modules/access/application/tenant_roles"
	resourceApp "nfxid/modules/access/application/resource"
	grpcHandler "nfxid/modules/access/interfaces/grpc/handler"
	"nfxid/pkgs/postgresqlx"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"
	tenantrolepb "nfxid/protos/gen/access/tenant_role"
	healthpb "nfxid/protos/gen/common/health"
	schemapb "nfxid/protos/gen/common/schema"

	"google.golang.org/grpc"
)

type Deps interface {
	TenantRoleAppSvc() *tenantrolesApp.Service
	ResourceSvc() *resourceApp.Service
	ServerTokenVerifier() token.Verifier
	Postgres() *postgresqlx.Connection
}

func NewServer(d Deps) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier())),
	}
	s := grpc.NewServer(opts...)

	tenantrolepb.RegisterTenantRoleServiceServer(s, grpcHandler.NewTenantRoleHandler(d.TenantRoleAppSvc()))
	healthpb.RegisterHealthServiceServer(s, grpcHandler.NewHealthHandler(d.ResourceSvc(), "access"))
	schemapb.RegisterSchemaServiceServer(s, grpcHandler.NewSchemaHandler(d.Postgres().DB(), "access"))

	return s
}
