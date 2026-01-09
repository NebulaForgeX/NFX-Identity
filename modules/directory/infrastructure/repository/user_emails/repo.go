package user_emails

import (
	"nfxid/modules/directory/domain/user_emails"
	"nfxid/modules/directory/infrastructure/repository/user_emails/check"
	"nfxid/modules/directory/infrastructure/repository/user_emails/create"
	"nfxid/modules/directory/infrastructure/repository/user_emails/delete"
	"nfxid/modules/directory/infrastructure/repository/user_emails/get"
	"nfxid/modules/directory/infrastructure/repository/user_emails/update"

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
