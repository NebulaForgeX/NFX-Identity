package profile

import (
	profileDomain "nfxid/modules/auth/domain/profile"
	"nfxid/modules/auth/infrastructure/grpcclient"
)

type Service struct {
	profileRepo     *profileDomain.Repo
	profileQuery    profileDomain.Query
	imageGRPCClient *grpcclient.ImageGRPCClient
}

func NewService(
	profileRepo *profileDomain.Repo,
	profileQuery profileDomain.Query,
	imageGRPCClient *grpcclient.ImageGRPCClient,
) *Service {
	return &Service{
		profileRepo:     profileRepo,
		profileQuery:    profileQuery,
		imageGRPCClient: imageGRPCClient,
	}
}
