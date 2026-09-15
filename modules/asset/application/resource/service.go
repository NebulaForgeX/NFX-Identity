package resource

import (
	"nfxidentity/pkgs/cachex"
	"nfxidentity/pkgs/kafkax"
	"nfxidentity/pkgs/postgresqlx"
)

type Service struct {
	postgres *postgresqlx.Connection
	cache    *cachex.Connection
	kafkaCfg *kafkax.Config
}

func NewService(
	postgres *postgresqlx.Connection,
	cache *cachex.Connection,
	kafkaCfg *kafkax.Config,
) *Service {
	return &Service{
		postgres: postgres,
		cache:    cache,
		kafkaCfg: kafkaCfg,
	}
}
