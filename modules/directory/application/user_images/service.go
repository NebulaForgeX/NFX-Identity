package user_images

import (
	directoryGrpc "nfxid/modules/directory/infrastructure/grpc"
	userImageDomain "nfxid/modules/directory/domain/user_images"
)

type Service struct {
	userImageRepo *userImageDomain.Repo
	grpcClients   *directoryGrpc.GRPCClients
}

func NewService(
	userImageRepo *userImageDomain.Repo,
	grpcClients *directoryGrpc.GRPCClients,
) *Service {
	return &Service{
		userImageRepo: userImageRepo,
		grpcClients:   grpcClients,
	}
}
