package update
import ("context"; "nfxidentity/modules/auth/domain/identity"; "nfxidentity/modules/auth/infrastructure/repository/identity/mapper")
func (h *Handler) Generic(ctx context.Context, i *identity.Identity) error {
	return h.db.WithContext(ctx).Save(mapper.ToModel(i)).Error
}
