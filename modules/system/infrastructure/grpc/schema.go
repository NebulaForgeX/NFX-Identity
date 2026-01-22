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
	schemapb "nfxid/protos/gen/common/schema"

	"google.golang.org/grpc"
)

// SchemaChecker schema 清空客户端管理器
type SchemaChecker struct {
	schemaClients map[string]*common.SchemaClient // serviceName -> schemaClient
	conns         []*grpc.ClientConn
	mu            sync.Mutex
}

// NewSchemaChecker 创建 schema 清空客户端管理器
func NewSchemaChecker() *SchemaChecker {
	return &SchemaChecker{
		schemaClients: make(map[string]*common.SchemaClient),
		conns:         make([]*grpc.ClientConn, 0),
	}
}

// initSchemaClients 初始化所有服务的 schema 清空客户端
func initSchemaClients(grpcClients *GRPCClients, cfg *config.GRPCClientConfig, serverCfg *config.ServerConfig, tokenProvider servertoken.TokenProvider) error {
	schemaChecker := NewSchemaChecker()
	
	// 构建 system 服务地址（自身）
	systemAddr := fmt.Sprintf("%s:%d", serverCfg.Host, serverCfg.GRPCPort)

	// 为所有服务创建 schema 清空连接
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
		// Schema 清空连接需要 token
		schemaConn, err := createConnection(addr, tokenProvider)
		if err != nil {
			logx.S().Warnf("Failed to create schema clear connection for %s: %v", serviceName, err)
			continue // 继续创建其他连接
		}
		schemaChecker.schemaClients[serviceName] = common.NewSchemaClient(schemapb.NewSchemaServiceClient(schemaConn))
		schemaChecker.conns = append(schemaChecker.conns, schemaConn)
		grpcClients.conns = append(grpcClients.conns, schemaConn)
	}

	grpcClients.schemaChecker = schemaChecker
	return nil
}

// ClearAllSchemas 清空所有 8 个服务的 schema（清空所有表数据，不删除表）
// 返回服务名 -> 清空结果的映射
func (c *GRPCClients) ClearAllSchemas(ctx context.Context) (map[string]*schemapb.ClearSchemaResponse, error) {
	if c.schemaChecker == nil {
		return nil, fmt.Errorf("schema checker not initialized")
	}

	c.schemaChecker.mu.Lock()
	defer c.schemaChecker.mu.Unlock()

	results := make(map[string]*schemapb.ClearSchemaResponse)
	services := constants.AllServices()

	for _, serviceName := range services {
		schemaClient, exists := c.schemaChecker.schemaClients[serviceName]
		if !exists {
			logx.S().Warnf("No schema clear client for service: %s", serviceName)
			errMsg := fmt.Sprintf("schema clear client not found for service: %s", serviceName)
			results[serviceName] = &schemapb.ClearSchemaResponse{
				Success:      false,
				ErrorMessage: &errMsg,
				TablesCleared: 0,
			}
			continue
		}

		// 创建带超时的 context
		clearCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		clearResp, err := schemaClient.ClearSchema(clearCtx)
		cancel()

		if err != nil {
			logx.S().Warnf("Schema clear failed for %s: %v", serviceName, err)
			errMsg := err.Error()
			results[serviceName] = &schemapb.ClearSchemaResponse{
				Success:      false,
				ErrorMessage: &errMsg,
				TablesCleared: 0,
			}
		} else {
			results[serviceName] = clearResp
			if clearResp.Success {
				logx.S().Infof("✅ Cleared schema for %s: %d tables cleared", serviceName, clearResp.TablesCleared)
			} else {
				errMsg := "unknown error"
				if clearResp.ErrorMessage != nil {
					errMsg = *clearResp.ErrorMessage
				}
				logx.S().Warnf("⚠️  Schema clear failed for %s: %s", serviceName, errMsg)
			}
		}
	}

	return results, nil
}
