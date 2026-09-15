package email
import (
	"nfxidentity/modules/auth/domain/email"
	"nfxidentity/modules/auth/infrastructure/repository/email/create"
	"nfxidentity/modules/auth/infrastructure/repository/email/get"
	"nfxidentity/modules/auth/infrastructure/repository/email/update"
	"gorm.io/gorm"
)
func NewRepo(db *gorm.DB) *email.Repo {
	return &email.Repo{Create: create.NewHandler(db), Get: get.NewHandler(db), Update: update.NewHandler(db)}
}
