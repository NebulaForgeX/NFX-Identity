package grpc

import (
	roleApp "nfxid/modules/access/application/roles"
	permissionApp "nfxid/modules/access/application/permissions"
	grantApp "nfxid/modules/access/application/grants"
	scopeApp "nfxid/modules/access/application/scopes"
	rolePermissionApp "nfxid/modules/access/application/role_permissions"
	scopePermissionApp "nfxid/modules/access/application/scope_permissions"
	grpcHandler "nfxid/modules/access/interfaces/grpc/handler"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"
	rolepb "nfxid/protos/gen/access/role"
	permissionpb "nfxid/protos/gen/access/permission"
	grantpb "nfxid/protos/gen/access/grant"
	scopepb "nfxid/protos/gen/access/scope"
	rolepermissionpb "nfxid/protos/gen/access/role_permission"
	scopepermissionpb "nfxid/protos/gen/access/scope_permission"

	"google.golang.org/grpc"
)

type Deps interface {
	RoleAppSvc() *roleApp.Service
	PermissionAppSvc() *permissionApp.Service
	GrantAppSvc() *grantApp.Service
	ScopeAppSvc() *scopeApp.Service
	RolePermissionAppSvc() *rolePermissionApp.Service
	ScopePermissionAppSvc() *scopePermissionApp.Service
	ServerTokenVerifier() token.Verifier
}

func NewServer(d Deps) *grpc.Server {
	// 创建 gRPC 服务器，添加认证拦截器（使用 ServerTokenVerifier 用于服务间通信）
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier())),
	}

	s := grpc.NewServer(opts...)

	// Register gRPC services
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

	return s
}
