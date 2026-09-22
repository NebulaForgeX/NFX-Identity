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
	MinIO       MinIOConfig        `koanf:"minio"`
	OTEL        otelx.Config       `koanf:"otel"`
	GRPCClient  GRPCClientConfig   `koanf:"grpc_client"`
}

type GRPCClientConfig struct {
	AuthAddr string `koanf:"auth_addr"`
}

type ServerConfig struct {
	Name      string                `koanf:"name"`
	Host      string                `koanf:"host"`
	HTTPPort  int                   `koanf:"http_port"`
	GRPCPort  int                   `koanf:"grpc_port"`
	AccessLog httpx.AccessLogConfig `koanf:"access_log"`
}

type MinIOConfig struct {
	Endpoint  string `koanf:"endpoint"`
	PublicURL string `koanf:"public_url"`
	LanURL    string `koanf:"lan_url"`
	AccessKey string `koanf:"access_key"`
	SecretKey string `koanf:"secret_key"`
	Bucket    string `koanf:"bucket"`
	UseSSL    bool   `koanf:"use_ssl"`
	Region    string `koanf:"region"`
}
