package create
import ("context"; "nfxidentity/modules/auth/domain/identity"; "nfxidentity/modules/auth/infrastructure/repository/identity/mapper")
func (h *Handler) New(ctx context.Context, i *identity.Identity) error {
	return h.db.WithContext(ctx).Create(mapper.ToModel(i)).Error
}
