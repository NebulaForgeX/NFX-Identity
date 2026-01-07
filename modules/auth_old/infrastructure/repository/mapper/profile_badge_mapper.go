package mapper

import (
	"nfxid/modules/auth/domain/profile_badge"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

func ProfileBadgeDomainToModel(pb *profile_badge.ProfileBadge) *models.ProfileBadge {
	if pb == nil {
		return nil
	}

	editable := pb.Editable()
	return &models.ProfileBadge{
		ID:          pb.ID(),
		ProfileID:   pb.ProfileID(),
		BadgeID:     pb.BadgeID(),
		Description: editable.Description,
		Level:       editable.Level,
		EarnedAt:    pb.EarnedAt(),
		CreatedAt:   pb.CreatedAt(),
		UpdatedAt:   pb.UpdatedAt(),
	}
}

func ProfileBadgeModelToDomain(m *models.ProfileBadge) *profile_badge.ProfileBadge {
	if m == nil {
		return nil
	}

	editable := profile_badge.ProfileBadgeEditable{
		Description: m.Description,
		Level:       m.Level,
	}

	state := profile_badge.ProfileBadgeState{
		ID:        m.ID,
		ProfileID: m.ProfileID,
		BadgeID:   m.BadgeID,
		Editable:  editable,
		EarnedAt:  m.EarnedAt,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}

	return profile_badge.NewProfileBadgeFromState(state)
}

func ProfileBadgeModelsToUpdates(m *models.ProfileBadge) map[string]any {
	return map[string]any{
		models.ProfileBadgeCols.Description: m.Description,
		models.ProfileBadgeCols.Level:       m.Level,
		models.ProfileBadgeCols.EarnedAt:    m.EarnedAt,
	}
}
