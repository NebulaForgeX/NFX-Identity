package get
import ("context"; "errors"; "nfxidentity/modules/auth/domain/forgerprofile"; "nfxidentity/modules/auth/infrastructure/rdb/models"; "nfxidentity/modules/auth/infrastructure/repository/forgerprofile/mapper"; "nfxidentity/pkgs/errx"; "github.com/google/uuid"; "gorm.io/gorm")
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*forgerprofile.Profile, error) {
	var m models.ForgerProfile
	if err := h.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { return nil, errx.NotFound("NOT_FOUND", "profile not found") }
		return nil, err
	}
	return mapper.ToDomain(&m), nil
}
