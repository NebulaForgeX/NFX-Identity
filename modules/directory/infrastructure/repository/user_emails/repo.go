package user_emails

import (
	"nfxidentity/modules/directory/domain/user_emails"
	"nfxidentity/modules/directory/infrastructure/repository/user_emails/check"
	"nfxidentity/modules/directory/infrastructure/repository/user_emails/create"
	"nfxidentity/modules/directory/infrastructure/repository/user_emails/delete"
	"nfxidentity/modules/directory/infrastructure/repository/user_emails/get"
	"nfxidentity/modules/directory/infrastructure/repository/user_emails/update"

	"gorm.io/gorm"
)

// NewRepo 创建一个新的 UserEmail repository
func NewRepo(db *gorm.DB) *user_emails.Repo {
	return &user_emails.Repo{
		Create: create.NewHandler(db),
		Get:    get.NewHandler(db),
		Check:  check.NewHandler(db),
		Update: update.NewHandler(db),
		Delete: delete.NewHandler(db),
	}
}
