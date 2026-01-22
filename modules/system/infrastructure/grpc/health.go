package grpc

import (
	"context"
	"fmt"
	"sync"
	"time"

	"nfxid/connections/common"
	"nfxid/constants"
	"nfxid/modules/system/config"
	"nfxid/pkgs/logx"
	"nfxid/pkgs/security/token/servertoken"
	healthpb "nfxid/protos/gen/common/health"

	"google.golang.org/grpc"
)

// HealthChecker 健康检查客户端管理器
type HealthChecker struct {
	healthClients map[string]*common.HealthClient // serviceName -> healthClient
	conns         []*grpc.ClientConn
	mu            sync.Mutex
}

// NewHealthChecker 创建健康检查客户端管理器
func NewHealthChecker() *HealthChecker {
	return &HealthChecker{
		healthClients: make(map[string]*common.HealthClient),
		conns:         make([]*grpc.ClientConn, 0),
	}
}

// initHealthClients 初始化所有服务的健康检查客户端
func initHealthClients(grpcClients *GRPCClients, cfg *config.GRPCClientConfig, serverCfg *config.ServerConfig, tokenProvider servertoken.TokenProvider) error {
	healthChecker := NewHealthChecker()
	
	// 构建 system 服务地址（自身）
	systemAddr := fmt.Sprintf("%s:%d", serverCfg.Host, serverCfg.GRPCPort)

	// 为所有服务创建健康检查连接（需要 token，因为这是正常的 gRPC 请求）
	serviceAddrs := map[string]string{
		constants.ServiceAccess:    cfg.AccessAddr,
		constants.ServiceAudit:     cfg.AuditAddr,
		constants.ServiceAuth:      cfg.AuthAddr,
		constants.ServiceClients:   cfg.ClientsAddr,
		constants.ServiceDirectory: cfg.DirectoryAddr,
		constants.ServiceImage:     cfg.ImageAddr,
		constants.ServiceSystem:    systemAddr,
		constants.ServiceTenants:   cfg.TenantsAddr,
	}

	for serviceName, addr := range serviceAddrs {
		if addr == "" {
			continue
		}
		// 健康检查连接需要 token（这是正常的 gRPC 请求，不是标准健康检查协议）
		healthConn, err := createConnection(addr, tokenProvider)
		if err != nil {
			logx.S().Warnf("Failed to create health check connection for %s: %v", serviceName, err)
			continue // 继续创建其他连接
		}
		healthChecker.healthClients[serviceName] = common.NewHealthClient(healthpb.NewHealthServiceClient(healthConn))
		healthChecker.conns = append(healthChecker.conns, healthConn)
		grpcClients.conns = append(grpcClients.conns, healthConn)
	}

	grpcClients.healthChecker = healthChecker
	return nil
}

// CheckAllServicesHealth 检查所有 8 个服务的健康状态（包括基础设施：数据库、Redis等）
// 返回服务名 -> 健康状态响应的映射
func (c *GRPCClients) CheckAllServicesHealth(ctx context.Context) (map[string]*healthpb.GetHealthResponse, error) {
	if c.healthChecker == nil {
		return nil, fmt.Errorf("health checker not initialized")
	}

	c.healthChecker.mu.Lock()
	defer c.healthChecker.mu.Unlock()

	results := make(map[string]*healthpb.GetHealthResponse)
	services := constants.AllServices()

	for _, serviceName := range services {
		healthClient, exists := c.healthChecker.healthClients[serviceName]
		if !exists {
			logx.S().Warnf("No health check client for service: %s", serviceName)
			results[serviceName] = &healthpb.GetHealthResponse{
				Healthy:     false,
				ServiceName: serviceName,
				CheckedAt:   time.Now().Unix(),
			}
			continue
		}

		// 创建带超时的 context
		checkCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
		healthResp, err := healthClient.GetHealth(checkCtx)
		cancel()

		if err != nil {
			logx.S().Warnf("Health check failed for %s: %v", serviceName, err)
			results[serviceName] = &healthpb.GetHealthResponse{
				Healthy:     false,
				ServiceName: serviceName,
				CheckedAt:   time.Now().Unix(),
			}
		} else {
			results[serviceName] = healthResp
		}
	}

	return results, nil
}
