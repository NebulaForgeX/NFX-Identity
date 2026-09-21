package email

import (
	"context"

	emailQuery "nfxidentity/modules/auth/query/email"
	"nfxidentity/pkgs/httpx"

	"github.com/google/uuid"
)

type ListInput struct {
	AccountID uuid.UUID
}

func (s *Service) List(ctx context.Context, in ListInput) (httpx.Page[emailQuery.EmailItemVO], error) {
	items, err := s.emails.List.ByAccountID(ctx, in.AccountID)
	if err != nil {
		return httpx.Page[emailQuery.EmailItemVO]{}, err
	}
	return httpx.NewPage(items, int64(len(items))), nil
}
