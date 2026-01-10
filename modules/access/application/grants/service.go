package grants

import (
	"context"
	grantAppResult "nfxid/modules/access/application/grants/results"
	grantDomain "nfxid/modules/access/domain/grants"

	"github.com/google/uuid"
)

type Service struct {
	grantRepo *grantDomain.Repo
}

func NewService(
	grantRepo *grantDomain.Repo,
) *Service {
	return &Service{
		grantRepo: grantRepo,
	}
}

// GetGrantsBySubject 根据主体获取授权列表
func (s *Service) GetGrantsBySubject(ctx context.Context, subjectType grantDomain.SubjectType, subjectID uuid.UUID) ([]grantAppResult.GrantRO, error) {
	grants, err := s.grantRepo.Get.BySubject(ctx, subjectType, subjectID)
	if err != nil {
		return nil, err
	}
	
	results := make([]grantAppResult.GrantRO, len(grants))
	for i, g := range grants {
		results[i] = grantAppResult.GrantMapper(g)
	}
	return results, nil
}
