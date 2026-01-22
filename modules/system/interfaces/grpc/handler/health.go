package handler

import (
	"context"
	"time"

	"nfxid/modules/system/application/resource"
	healthpb "nfxid/protos/gen/common/health"
)

// HealthHandler 健康检查处理器
type HealthHandler struct {
	healthpb.UnimplementedHealthServiceServer
	resourceSvc *resource.Service
	serviceName string
}

// NewHealthHandler 创建健康检查处理器
func NewHealthHandler(resourceSvc *resource.Service, serviceName string) *HealthHandler {
	return &HealthHandler{
		resourceSvc: resourceSvc,
		serviceName: serviceName,
	}
}

// GetHealth 获取服务健康状态（包括数据库、Redis等基础设施）
func (h *HealthHandler) GetHealth(ctx context.Context, req *healthpb.GetHealthRequest) (*healthpb.GetHealthResponse, error) {
	// 构建基础设施健康状态
	infra := &healthpb.InfrastructureHealth{
		Others: make(map[string]*healthpb.ResourceHealth),
	}

	allHealthy := true
	now := time.Now().Unix()

	// 检查 PostgreSQL 数据库健康状态
	postgresErr := h.resourceSvc.CheckPostgres(ctx)
	dbHealth := &healthpb.ResourceHealth{
		Healthy:   postgresErr == nil,
		CheckedAt: now,
	}
	if postgresErr != nil {
		errMsg := postgresErr.Error()
		dbHealth.ErrorMessage = &errMsg
		allHealthy = false
	}
	infra.Database = dbHealth

	// 检查 Redis 健康状态
	redisErr := h.resourceSvc.CheckRedis(ctx)
	redisHealth := &healthpb.ResourceHealth{
		Healthy:   redisErr == nil,
		CheckedAt: now,
	}
	if redisErr != nil {
		errMsg := redisErr.Error()
		redisHealth.ErrorMessage = &errMsg
		allHealthy = false
	}
	infra.Redis = redisHealth

	// 检查 Kafka 健康状态
	kafkaErr := h.resourceSvc.CheckKafka(ctx)
	if kafkaErr != nil {
		errMsg := kafkaErr.Error()
		infra.Others["kafka"] = &healthpb.ResourceHealth{
			Healthy:     false,
			ErrorMessage: &errMsg,
			CheckedAt:   now,
		}
		allHealthy = false
	} else {
		infra.Others["kafka"] = &healthpb.ResourceHealth{
			Healthy:   true,
			CheckedAt: now,
		}
	}

	// 检查 RabbitMQ 健康状态
	rabbitMQErr := h.resourceSvc.CheckRabbitMQ(ctx)
	if rabbitMQErr != nil {
		errMsg := rabbitMQErr.Error()
		infra.Others["rabbitmq"] = &healthpb.ResourceHealth{
			Healthy:     false,
			ErrorMessage: &errMsg,
			CheckedAt:   now,
		}
		allHealthy = false
	} else {
		infra.Others["rabbitmq"] = &healthpb.ResourceHealth{
			Healthy:   true,
			CheckedAt: now,
		}
	}

	return &healthpb.GetHealthResponse{
		Healthy:        allHealthy,
		Infrastructure: infra,
		ServiceName:    h.serviceName,
		CheckedAt:      now,
	}, nil
}
