package resource

import (
	"nfxid/pkgs/cachex"
	"nfxid/pkgs/kafkax"
	"nfxid/pkgs/postgresqlx"
	"nfxid/pkgs/rabbitmqx"
)

// Service 资源健康检查服务
type Service struct {
	postgres   *postgresqlx.Connection
	cache      *cachex.Connection
	kafkaCfg   *kafkax.Config
	rabbitMQCfg *rabbitmqx.Config
}

// NewService 创建资源健康检查服务
func NewService(
	postgres *postgresqlx.Connection,
	cache *cachex.Connection,
	kafkaCfg *kafkax.Config,
	rabbitMQCfg *rabbitmqx.Config,
) *Service {
	return &Service{
		postgres:   postgres,
		cache:      cache,
		kafkaCfg:   kafkaCfg,
		rabbitMQCfg: rabbitMQCfg,
	}
}
