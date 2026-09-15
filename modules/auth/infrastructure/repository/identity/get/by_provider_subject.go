package get
import ("context"; "errors"; "nfxidentity/modules/auth/domain/identity"; "nfxidentity/modules/auth/infrastructure/rdb/models"; "nfxidentity/modules/auth/infrastructure/repository/identity/mapper"; "nfxidentity/pkgs/errx"; "gorm.io/gorm")
func (h *Handler) ByProviderSubject(ctx context.Context, provider, subject string) (*identity.Identity, error) {
	var m models.Identity
	if err := h.db.WithContext(ctx).Where("identity_provider = ? AND provider_subject = ? AND deleted_at IS NULL", provider, subject).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { return nil, errx.NotFound("NOT_FOUND", "identity not found") }
		return nil, err
	}
	return mapper.ToDomain(&m), nil
}
