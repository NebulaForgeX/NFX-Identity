package mapper

import (
	profileDomain "nfxidentity/modules/auth/domain/profile"
	accountpb "nfxidentity/protos/gen/auth/account"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func AuthorityProfileAvatarToProto(a *profileDomain.AuthorityProfileAvatar) *accountpb.AuthorityProfileAvatar {
	if a == nil {
		return nil
	}
	out := &accountpb.AuthorityProfileAvatar{
		Id:        a.ID().String(),
		ProfileId: a.ProfileID().String(),
		ImageId:   a.ImageID().String(),
		IsActive:  a.IsActive(),
		CreatedAt: timestamppb.New(a.CreatedAt()),
		UpdatedAt: timestamppb.New(a.UpdatedAt()),
	}
	if t := a.DeletedAt(); t != nil {
		out.DeletedAt = timestamppb.New(*t)
	}
	return out
}
