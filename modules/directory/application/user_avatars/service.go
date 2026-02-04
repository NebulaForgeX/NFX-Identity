package user_avatars

import (
	userAvatarDomain "nfxid/modules/directory/domain/user_avatars"
	directoryGrpc "nfxid/modules/directory/infrastructure/grpc"
	"nfxid/pkgs/kafkax/eventbus"
)

type Service struct {
	userAvatarRepo *userAvatarDomain.Repo
	grpcClients    *directoryGrpc.GRPCClients
	busPublisher   *eventbus.BusPublisher
}

func NewService(
	userAvatarRepo *userAvatarDomain.Repo,
	grpcClients *directoryGrpc.GRPCClients,
	busPublisher *eventbus.BusPublisher,
) *Service {
	return &Service{
		userAvatarRepo: userAvatarRepo,
		grpcClients:    grpcClients,
		busPublisher:   busPublisher,
	}
}
