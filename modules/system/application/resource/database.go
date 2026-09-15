package resource

import (
	"context"
)

func (s *Service) CheckPostgres(ctx context.Context) error {
	if s.postgres == nil {
		return context.Canceled
	}
	return s.postgres.Check(ctx)
}
