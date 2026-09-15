package get
import ("context"; "errors"; "nfxidentity/modules/auth/domain/authorityprofile"; "nfxidentity/modules/auth/infrastructure/rdb/models"; "nfxidentity/modules/auth/infrastructure/repository/authorityprofile/mapper"; "nfxidentity/pkgs/errx"; "github.com/google/uuid"; "gorm.io/gorm")
func (h *Handler) ByAccountAndID(ctx context.Context, accountID, id uuid.UUID) (*authorityprofile.Profile, error) {
	var m models.AuthorityProfile
	if err := h.db.WithContext(ctx).Where("id = ? AND account_id = ? AND deleted_at IS NULL", id, accountID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { return nil, errx.NotFound("NOT_FOUND", "profile not found") }
		return nil, err
	}
	return mapper.ToDomain(&m), nil
}
