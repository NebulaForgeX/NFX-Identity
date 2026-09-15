package get
import ("context"; "errors"; "nfxidentity/modules/auth/domain/identity"; "nfxidentity/modules/auth/infrastructure/rdb/models"; "nfxidentity/modules/auth/infrastructure/repository/identity/mapper"; "nfxidentity/pkgs/errx"; "github.com/google/uuid"; "gorm.io/gorm")
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*identity.Identity, error) {
	var m models.Identity
	if err := h.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { return nil, errx.NotFound("NOT_FOUND", "identity not found") }
		return nil, err
	}
	return mapper.ToDomain(&m), nil
}
