package grpc

import (
	resourceApp "nfxidentity/modules/access/application/resource"
	superadminsApp "nfxidentity/modules/access/application/super_admins"
	tenantrolesApp "nfxidentity/modules/access/application/tenant_roles"
	grpcHandler "nfxidentity/modules/access/interfaces/grpc/handler"
	"nfxidentity/pkgs/grpcx/interceptor"
	"nfxidentity/pkgs/postgresqlx"
	"nfxidentity/pkgs/security/token"
	"nfxidentity/pkgs/security/token/servertoken"
	superadminpb "nfxidentity/protos/gen/access/super_admin"
	tenantrolepb "nfxidentity/protos/gen/access/tenant_role"
	healthpb "nfxidentity/protos/gen/common/health"
	schemapb "nfxidentity/protos/gen/common/schema"

	"google.golang.org/grpc"
)

type Deps interface {
	TenantRoleAppSvc() *tenantrolesApp.Service
	SuperAdminAppSvc() *superadminsApp.Service
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

	superadminpb.RegisterSuperAdminServiceServer(s, grpcHandler.NewSuperAdminHandler(d.SuperAdminAppSvc()))
	tenantrolepb.RegisterTenantRoleServiceServer(s, grpcHandler.NewTenantRoleHandler(d.TenantRoleAppSvc()))
	healthpb.RegisterHealthServiceServer(s, grpcHandler.NewHealthHandler(d.ResourceSvc(), "access"))
	schemapb.RegisterSchemaServiceServer(s, grpcHandler.NewSchemaHandler(d.Postgres().DB(), "access"))

	return s
}
