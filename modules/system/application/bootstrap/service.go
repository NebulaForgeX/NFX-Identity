package bootstrap

import (
	systemStateDomain "nfxid/modules/system/domain/system_state"
	grpcClients "nfxid/modules/system/infrastructure/grpc"
)

// Service 系统初始化服务
// 负责系统首次初始化时的所有基础数据创建
type Service struct {
	systemStateRepo *systemStateDomain.Repo
	grpcClients     *grpcClients.GRPCClients // gRPC 客户端（通过依赖注入）
}

func NewService(
	systemStateRepo *systemStateDomain.Repo,
	grpcClients *grpcClients.GRPCClients,
) *Service {
	return &Service{
		systemStateRepo: systemStateRepo,
		grpcClients:     grpcClients,
	}
}
