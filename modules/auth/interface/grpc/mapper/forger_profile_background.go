package mapper

import (
	profileDomain "nfxidentity/modules/auth/domain/profile"
	accountpb "nfxidentity/protos/gen/auth/account"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ForgerProfileBackgroundToProto(b *profileDomain.ForgerProfileBackground) *accountpb.ForgerProfileBackground {
	if b == nil {
		return nil
	}
	out := &accountpb.ForgerProfileBackground{
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

func ForgerProfileBackgroundsToProto(items []*profileDomain.ForgerProfileBackground) []*accountpb.ForgerProfileBackground {
	if len(items) == 0 {
		return nil
	}
	out := make([]*accountpb.ForgerProfileBackground, 0, len(items))
	for _, item := range items {
		out = append(out, ForgerProfileBackgroundToProto(item))
	}
	return out
}
