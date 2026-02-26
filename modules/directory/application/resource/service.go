package resource

import (
	"nfxidentity/pkgs/cachex"
	"nfxidentity/pkgs/kafkax"
	"nfxidentity/pkgs/postgresqlx"
	"nfxidentity/pkgs/rabbitmqx"
)

// Service 资源健康检查服务
type Service struct {
	postgres    *postgresqlx.Connection
	cache       *cachex.Connection
	kafkaCfg    *kafkax.Config
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
		postgres:    postgres,
		cache:       cache,
		kafkaCfg:    kafkaCfg,
		rabbitMQCfg: rabbitMQCfg,
	}
}
