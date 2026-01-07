package config

import (
	"nfxid/pkgs/cache"
	"nfxid/pkgs/env"
	"nfxid/pkgs/kafkax"
	"nfxid/pkgs/logx"
	"nfxid/pkgs/mongodbx"
	"nfxid/pkgs/postgresqlx"
)

type Config struct {
	Env         env.Env
	Server      ServerConfig       `koanf:"server"`
	PostgreSQL  postgresqlx.Config `koanf:"postgresql"`
	Mongo       mongodbx.Config    `koanf:"mongodb"`
	Cache       cache.ConnConfig   `koanf:"cache"`
	Logger      logx.LoggerConfig  `koanf:"logger"`
	Token       TokenConfig        `koanf:"token"`
	Storage     StorageConfig      `koanf:"storage"`
	KafkaConfig kafkax.Config      `koanf:"kafka"`
	GRPCClient  GRPCClientConfig   `koanf:"grpc_client"`
}

type GRPCClientConfig struct {
	AuthAddr string `koanf:"auth_addr"` // 例如: "localhost:10011" 或 "auth:10011"
}

type TokenConfig struct {
	SecretKey string `koanf:"secret_key"`
	Issuer    string `koanf:"issuer"`
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
