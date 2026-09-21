package login

import (
	repofactory "nfxidentity/modules/auth/infrastructure/repository/factory"
	profileQuery "nfxidentity/modules/auth/query/profile"
	"nfxidentity/pkgs/kafkax/eventbus"
	"nfxidentity/pkgs/tokenx"
	"nfxidentity/pkgs/transaction"
)

type Service struct {
	repoFactory  *repofactory.TxRepoFactory
	txManager    transaction.TxManager
	bus          *eventbus.BusPublisher
	tokenx       *tokenx.Tokenx
	profileQuery *profileQuery.Query
}

func NewService(
	repoFactory *repofactory.TxRepoFactory,
	txManager transaction.TxManager,
	bus *eventbus.BusPublisher,
	tokenxInst *tokenx.Tokenx,
	profileQuery *profileQuery.Query,
) *Service {
	return &Service{
		repoFactory:  repoFactory,
		txManager:    txManager,
		bus:          bus,
		tokenx:       tokenxInst,
		profileQuery: profileQuery,
	}
}
