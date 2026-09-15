package get
import ("context"; "nfxidentity/modules/auth/domain/authorityprofile"; "nfxidentity/modules/auth/infrastructure/rdb/models"; "nfxidentity/modules/auth/infrastructure/repository/authorityprofile/mapper"; "github.com/google/uuid")
func (h *Handler) ByAccountID(ctx context.Context, accountID uuid.UUID) ([]*authorityprofile.Profile, error) {
	var rows []models.AuthorityProfile
	if err := h.db.WithContext(ctx).Where("account_id = ? AND deleted_at IS NULL", accountID).Find(&rows).Error; err != nil { return nil, err }
	out := make([]*authorityprofile.Profile, 0, len(rows))
	for i := range rows { out = append(out, mapper.ToDomain(&rows[i])) }
	return out, nil
}
