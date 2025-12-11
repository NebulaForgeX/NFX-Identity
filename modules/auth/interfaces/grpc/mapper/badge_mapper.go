package mapper

import (
	badgeAppViews "nfxid/modules/auth/application/badge/views"
	badgepb "nfxid/protos/gen/auth/badge"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// BadgeViewToProto 将 BadgeView 转换为 proto Badge 消息
func BadgeViewToProto(v *badgeAppViews.BadgeView) *badgepb.Badge {
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

	return badge
}
