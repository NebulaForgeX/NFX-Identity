package mapper

import (
	profileDomain "nfxidentity/modules/auth/domain/profile"
	accountpb "nfxidentity/protos/gen/auth/account"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func AuthorityProfileBackgroundToProto(b *profileDomain.AuthorityProfileBackground) *accountpb.AuthorityProfileBackground {
	if b == nil {
		return nil
	}
	out := &accountpb.AuthorityProfileBackground{
		Id:        b.ID().String(),
		ProfileId: b.ProfileID().String(),
		ImageId:   b.ImageID().String(),
		SortOrder: int32(b.SortOrder()),
		CreatedAt: timestamppb.New(b.CreatedAt()),
		UpdatedAt: timestamppb.New(b.UpdatedAt()),
	}
	if t := b.DeletedAt(); t != nil {
		out.DeletedAt = timestamppb.New(*t)
	}
	return out
}

func AuthorityProfileBackgroundsToProto(items []*profileDomain.AuthorityProfileBackground) []*accountpb.AuthorityProfileBackground {
	if len(items) == 0 {
		return nil
	}
	out := make([]*accountpb.AuthorityProfileBackground, 0, len(items))
	for _, item := range items {
		out = append(out, AuthorityProfileBackgroundToProto(item))
	}
	return out
}
