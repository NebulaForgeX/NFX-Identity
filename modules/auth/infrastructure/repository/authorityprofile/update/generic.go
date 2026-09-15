package update
import ("context"; "nfxidentity/modules/auth/domain/authorityprofile"; "nfxidentity/modules/auth/infrastructure/repository/authorityprofile/mapper")
func (h *Handler) Generic(ctx context.Context, p *authorityprofile.Profile) error {
	return h.db.WithContext(ctx).Save(mapper.ToModel(p)).Error
}
