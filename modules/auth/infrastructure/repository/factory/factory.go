package factory

import (
	imagesDomain "nfxidentity/modules/asset/domain/images"
	imagesRepo "nfxidentity/modules/asset/infrastructure/repository/images"
	accountDomain "nfxidentity/modules/auth/domain/account"
	emailDomain "nfxidentity/modules/auth/domain/email"
	identityDomain "nfxidentity/modules/auth/domain/identity"
	phoneDomain "nfxidentity/modules/auth/domain/phone"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	refreshDomain "nfxidentity/modules/auth/domain/refreshtoken"
	accountRepo "nfxidentity/modules/auth/infrastructure/repository/account"
	emailRepo "nfxidentity/modules/auth/infrastructure/repository/email"
	identityRepo "nfxidentity/modules/auth/infrastructure/repository/identity"
	phoneRepo "nfxidentity/modules/auth/infrastructure/repository/phone"
	profileRepo "nfxidentity/modules/auth/infrastructure/repository/profile"
	refreshRepo "nfxidentity/modules/auth/infrastructure/repository/refreshtoken"
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
func (f *TxRepoFactory) Profile(uow transaction.UoW) *profileDomain.Repo {
	return profileRepo.NewRepo(f.dbOr(uow))
}
func (f *TxRepoFactory) Image(uow transaction.UoW) *imagesDomain.Repo {
	return imagesRepo.NewRepo(f.dbOr(uow))
}
