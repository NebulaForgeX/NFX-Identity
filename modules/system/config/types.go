package config

import (
	"nfxidentity/pkgs/cachex"
	"nfxidentity/pkgs/connections/otelx"
	"nfxidentity/pkgs/env"
	"nfxidentity/pkgs/httpx"
	"nfxidentity/pkgs/kafkax"
	"nfxidentity/pkgs/logx"
	"nfxidentity/pkgs/postgresqlx"
	"nfxidentity/pkgs/tokenx"
)

type Config struct {
	Env            env.Env
	Server         ServerConfig       `koanf:"server"`
	PostgreSQL     postgresqlx.Config `koanf:"postgresql"`
	Cache          cachex.ConnConfig  `koanf:"cache"`
	Logger         logx.LoggerConfig  `koanf:"logger"`
	KafkaConfig kafkax.Config      `koanf:"kafka"`
	GRPCClient  GRPCClientConfig  `koanf:"grpc_client"`
	Token       tokenx.Config     `koanf:"token"`
	I18n        I18nConfig        `koanf:"i18n"`
	OTEL        otelx.Config      `koanf:"otel"`
}

// I18nConfig 错误码翻译 JSON 目录（挂载路径，外部更新即生效）
type I18nConfig struct {
	ErrorsLangsPath string `koanf:"errors_langs_path"` // 如 ./data/errors/langs，与 data 一样可挂载
}

type GRPCClientConfig struct {
	AuthAddr  string `koanf:"auth_addr"`
	AssetAddr string `koanf:"asset_addr"`
}

type ServerConfig struct {
	Name      string                `koanf:"name"`
	Host      string                `koanf:"host"`
	HTTPPort  int                   `koanf:"http_port"`
	GRPCPort  int                   `koanf:"grpc_port"`
	AccessLog httpx.AccessLogConfig `koanf:"access_log"`
}
