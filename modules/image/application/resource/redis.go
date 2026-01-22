package resource

import (
	"context"
	"errors"
)

// CheckRedis 检查 Redis 连接的健康状态
func (s *Service) CheckRedis(ctx context.Context) error {
	if s.cache == nil {
		return errors.New("redis connection not initialized")
	}
	return s.cache.Check(ctx)
}
