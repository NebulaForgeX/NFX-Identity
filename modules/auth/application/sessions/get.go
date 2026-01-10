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

// GetSessionsByUserID 根据UserID获取会话列表
func (s *Service) GetSessionsByUserID(ctx context.Context, userID uuid.UUID) ([]sessionResult.SessionRO, error) {
	domainEntities, err := s.sessionRepo.Get.ByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	results := make([]sessionResult.SessionRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = sessionResult.SessionMapper(entity)
	}
	return results, nil
}
