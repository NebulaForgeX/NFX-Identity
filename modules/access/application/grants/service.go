package grants

import (
	"context"
	"fmt"
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

// GetGrantsBySubjectString 根据主体获取授权列表（接受字符串类型的 subjectType）
// 这个方法用于 HTTP handler，避免 handler 直接依赖 domain 层
func (s *Service) GetGrantsBySubjectString(ctx context.Context, subjectTypeStr string, subjectID uuid.UUID) ([]grantAppResult.GrantRO, error) {
	// 在 application 层进行类型转换和验证
	subjectType := grantDomain.SubjectType(subjectTypeStr)
	if subjectType != grantDomain.SubjectTypeUser && subjectType != grantDomain.SubjectTypeClient {
		return nil, fmt.Errorf("invalid subject_type: %s, must be USER or CLIENT", subjectTypeStr)
	}
	
	return s.GetGrantsBySubject(ctx, subjectType, subjectID)
}
