package get
import (
	"context"; "errors"
	"nfxidentity/modules/auth/domain/email"
	"nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/modules/auth/infrastructure/repository/email/mapper"
	"nfxidentity/pkgs/errx"
	"github.com/google/uuid"
	"gorm.io/gorm"
)
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (*email.Email, error) {
	var m models.Email
	if err := h.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { return nil, errx.NotFound("NOT_FOUND", "email not found") }
		return nil, err
	}
	return mapper.ToDomain(&m), nil
}
