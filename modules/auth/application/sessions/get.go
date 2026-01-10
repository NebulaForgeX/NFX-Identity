package sessions

import (
	"context"
	sessionResult "nfxid/modules/auth/application/sessions/results"

	"github.com/google/uuid"
)

// GetSession 根据ID获取会话
func (s *Service) GetSession(ctx context.Context, sessionID uuid.UUID) (sessionResult.SessionRO, error) {
	domainEntity, err := s.sessionRepo.Get.ByID(ctx, sessionID)
	if err != nil {
		return sessionResult.SessionRO{}, err
	}
	return sessionResult.SessionMapper(domainEntity), nil
}

// GetSessionBySessionID 根据SessionID获取会话
func (s *Service) GetSessionBySessionID(ctx context.Context, sessionID string) (sessionResult.SessionRO, error) {
	domainEntity, err := s.sessionRepo.Get.BySessionID(ctx, sessionID)
	if err != nil {
		return sessionResult.SessionRO{}, err
	}
	return sessionResult.SessionMapper(domainEntity), nil
}
