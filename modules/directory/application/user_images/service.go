package user_images

import (
	userImageDomain "nfxid/modules/directory/domain/user_images"
	directoryGrpc "nfxid/modules/directory/infrastructure/grpc"
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
