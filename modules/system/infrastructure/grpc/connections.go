package grpc

import (
	"fmt"

	"nfxid/connections/access"
	"nfxid/connections/auth"
	"nfxid/connections/directory"
	"nfxid/connections/image"
	"nfxid/pkgs/grpcx"
	"nfxid/pkgs/security/token/servertoken"
	"nfxid/pkgs/tokenx"

	applicationroleassignmentpb "nfxid/protos/gen/access/application_role_assignment"
	applicationrolepb "nfxid/protos/gen/access/application_role"
	superadminpb "nfxid/protos/gen/access/super_admin"
	tenantroleassignmentpb "nfxid/protos/gen/access/tenant_role_assignment"
	tenantrolepb "nfxid/protos/gen/access/tenant_role"
	usercredentialpb "nfxid/protos/gen/auth/user_credential"
	userpb "nfxid/protos/gen/directory/user"
	useremailpb "nfxid/protos/gen/directory/user_email"
	userphonepb "nfxid/protos/gen/directory/user_phone"
	userpreferencepb "nfxid/protos/gen/directory/user_preference"
	userprofilepb "nfxid/protos/gen/directory/user_profile"
	imagepb "nfxid/protos/gen/image/image"

	"google.golang.org/grpc"
)

// DirectoryClient Directory 服务客户端（只包含需要的服务）
type DirectoryClient struct {
	User           *directory.UserClient
	UserEmail      *directory.UserEmailClient
	UserPhone      *directory.UserPhoneClient
	UserProfile    *directory.UserProfileClient
	UserPreference *directory.UserPreferenceClient
}

// NewDirectoryClient 创建 Directory 客户端
func NewDirectoryClient(conn *grpc.ClientConn) *DirectoryClient {
	return &DirectoryClient{
		User:           directory.NewUserClient(userpb.NewUserServiceClient(conn)),
		UserEmail:      directory.NewUserEmailClient(useremailpb.NewUserEmailServiceClient(conn)),
		UserPhone:      directory.NewUserPhoneClient(userphonepb.NewUserPhoneServiceClient(conn)),
		UserProfile:    directory.NewUserProfileClient(userprofilepb.NewUserProfileServiceClient(conn)),
		UserPreference: directory.NewUserPreferenceClient(userpreferencepb.NewUserPreferenceServiceClient(conn)),
	}
}

// AccessClient Access 服务客户端（只包含需要的服务）
type AccessClient struct {
	Client *access.Client
}

// NewAccessClient 创建 Access 客户端
func NewAccessClient(conn *grpc.ClientConn) *AccessClient {
	return &AccessClient{
		Client: access.NewClient(
			superadminpb.NewSuperAdminServiceClient(conn),
			tenantrolepb.NewTenantRoleServiceClient(conn),
			tenantroleassignmentpb.NewTenantRoleAssignmentServiceClient(conn),
			applicationrolepb.NewApplicationRoleServiceClient(conn),
			applicationroleassignmentpb.NewApplicationRoleAssignmentServiceClient(conn),
		),
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

// ImageClient Image 服务客户端（初始化时清空存储等）
type ImageClient struct {
	Image *image.ImageClient
}

// NewImageClient 创建 Image 客户端
func NewImageClient(conn *grpc.ClientConn) *ImageClient {
	return &ImageClient{
		Image: image.NewImageClient(imagepb.NewImageServiceClient(conn)),
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
