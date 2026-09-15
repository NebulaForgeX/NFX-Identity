package resource

import (
	"context"
)

func (s *Service) CheckRedis(ctx context.Context) error {
	if s.cache == nil {
		return context.Canceled
	}
	return s.cache.Check(ctx)
}
