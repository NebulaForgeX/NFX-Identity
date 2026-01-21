package grpc

import (
	"fmt"

	"nfxid/connections/access"
	"nfxid/connections/auth"
	"nfxid/connections/directory"
	"nfxid/pkgs/grpcx"
	"nfxid/pkgs/security/token/servertoken"
	"nfxid/pkgs/tokenx"

	grantpb "nfxid/protos/gen/access/grant"
	permissionpb "nfxid/protos/gen/access/permission"
	rolepb "nfxid/protos/gen/access/role"
	rolepermissionpb "nfxid/protos/gen/access/role_permission"
	usercredentialpb "nfxid/protos/gen/auth/user_credential"
	userpb "nfxid/protos/gen/directory/user"

	"google.golang.org/grpc"
)

// DirectoryClient Directory 服务客户端（只包含需要的服务）
type DirectoryClient struct {
	User *directory.UserClient
}

// NewDirectoryClient 创建 Directory 客户端
func NewDirectoryClient(conn *grpc.ClientConn) *DirectoryClient {
	return &DirectoryClient{
		User: directory.NewUserClient(userpb.NewUserServiceClient(conn)),
	}
}

// AccessClient Access 服务客户端（只包含需要的服务）
type AccessClient struct {
	Role           *access.RoleClient
	Permission     *access.PermissionClient
	RolePermission *access.RolePermissionClient
	Grant          *access.GrantClient
}

// NewAccessClient 创建 Access 客户端
func NewAccessClient(conn *grpc.ClientConn) *AccessClient {
	return &AccessClient{
		Role:           access.NewRoleClient(rolepb.NewRoleServiceClient(conn)),
		Permission:     access.NewPermissionClient(permissionpb.NewPermissionServiceClient(conn)),
		RolePermission: access.NewRolePermissionClient(rolepermissionpb.NewRolePermissionServiceClient(conn)),
		Grant:          access.NewGrantClient(grantpb.NewGrantServiceClient(conn)),
	}
}

// AuthClient Auth 服务客户端（只包含需要的服务）
type AuthClient struct {
	UserCredential *auth.UserCredentialClient
}

// NewAuthClient 创建 Auth 客户端
func NewAuthClient(conn *grpc.ClientConn) *AuthClient {
	return &AuthClient{
		UserCredential: auth.NewUserCredentialClient(usercredentialpb.NewUserCredentialServiceClient(conn)),
	}
}

// createConnection 创建 gRPC 连接
func createConnection(addr string, tokenProvider servertoken.TokenProvider) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		addr,
		grpcx.DefaultClientOptions(tokenProvider)...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC client: %w", err)
	}
	return conn, nil
}

// createTokenProvider 创建 server token provider
func createTokenProvider(tokenCfg *tokenx.Config) servertoken.TokenProvider {
	return servertoken.NewProvider(
		&servertoken.HMACSigner{Key: []byte(tokenCfg.SecretKey)},
		tokenCfg.Issuer,
		"system-service", // service ID
	)
}