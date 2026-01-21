package grpc

import (
	"context"
	"fmt"
	"sync"

	"nfxid/modules/system/config"
	"nfxid/pkgs/grpcx"
	"nfxid/pkgs/security/token/servertoken"
	"nfxid/pkgs/tokenx"

	userpb "nfxid/protos/gen/directory/user"
	rolepb "nfxid/protos/gen/access/role"
	permissionpb "nfxid/protos/gen/access/permission"
	grantpb "nfxid/protos/gen/access/grant"
	rolepermissionpb "nfxid/protos/gen/access/role_permission"
	usercredentialpb "nfxid/protos/gen/auth/user_credential"

	"google.golang.org/grpc"
)

// Clients gRPC 客户端集合
// 用于系统初始化时调用其他服务
type Clients struct {
	DirectoryClient        userpb.UserServiceClient
	RoleClient            rolepb.RoleServiceClient
	PermissionClient      permissionpb.PermissionServiceClient
	GrantClient           grantpb.GrantServiceClient
	RolePermissionClient  rolepermissionpb.RolePermissionServiceClient
	UserCredentialClient  usercredentialpb.UserCredentialServiceClient

	conns []*grpc.ClientConn
	mu    sync.Mutex
}

// NewClients 创建 gRPC 客户端连接
func NewClients(ctx context.Context, cfg *config.GRPCClientConfig, tokenCfg *tokenx.Config) (*Clients, error) {
	// 创建 server token provider
	tokenProvider := servertoken.NewProvider(
		&servertoken.HMACSigner{Key: []byte(tokenCfg.SecretKey)},
		tokenCfg.Issuer,
		"system-service", // service ID
	)

	clients := &Clients{
		conns: make([]*grpc.ClientConn, 0),
	}

	// 连接 Directory 服务
	directoryConn, err := grpc.NewClient(
		cfg.DirectoryAddr,
		grpcx.DefaultClientOptions(tokenProvider)...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create directory client: %w", err)
	}
	clients.conns = append(clients.conns, directoryConn)
	clients.DirectoryClient = userpb.NewUserServiceClient(directoryConn)

	// 连接 Access 服务
	accessConn, err := grpc.NewClient(
		cfg.AccessAddr,
		grpcx.DefaultClientOptions(tokenProvider)...,
	)
	if err != nil {
		clients.Close()
		return nil, fmt.Errorf("failed to create access client: %w", err)
	}
	clients.conns = append(clients.conns, accessConn)
	clients.RoleClient = rolepb.NewRoleServiceClient(accessConn)
	clients.PermissionClient = permissionpb.NewPermissionServiceClient(accessConn)
	clients.GrantClient = grantpb.NewGrantServiceClient(accessConn)
	clients.RolePermissionClient = rolepermissionpb.NewRolePermissionServiceClient(accessConn)

	// 连接 Auth 服务
	authConn, err := grpc.NewClient(
		cfg.AuthAddr,
		grpcx.DefaultClientOptions(tokenProvider)...,
	)
	if err != nil {
		clients.Close()
		return nil, fmt.Errorf("failed to create auth client: %w", err)
	}
	clients.conns = append(clients.conns, authConn)
	clients.UserCredentialClient = usercredentialpb.NewUserCredentialServiceClient(authConn)

	return clients, nil
}

// Close 关闭所有 gRPC 连接
func (c *Clients) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var errs []error
	for _, conn := range c.conns {
		if err := conn.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors closing gRPC connections: %v", errs)
	}

	return nil
}
