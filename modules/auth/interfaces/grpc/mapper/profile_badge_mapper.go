package mapper

import (
	badgeAppViews "nfxid/modules/auth/application/badge/views"
	profileBadgeAppViews "nfxid/modules/auth/application/profile_badge/views"
	badgepb "nfxid/protos/gen/auth/badge"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// ProfileBadgeViewToProto 将 ProfileBadgeView 转换为 proto ProfileBadge 消息
func ProfileBadgeViewToProto(v *profileBadgeAppViews.ProfileBadgeView, badgeView *badgeAppViews.BadgeView) *badgepb.ProfileBadge {
	if v == nil {
		return nil
	}

	profileBadge := &badgepb.ProfileBadge{
		Id:        v.ID.String(),
		ProfileId: v.ProfileID.String(),
		BadgeId:   v.BadgeID.String(),
		EarnedAt:  timestamppb.New(v.EarnedAt),
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.Description != nil {
		profileBadge.Description = v.Description
	}
	if v.Level != nil {
		level := int32(*v.Level)
		profileBadge.Level = &level
	}

	// Fill Badge information if provided
	if badgeView != nil {
		profileBadge.Badge = BadgeViewToProto(badgeView)
	}

	return profileBadge
}
