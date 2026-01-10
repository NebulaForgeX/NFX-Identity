package mapper

import (
	userBadgeAppResult "nfxid/modules/directory/application/user_badges/results"
	userbadgepb "nfxid/protos/gen/directory/user_badge"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserBadgeROToProto 将 UserBadgeRO 转换为 proto UserBadge 消息
func UserBadgeROToProto(v *userBadgeAppResult.UserBadgeRO) *userbadgepb.UserBadge {
	if v == nil {
		return nil
	}

	userBadge := &userbadgepb.UserBadge{
		Id:        v.ID.String(),
		UserId:    v.UserID.String(),
		BadgeId:   v.BadgeID.String(),
		EarnedAt:  timestamppb.New(v.EarnedAt),
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.Description != "" {
		userBadge.Description = &v.Description
	}

	if v.Level > 0 {
		level := int32(v.Level)
		userBadge.Level = &level
	}

	return userBadge
}

// UserBadgeListROToProto 批量转换 UserBadgeRO 到 proto UserBadge
func UserBadgeListROToProto(results []userBadgeAppResult.UserBadgeRO) []*userbadgepb.UserBadge {
	userBadges := make([]*userbadgepb.UserBadge, len(results))
	for i, v := range results {
		userBadges[i] = UserBadgeROToProto(&v)
	}
	return userBadges
}
