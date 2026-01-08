package config

import (
	"nfxid/pkgs/cache"
	"nfxid/pkgs/env"
	"nfxid/pkgs/kafkax"
	"nfxid/pkgs/logx"
	"nfxid/pkgs/mongodbx"
	"nfxid/pkgs/postgresqlx"
	"nfxid/pkgs/tokenx"
)

type Config struct {
	Env         env.Env
	Server      ServerConfig       `koanf:"server"`
	PostgreSQL  postgresqlx.Config `koanf:"postgresql"`
	Mongo       mongodbx.Config    `koanf:"mongodb"`
	Cache       cache.ConnConfig   `koanf:"cache"`
	Logger      logx.LoggerConfig  `koanf:"logger"`
	Token       tokenx.Config      `koanf:"token"`
	Storage     StorageConfig      `koanf:"storage"`
	KafkaConfig kafkax.Config      `koanf:"kafka"`
	Email       EmailConfig        `koanf:"email"`
	GRPCClient  GRPCClientConfig   `koanf:"grpc_client"`
}

type GRPCClientConfig struct {
	ImageAddr string `koanf:"image_addr"` // 例如: "localhost:10013" 或 "image:50051"
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

type EmailConfig struct {
	SMTPHost     string `koanf:"smtp_host"`
	SMTPPort     int    `koanf:"smtp_port"`
	SMTPUser     string `koanf:"smtp_user"`
	SMTPPassword string `koanf:"smtp_password"`
	SMTPFrom     string `koanf:"smtp_from"`
}
