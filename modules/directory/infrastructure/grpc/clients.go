package grpc

import (
	"context"
	"fmt"
	"sync"

	"nfxid/connections/image"
	"nfxid/modules/directory/config"
	"nfxid/pkgs/tokenx"

	imagepb "nfxid/protos/gen/image/image"

	"google.golang.org/grpc"
)

// GRPCClients Directory 模块调用的 gRPC 客户端集合（如 Image 服务，用于 create user_avatar/user_image 前校验 image 存在）
type GRPCClients struct {
	ImageClient *image.ImageClient

	conns []*grpc.ClientConn
	mu    sync.Mutex
}

// NewGRPCClients 创建 gRPC 客户端连接；directory 的 user_avatar/user_image 依赖 Image，必须配 image_addr
func NewGRPCClients(ctx context.Context, cfg *config.GRPCClientConfig, serverCfg *config.ServerConfig, tokenCfg *tokenx.Config) (*GRPCClients, error) {
	_ = ctx
	_ = serverCfg

	if cfg.ImageAddr == "" {
		return nil, fmt.Errorf("directory requires grpc_client.image_addr to be set")
	}
	tokenProvider := createTokenProvider(tokenCfg)
	clients := &GRPCClients{conns: make([]*grpc.ClientConn, 0)}

	imageConn, err := createConnection(cfg.ImageAddr, tokenProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to create image connection: %w", err)
	}
	clients.conns = append(clients.conns, imageConn)
	clients.ImageClient = image.NewImageClient(imagepb.NewImageServiceClient(imageConn))

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
