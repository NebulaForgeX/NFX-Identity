package identity
import (
	"nfxidentity/modules/auth/domain/identity"
	"nfxidentity/modules/auth/infrastructure/repository/identity/create"
	"nfxidentity/modules/auth/infrastructure/repository/identity/get"
	"nfxidentity/modules/auth/infrastructure/repository/identity/update"
	"gorm.io/gorm"
)
func NewRepo(db *gorm.DB) *identity.Repo {
	return &identity.Repo{Create: create.NewHandler(db), Get: get.NewHandler(db), Update: update.NewHandler(db)}
}
