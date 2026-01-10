package config

import (
	"nfxid/pkgs/cache"
	"nfxid/pkgs/env"
	"nfxid/pkgs/kafkax"
	"nfxid/pkgs/logx"
	"nfxid/pkgs/postgresqlx"
	"nfxid/pkgs/rabbitmqx"
)

type Config struct {
	Env          env.Env
	Server       ServerConfig       `koanf:"server"`
	PostgreSQL   postgresqlx.Config `koanf:"postgresql"`
	Cache        cache.ConnConfig   `koanf:"cache"`
	Logger       logx.LoggerConfig  `koanf:"logger"`
	Storage      StorageConfig      `koanf:"storage"`
	KafkaConfig  kafkax.Config      `koanf:"kafka"`
	RabbitMQConfig rabbitmqx.Config `koanf:"rabbitmq"`
	GRPCClient   GRPCClientConfig   `koanf:"grpc_client"`
}

type GRPCClientConfig struct {
	AccessAddr    string `koanf:"access_addr"`    // access service address, e.g., "localhost:10000" or "access:50051"
	AuthAddr      string `koanf:"auth_addr"`      // auth service address, e.g., "localhost:10001" or "auth:50051"
	AuditAddr     string `koanf:"audit_addr"`     // audit service address, e.g., "localhost:10002" or "audit:50051"
	ClientsAddr   string `koanf:"clients_addr"`   // clients service address, e.g., "localhost:10003" or "clients:50051"
	DirectoryAddr string `koanf:"directory_addr"` // directory service address, e.g., "localhost:10004" or "directory:50051"
	SystemAddr    string `koanf:"system_addr"`    // system service address, e.g., "localhost:10006" or "system:50051"
	TenantsAddr   string `koanf:"tenants_addr"`   // tenants service address, e.g., "localhost:10007" or "tenants:50051"
}

type ServerConfig struct {
	Name     string `koanf:"name"`
	Host     string `koanf:"host"`
	HTTPPort int    `koanf:"http_port"`
	GRPCPort int    `koanf:"grpc_port"`
}

type StorageConfig struct {
	BasePath string `koanf:"base_path"`
}
