package grpc

import (
	"context"
	"fmt"
	"sync"

	authconn "nfxidentity/connections/auth"
	"nfxidentity/modules/system/config"
	"nfxidentity/pkgs/tokenx"
	accountpb "nfxidentity/protos/gen/auth/account"

	"google.golang.org/grpc"
)

type GRPCClients struct {
	healthChecker *HealthChecker
	schemaChecker *SchemaChecker
	auth          *authconn.Client
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
	if err := initAuthClient(grpcClients, cfg, tokenCfg); err != nil {
		grpcClients.Close()
		return nil, fmt.Errorf("failed to initialize auth client: %w", err)
	}
	return grpcClients, nil
}

func initAuthClient(grpcClients *GRPCClients, cfg *config.GRPCClientConfig, tokenCfg *tokenx.Config) error {
	if cfg.AuthAddr == "" {
		return fmt.Errorf("auth gRPC address is empty")
	}
	client, err := authconn.Dial(authconn.GRPCConfig{
		Addr:           cfg.AuthAddr,
		TokenSecretKey: tokenCfg.SecretKey,
		TokenIssuer:    tokenCfg.Issuer,
		CallerService:  "system-service",
	})
	if err != nil {
		return err
	}
	grpcClients.auth = client
	return nil
}

func (c *GRPCClients) BootstrapOwner(ctx context.Context, username, password, email, phone string) (*accountpb.BootstrapOwnerResponse, error) {
	if c.auth == nil || c.auth.Account == nil {
		return nil, fmt.Errorf("auth account client not initialized")
	}
	return c.auth.Account.BootstrapOwner(ctx, username, password, email, phone)
}

func (c *GRPCClients) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	var errs []error
	if c.auth != nil {
		if err := c.auth.Close(); err != nil {
			errs = append(errs, err)
		}
		c.auth = nil
	}
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
