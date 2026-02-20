package grpc

import (
	resourceApp "nfxid/modules/access/application/resource"
	superadminsApp "nfxid/modules/access/application/super_admins"
	tenantrolesApp "nfxid/modules/access/application/tenant_roles"
	grpcHandler "nfxid/modules/access/interfaces/grpc/handler"
	"nfxid/pkgs/postgresqlx"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"
	healthpb "nfxid/protos/gen/common/health"
	schemapb "nfxid/protos/gen/common/schema"
	superadminpb "nfxid/protos/gen/access/super_admin"
	tenantrolepb "nfxid/protos/gen/access/tenant_role"

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
		grpc.UnaryInterceptor(servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier())),
	}
	s := grpc.NewServer(opts...)

	superadminpb.RegisterSuperAdminServiceServer(s, grpcHandler.NewSuperAdminHandler(d.SuperAdminAppSvc()))
	tenantrolepb.RegisterTenantRoleServiceServer(s, grpcHandler.NewTenantRoleHandler(d.TenantRoleAppSvc()))
	healthpb.RegisterHealthServiceServer(s, grpcHandler.NewHealthHandler(d.ResourceSvc(), "access"))
	schemapb.RegisterSchemaServiceServer(s, grpcHandler.NewSchemaHandler(d.Postgres().DB(), "access"))

	return s
}
