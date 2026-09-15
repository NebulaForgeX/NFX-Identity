package resource

import (
	"context"
)

func (s *Service) CheckKafka(ctx context.Context) error {
	if s.kafkaCfg == nil {
		return context.Canceled
	}
	return nil
}
