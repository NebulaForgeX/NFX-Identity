package grpc

import (
	"context"
	"fmt"
	"sync"

	"nfxid/modules/system/config"
	"nfxid/pkgs/tokenx"

	"google.golang.org/grpc"
)

// GRPCClients gRPC 客户端集合
// 用于系统初始化时调用其他服务
type GRPCClients struct {
	DirectoryClient *DirectoryClient
	AccessClient    *AccessClient
	AuthClient      *AuthClient

	healthChecker *HealthChecker // 健康检查客户端管理器
	conns         []*grpc.ClientConn
	mu            sync.Mutex
}

// NewClients 创建 gRPC 客户端连接
func NewGRPCClients(ctx context.Context, cfg *config.GRPCClientConfig, serverCfg *config.ServerConfig, tokenCfg *tokenx.Config) (*GRPCClients, error) {
	// 创建 server token provider
	tokenProvider := createTokenProvider(tokenCfg)

	grpcClients := &GRPCClients{
		conns: make([]*grpc.ClientConn, 0),
	}

	// 连接 Directory 服务
	directoryConn, err := createConnection(cfg.DirectoryAddr, tokenProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to create directory connection: %w", err)
	}
	grpcClients.conns = append(grpcClients.conns, directoryConn)
	grpcClients.DirectoryClient = NewDirectoryClient(directoryConn)

	// 连接 Access 服务
	accessConn, err := createConnection(cfg.AccessAddr, tokenProvider)
	if err != nil {
		grpcClients.Close()
		return nil, fmt.Errorf("failed to create access connection: %w", err)
	}
	grpcClients.conns = append(grpcClients.conns, accessConn)
	grpcClients.AccessClient = NewAccessClient(accessConn)

	// 连接 Auth 服务
	authConn, err := createConnection(cfg.AuthAddr, tokenProvider)
	if err != nil {
		grpcClients.Close()
		return nil, fmt.Errorf("failed to create auth connection: %w", err)
	}
	grpcClients.conns = append(grpcClients.conns, authConn)
	grpcClients.AuthClient = NewAuthClient(authConn)

	// 初始化健康检查客户端
	if err := initHealthClients(grpcClients, cfg, serverCfg, tokenProvider); err != nil {
		grpcClients.Close()
		return nil, fmt.Errorf("failed to initialize health clients: %w", err)
	}

	return grpcClients, nil
}

// Close 关闭所有 gRPC 连接
func (c *GRPCClients) Close() error {
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
