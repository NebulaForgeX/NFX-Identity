package mapper

import (
	badgeAppResult "nfxid/modules/directory/application/badges/results"
	badgepb "nfxid/protos/gen/directory/badge"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// BadgeROToProto 将 BadgeRO 转换为 proto Badge 消息
func BadgeROToProto(v *badgeAppResult.BadgeRO) *badgepb.Badge {
	if v == nil {
		return nil
	}

	badge := &badgepb.Badge{
		Id:        v.ID.String(),
		Name:      v.Name,
		IsSystem:  v.IsSystem,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.Description != nil {
		badge.Description = v.Description
	}

	if v.IconURL != nil {
		badge.IconUrl = v.IconURL
	}

	if v.Color != nil {
		badge.Color = v.Color
	}

	if v.Category != nil {
		badge.Category = v.Category
	}

	if v.DeletedAt != nil {
		badge.DeletedAt = timestamppb.New(*v.DeletedAt)
	}

	return badge
}

// BadgeListROToProto 批量转换 BadgeRO 到 proto Badge
func BadgeListROToProto(results []badgeAppResult.BadgeRO) []*badgepb.Badge {
	badges := make([]*badgepb.Badge, len(results))
	for i, v := range results {
		badges[i] = BadgeROToProto(&v)
	}
	return badges
}
