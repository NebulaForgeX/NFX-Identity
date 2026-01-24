package grpc

import (
	"fmt"

	"nfxid/connections/access"
	"nfxid/connections/directory"
	"nfxid/pkgs/grpcx"
	"nfxid/pkgs/security/token/servertoken"
	"nfxid/pkgs/tokenx"

	grantpb "nfxid/protos/gen/access/grant"
	permissionpb "nfxid/protos/gen/access/permission"
	rolepb "nfxid/protos/gen/access/role"
	rolepermissionpb "nfxid/protos/gen/access/role_permission"
	useremailpb "nfxid/protos/gen/directory/user_email"
	userpb "nfxid/protos/gen/directory/user"
	userphonepb "nfxid/protos/gen/directory/user_phone"
	userpreferencepb "nfxid/protos/gen/directory/user_preference"
	userprofilepb "nfxid/protos/gen/directory/user_profile"

	"google.golang.org/grpc"
)

// DirectoryClient Directory 服务客户端（只包含需要的服务）
type DirectoryClient struct {
	User          *directory.UserClient
	UserEmail     *directory.UserEmailClient
	UserPhone     *directory.UserPhoneClient
	UserProfile   *directory.UserProfileClient
	UserPreference *directory.UserPreferenceClient
}

// NewDirectoryClient 创建 Directory 客户端
func NewDirectoryClient(conn *grpc.ClientConn) *DirectoryClient {
	return &DirectoryClient{
		User:          directory.NewUserClient(userpb.NewUserServiceClient(conn)),
		UserEmail:     directory.NewUserEmailClient(useremailpb.NewUserEmailServiceClient(conn)),
		UserPhone:     directory.NewUserPhoneClient(userphonepb.NewUserPhoneServiceClient(conn)),
		UserProfile:   directory.NewUserProfileClient(userprofilepb.NewUserProfileServiceClient(conn)),
		UserPreference: directory.NewUserPreferenceClient(userpreferencepb.NewUserPreferenceServiceClient(conn)),
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

// createTokenProvider 创建 server token provider
func createTokenProvider(tokenCfg *tokenx.Config) servertoken.TokenProvider {
	return servertoken.NewProvider(
		&servertoken.HMACSigner{Key: []byte(tokenCfg.SecretKey)},
		tokenCfg.Issuer,
		"auth-service", // service ID
	)
}
