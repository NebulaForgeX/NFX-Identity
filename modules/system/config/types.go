package config

import (
	"nfxidentity/pkgs/cachex"
	"nfxidentity/pkgs/env"
	"nfxidentity/pkgs/httpx"
	"nfxidentity/pkgs/kafkax"
	"nfxidentity/pkgs/logx"
	"nfxidentity/pkgs/postgresqlx"
	"nfxidentity/pkgs/rabbitmqx"
	"nfxidentity/pkgs/tokenx"
)

type Config struct {
	Env            env.Env
	Server         ServerConfig       `koanf:"server"`
	PostgreSQL     postgresqlx.Config `koanf:"postgresql"`
	Cache          cachex.ConnConfig  `koanf:"cache"`
	Logger         logx.LoggerConfig  `koanf:"logger"`
	KafkaConfig    kafkax.Config      `koanf:"kafka"`
	RabbitMQConfig rabbitmqx.Config   `koanf:"rabbitmq"`
	GRPCClient     GRPCClientConfig   `koanf:"grpc_client"`
	Token          tokenx.Config      `koanf:"token"`
	I18n           I18nConfig         `koanf:"i18n"`
}

// I18nConfig 错误码翻译 JSON 目录（挂载路径，外部更新即生效）
type I18nConfig struct {
	ErrorsLangsPath string `koanf:"errors_langs_path"` // 如 ./data/errors/langs，与 data 一样可挂载
}

type GRPCClientConfig struct {
	AuthAddr      string `koanf:"auth_addr"`      // auth service address, e.g., "localhost:10001" or "auth:50051"
	AccessAddr    string `koanf:"access_addr"`    // access service address, e.g., "localhost:10000" or "access:50051"
	AuditAddr     string `koanf:"audit_addr"`     // audit service address, e.g., "localhost:10002" or "audit:50051"
	ClientsAddr   string `koanf:"clients_addr"`   // clients service address, e.g., "localhost:10003" or "clients:50051"
	DirectoryAddr string `koanf:"directory_addr"` // directory service address, e.g., "localhost:10004" or "directory:50051"
	ImageAddr     string `koanf:"image_addr"`     // image service address, e.g., "localhost:10005" or "image:50051"
	TenantsAddr   string `koanf:"tenants_addr"`   // tenants service address, e.g., "localhost:10007" or "tenants:50051"
}

type ServerConfig struct {
	Name      string                `koanf:"name"`
	Host      string                `koanf:"host"`
	HTTPPort  int                   `koanf:"http_port"`
	GRPCPort  int                   `koanf:"grpc_port"`
	AccessLog httpx.AccessLogConfig `koanf:"access_log"`
}
