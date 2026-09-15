package create
import ("context"; "nfxidentity/modules/auth/domain/phone"; "nfxidentity/modules/auth/infrastructure/repository/phone/mapper")
func (h *Handler) New(ctx context.Context, p *phone.Phone) error {
	return h.db.WithContext(ctx).Create(mapper.ToModel(p)).Error
}
