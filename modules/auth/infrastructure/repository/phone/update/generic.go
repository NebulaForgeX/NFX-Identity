package update
import ("context"; "nfxidentity/modules/auth/domain/phone"; "nfxidentity/modules/auth/infrastructure/repository/phone/mapper")
func (h *Handler) Generic(ctx context.Context, p *phone.Phone) error {
	return h.db.WithContext(ctx).Save(mapper.ToModel(p)).Error
}
