package profile

import (
	profileQueries "nfxid/modules/auth/application/profile/queries"
	profileDomain "nfxid/modules/auth/domain/profile"
	"nfxid/modules/auth/infrastructure/grpcclient"
)

type Service struct {
	profileRepo     profileDomain.Repo
	profileQuery    profileQueries.ProfileQuery
	imageGRPCClient *grpcclient.ImageGRPCClient
}

func NewService(
	profileRepo profileDomain.Repo,
	profileQuery profileQueries.ProfileQuery,
	imageGRPCClient *grpcclient.ImageGRPCClient,
) *Service {
	return &Service{
		profileRepo:     profileRepo,
		profileQuery:    profileQuery,
		imageGRPCClient: imageGRPCClient,
	}
}
