package resource

import (
	"context"
	"errors"
)

// CheckMySQL 检查 MySQL 数据库连接的健康状态
func (s *Service) CheckMySQL(ctx context.Context) error {
	// TODO: 实现 MySQL 健康检查逻辑
	// 例如：执行简单的 SQL 查询（如 SELECT 1）来验证 MySQL 连接是否正常
	return nil
}

// CheckPostgres 检查 PostgreSQL 数据库连接的健康状态
func (s *Service) CheckPostgres(ctx context.Context) error {
	if s.postgres == nil {
		return errors.New("postgres connection not initialized")
	}
	return s.postgres.Check(ctx)
}

// CheckDynamoDB 检查 DynamoDB 连接的健康状态
func (s *Service) CheckDynamoDB(ctx context.Context) error {
	// TODO: 实现 DynamoDB 健康检查逻辑
	// 例如：执行简单的查询来验证 DynamoDB 连接是否正常
	return nil
}
