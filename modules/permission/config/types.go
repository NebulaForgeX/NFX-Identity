package config

import (
	"nfxid/pkgs/cache"
	"nfxid/pkgs/env"
	"nfxid/pkgs/kafkax"
	"nfxid/pkgs/logx"
	"nfxid/pkgs/mongodbx"
	"nfxid/pkgs/postgresqlx"
	"nfxid/pkgs/tokenx"
	"time"
)

type Config struct {
	Env         env.Env
	Server      ServerConfig       `koanf:"server"`
	PostgreSQL  postgresqlx.Config `koanf:"postgresql"`
	Mongo       mongodbx.Config    `koanf:"mongodb"`
	Cache       cache.ConnConfig   `koanf:"cache"`
	Logger      logx.LoggerConfig  `koanf:"logger"`
	Token       TokenConfig        `koanf:"token"`
	KafkaConfig kafkax.Config      `koanf:"kafka"`
	GRPCClient  GRPCClientConfig   `koanf:"grpc_client"`
}

type GRPCClientConfig struct {
	AuthAddr string `koanf:"auth_addr"` // 例如: "localhost:10012" 或 "auth:50051"
}

// TokenConfig Token 配置（TOML 中使用字符串）
type TokenConfig struct {
	SecretKey       string `koanf:"secret_key"`
	Issuer          string `koanf:"issuer"`
	AccessTokenTTL  string `koanf:"access_token_ttl"`  // 例如: "15m"
	RefreshTokenTTL string `koanf:"refresh_token_ttl"` // 例如: "168h"
	Algorithm       string `koanf:"algorithm"`
}

// ToTokenxConfig 转换为 tokenx.Config
func (tc TokenConfig) ToTokenxConfig() (tokenx.Config, error) {
	accessTTL, err := time.ParseDuration(tc.AccessTokenTTL)
	if err != nil {
		return tokenx.Config{}, err
	}

	refreshTTL, err := time.ParseDuration(tc.RefreshTokenTTL)
	if err != nil {
		return tokenx.Config{}, err
	}

	algorithm := tc.Algorithm
	if algorithm == "" {
		algorithm = "HS256"
	}

	return tokenx.Config{
		SecretKey:       tc.SecretKey,
		Issuer:          tc.Issuer,
		AccessTokenTTL:  accessTTL,
		RefreshTokenTTL: refreshTTL,
		Algorithm:       algorithm,
	}, nil
}

type ServerConfig struct {
	Name     string `koanf:"name"`
	Host     string `koanf:"host"`
	HTTPPort int    `koanf:"http_port"`
	GRPCPort int    `koanf:"grpc_port"`
}

