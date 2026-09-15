package grpc

import (
	"context"
	"fmt"
	"sync"

	"nfxidentity/modules/system/config"
	"nfxidentity/pkgs/tokenx"

	"google.golang.org/grpc"
)

type GRPCClients struct {
	healthChecker *HealthChecker
	schemaChecker *SchemaChecker
	conns         []*grpc.ClientConn
	mu            sync.Mutex
}

func NewGRPCClients(ctx context.Context, cfg *config.GRPCClientConfig, serverCfg *config.ServerConfig, tokenCfg *tokenx.Config) (*GRPCClients, error) {
	tokenProvider := createTokenProvider(tokenCfg)
	grpcClients := &GRPCClients{conns: make([]*grpc.ClientConn, 0)}
	if err := initHealthClients(grpcClients, cfg, serverCfg, tokenProvider); err != nil {
		grpcClients.Close()
		return nil, fmt.Errorf("failed to initialize health clients: %w", err)
	}
	if err := initSchemaClients(grpcClients, cfg, serverCfg, tokenProvider); err != nil {
		grpcClients.Close()
		return nil, fmt.Errorf("failed to initialize schema clients: %w", err)
	}
	return grpcClients, nil
}

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
