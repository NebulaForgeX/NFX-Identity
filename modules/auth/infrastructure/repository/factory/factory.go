package factory

import (
	imagesDomain "nfxidentity/modules/asset/domain/images"
	imagesRepo "nfxidentity/modules/asset/infrastructure/repository/images"
	accountDomain "nfxidentity/modules/auth/domain/account"
	authorityDomain "nfxidentity/modules/auth/domain/authorityprofile"
	avatarDomain "nfxidentity/modules/auth/domain/avatar"
	backgroundDomain "nfxidentity/modules/auth/domain/background"
	emailDomain "nfxidentity/modules/auth/domain/email"
	forgerDomain "nfxidentity/modules/auth/domain/forgerprofile"
	identityDomain "nfxidentity/modules/auth/domain/identity"
	phoneDomain "nfxidentity/modules/auth/domain/phone"
	refreshDomain "nfxidentity/modules/auth/domain/refreshtoken"
	settingsDomain "nfxidentity/modules/auth/domain/settings"
	accountRepo "nfxidentity/modules/auth/infrastructure/repository/account"
	authorityRepo "nfxidentity/modules/auth/infrastructure/repository/authorityprofile"
	avatarRepo "nfxidentity/modules/auth/infrastructure/repository/avatar"
	backgroundRepo "nfxidentity/modules/auth/infrastructure/repository/background"
	emailRepo "nfxidentity/modules/auth/infrastructure/repository/email"
	forgerRepo "nfxidentity/modules/auth/infrastructure/repository/forgerprofile"
	identityRepo "nfxidentity/modules/auth/infrastructure/repository/identity"
	phoneRepo "nfxidentity/modules/auth/infrastructure/repository/phone"
	refreshRepo "nfxidentity/modules/auth/infrastructure/repository/refreshtoken"
	settingsRepo "nfxidentity/modules/auth/infrastructure/repository/settings"
	"nfxidentity/pkgs/transaction"

	"gorm.io/gorm"
)

type TxRepoFactory struct {
	db *gorm.DB
}

func NewTxRepoFactory(db *gorm.DB) *TxRepoFactory {
	return &TxRepoFactory{db: db}
}

func (f *TxRepoFactory) dbOr(uow transaction.UoW) *gorm.DB {
	if uow.DB != nil {
		return uow.DB
	}
	return f.db
}

func (f *TxRepoFactory) Account(uow transaction.UoW) *accountDomain.Repo {
	return accountRepo.NewRepo(f.dbOr(uow))
}
func (f *TxRepoFactory) Email(uow transaction.UoW) *emailDomain.Repo {
	return emailRepo.NewRepo(f.dbOr(uow))
}
func (f *TxRepoFactory) Phone(uow transaction.UoW) *phoneDomain.Repo {
	return phoneRepo.NewRepo(f.dbOr(uow))
}
func (f *TxRepoFactory) Identity(uow transaction.UoW) *identityDomain.Repo {
	return identityRepo.NewRepo(f.dbOr(uow))
}
func (f *TxRepoFactory) RefreshToken(uow transaction.UoW) *refreshDomain.Repo {
	return refreshRepo.NewRepo(f.dbOr(uow))
}
func (f *TxRepoFactory) Forger(uow transaction.UoW) *forgerDomain.Repo {
	return forgerRepo.NewRepo(f.dbOr(uow))
}
func (f *TxRepoFactory) Authority(uow transaction.UoW) *authorityDomain.Repo {
	return authorityRepo.NewRepo(f.dbOr(uow))
}
func (f *TxRepoFactory) Avatar(uow transaction.UoW) *avatarDomain.Repo {
	return avatarRepo.NewRepo(f.dbOr(uow))
}
func (f *TxRepoFactory) Background(uow transaction.UoW) *backgroundDomain.Repo {
	return backgroundRepo.NewRepo(f.dbOr(uow))
}
func (f *TxRepoFactory) Settings(uow transaction.UoW) *settingsDomain.Repo {
	return settingsRepo.NewRepo(f.dbOr(uow))
}
func (f *TxRepoFactory) Image(uow transaction.UoW) *imagesDomain.Repo {
	return imagesRepo.NewRepo(f.dbOr(uow))
}
