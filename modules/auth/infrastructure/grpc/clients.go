package grpc

import (
	"context"
	"fmt"
	"sync"

	"nfxid/modules/auth/config"
	"nfxid/pkgs/tokenx"

	"google.golang.org/grpc"
)

// GRPCClients gRPC 客户端集合（Directory 和 Access，用于登录解析用户信息和角色）
type GRPCClients struct {
	DirectoryClient *DirectoryClient
	AccessClient    *AccessClient

	conns []*grpc.ClientConn
	mu    sync.Mutex
}

// NewGRPCClients 创建 gRPC 客户端连接
func NewGRPCClients(ctx context.Context, cfg *config.GRPCClientConfig, serverCfg *config.ServerConfig, tokenCfg *tokenx.Config) (*GRPCClients, error) {
	_ = ctx
	_ = serverCfg

	tokenProvider := createTokenProvider(tokenCfg)
	clients := &GRPCClients{conns: make([]*grpc.ClientConn, 0)}

	// 连接 Directory 服务
	if cfg.DirectoryAddr != "" {
		directoryConn, err := createConnection(cfg.DirectoryAddr, tokenProvider)
		if err != nil {
			return nil, fmt.Errorf("failed to create directory connection: %w", err)
		}
		clients.conns = append(clients.conns, directoryConn)
		clients.DirectoryClient = NewDirectoryClient(directoryConn)
	}

	// 连接 Access 服务
	if cfg.AccessAddr != "" {
		accessConn, err := createConnection(cfg.AccessAddr, tokenProvider)
		if err != nil {
			return nil, fmt.Errorf("failed to create access connection: %w", err)
		}
		clients.conns = append(clients.conns, accessConn)
		clients.AccessClient = NewAccessClient(accessConn)
	}

	return clients, nil
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
