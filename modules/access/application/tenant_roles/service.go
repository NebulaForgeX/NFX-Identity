package tenant_roles

import (
	domain "nfxidentity/modules/access/domain/tenant_roles"
)

type Service struct {
	repo *domain.Repo
}

func NewService(repo *domain.Repo) *Service {
	return &Service{repo: repo}
}
