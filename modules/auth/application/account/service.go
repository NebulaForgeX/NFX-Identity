package account

import (
	"context"

	repofactory "nfxidentity/modules/auth/infrastructure/repository/factory"
	accountQuery "nfxidentity/modules/auth/query/account"
	emailQuery "nfxidentity/modules/auth/query/email"
	phoneQuery "nfxidentity/modules/auth/query/phone"
	profileQuery "nfxidentity/modules/auth/query/profile"
	"nfxidentity/pkgs/email"
	"nfxidentity/pkgs/kafkax/eventbus"
	"nfxidentity/pkgs/tokenx"
	"nfxidentity/pkgs/transaction"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	tx          transaction.TxManager
	repoFactory *repofactory.TxRepoFactory
	accounts    *accountQuery.Query
	emails      *emailQuery.Query
	phones      *phoneQuery.Query
	profiles    *profileQuery.Query
	tokens      *tokenx.Tokenx
	redis       *redis.Client
	mail        *email.EmailService
	bus         *eventbus.BusPublisher
	checkFn     func(ctx context.Context, email, code string) bool
}

func NewService(
	tx transaction.TxManager,
	repoFactory *repofactory.TxRepoFactory,
	accounts *accountQuery.Query,
	emails *emailQuery.Query,
	phones *phoneQuery.Query,
	profiles *profileQuery.Query,
	tokens *tokenx.Tokenx,
	redisClient *redis.Client,
	mail *email.EmailService,
	bus *eventbus.BusPublisher,
) *Service {
	s := &Service{
		tx: tx, repoFactory: repoFactory, accounts: accounts, emails: emails, phones: phones, profiles: profiles,
		tokens: tokens, redis: redisClient, mail: mail, bus: bus,
	}
	if redisClient != nil {
		s.checkFn = s.redisCheckCode
	}
	return s
}
