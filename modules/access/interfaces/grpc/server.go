package grpc

import (
	actionRequirementApp "nfxid/modules/access/application/action_requirements"
	actionApp "nfxid/modules/access/application/actions"
	grantApp "nfxid/modules/access/application/grants"
	permissionApp "nfxid/modules/access/application/permissions"
	resourceApp "nfxid/modules/access/application/resource"
	rolePermissionApp "nfxid/modules/access/application/role_permissions"
	roleApp "nfxid/modules/access/application/roles"
	scopePermissionApp "nfxid/modules/access/application/scope_permissions"
	scopeApp "nfxid/modules/access/application/scopes"
	grpcHandler "nfxid/modules/access/interfaces/grpc/handler"
	"nfxid/pkgs/postgresqlx"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"
	actionpb "nfxid/protos/gen/access/action"
	actionrequirementpb "nfxid/protos/gen/access/action_requirement"
	grantpb "nfxid/protos/gen/access/grant"
	permissionpb "nfxid/protos/gen/access/permission"
	rolepb "nfxid/protos/gen/access/role"
	rolepermissionpb "nfxid/protos/gen/access/role_permission"
	scopepb "nfxid/protos/gen/access/scope"
	scopepermissionpb "nfxid/protos/gen/access/scope_permission"
	healthpb "nfxid/protos/gen/common/health"
	schemapb "nfxid/protos/gen/common/schema"

	"google.golang.org/grpc"
)

type Deps interface {
	ActionAppSvc() *actionApp.Service
	ActionRequirementAppSvc() *actionRequirementApp.Service
	RoleAppSvc() *roleApp.Service
	PermissionAppSvc() *permissionApp.Service
	GrantAppSvc() *grantApp.Service
	ScopeAppSvc() *scopeApp.Service
	RolePermissionAppSvc() *rolePermissionApp.Service
	ScopePermissionAppSvc() *scopePermissionApp.Service
	ResourceSvc() *resourceApp.Service
	ServerTokenVerifier() token.Verifier
	Postgres() *postgresqlx.Connection
}

func NewServer(d Deps) *grpc.Server {
	// 创建 gRPC 服务器，添加认证拦截器（使用 ServerTokenVerifier 用于服务间通信）
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier())),
	}

	s := grpc.NewServer(opts...)

	// Register gRPC services
	actionpb.RegisterActionServiceServer(s, grpcHandler.NewActionHandler(d.ActionAppSvc()))
	actionrequirementpb.RegisterActionRequirementServiceServer(s, grpcHandler.NewActionRequirementHandler(d.ActionRequirementAppSvc()))
	rolepb.RegisterRoleServiceServer(s, grpcHandler.NewRoleHandler(
		d.RoleAppSvc(),
	))

	permissionpb.RegisterPermissionServiceServer(s, grpcHandler.NewPermissionHandler(
		d.PermissionAppSvc(),
	))

	grantpb.RegisterGrantServiceServer(s, grpcHandler.NewGrantHandler(
		d.GrantAppSvc(),
	))

	scopepb.RegisterScopeServiceServer(s, grpcHandler.NewScopeHandler(
		d.ScopeAppSvc(),
	))

	rolepermissionpb.RegisterRolePermissionServiceServer(s, grpcHandler.NewRolePermissionHandler(
		d.RolePermissionAppSvc(),
	))

	scopepermissionpb.RegisterScopePermissionServiceServer(s, grpcHandler.NewScopePermissionHandler(
		d.ScopePermissionAppSvc(),
	))

	// Register health check service
	healthpb.RegisterHealthServiceServer(s, grpcHandler.NewHealthHandler(d.ResourceSvc(), "access"))

	// Register schema service
	schemapb.RegisterSchemaServiceServer(s, grpcHandler.NewSchemaHandler(d.Postgres().DB(), "access"))

	return s
}
