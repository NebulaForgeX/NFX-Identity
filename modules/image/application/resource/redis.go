package resource

import (
	"context"

	"nfxid/pkgs/errx"
)

// CheckRedis 检查 Redis 连接的健康状态
func (s *Service) CheckRedis(ctx context.Context) error {
	if s.cache == nil {
		return errx.FailedPrecond("REDIS_NOT_INITIALIZED", "redis connection not initialized")
	}
	return s.cache.Check(ctx)
}
