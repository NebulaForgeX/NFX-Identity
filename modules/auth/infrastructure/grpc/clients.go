package grpc

import (
	"context"
	"fmt"
	"sync"

	"nfxid/modules/auth/config"
	"nfxid/pkgs/tokenx"

	"google.golang.org/grpc"
)

// GRPCClients gRPC 客户端集合（仅 Directory，用于登录解析邮箱/手机 → user_id）
type GRPCClients struct {
	DirectoryClient *DirectoryClient

	conns []*grpc.ClientConn
	mu    sync.Mutex
}

// NewGRPCClients 创建 gRPC 客户端连接
func NewGRPCClients(ctx context.Context, cfg *config.GRPCClientConfig, serverCfg *config.ServerConfig, tokenCfg *tokenx.Config) (*GRPCClients, error) {
	_ = ctx
	_ = serverCfg

	tokenProvider := createTokenProvider(tokenCfg)
	clients := &GRPCClients{conns: make([]*grpc.ClientConn, 0)}

	if cfg.DirectoryAddr == "" {
		return clients, nil
	}

	directoryConn, err := createConnection(cfg.DirectoryAddr, tokenProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to create directory connection: %w", err)
	}
	clients.conns = append(clients.conns, directoryConn)
	clients.DirectoryClient = NewDirectoryClient(directoryConn)

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
