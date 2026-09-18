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
	Env         env.Env
	Server      ServerConfig       `koanf:"server"`
	PostgreSQL  postgresqlx.Config `koanf:"postgresql"`
	Cache       cachex.ConnConfig  `koanf:"cache"`
	Logger      logx.LoggerConfig  `koanf:"logger"`
	Token       tokenx.Config      `koanf:"token"`
	KafkaConfig kafkax.Config      `koanf:"kafka"`
	Email       EmailConfig        `koanf:"email"`
	OTEL        otelx.Config       `koanf:"otel"`
	GRPCClient  GRPCClientConfig   `koanf:"grpc_client"`
}

type GRPCClientConfig struct {
	AssetAddr  string `koanf:"asset_addr"`
	SystemAddr string `koanf:"system_addr"`
}

type ServerConfig struct {
	Name      string                `koanf:"name"`
	Host      string                `koanf:"host"`
	HTTPPort  int                   `koanf:"http_port"`
	GRPCPort  int                   `koanf:"grpc_port"`
	AccessLog httpx.AccessLogConfig `koanf:"access_log"`
}

type EmailConfig struct {
	SMTPHost     string `koanf:"smtp_host"`
	SMTPPort     int    `koanf:"smtp_port"`
	SMTPUser     string `koanf:"smtp_user"`
	SMTPPassword string `koanf:"smtp_password"`
	SMTPFrom     string `koanf:"smtp_from"`
}
