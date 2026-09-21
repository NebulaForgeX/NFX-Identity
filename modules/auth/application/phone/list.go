package phone

import (
	"context"

	phoneQuery "nfxidentity/modules/auth/query/phone"
	"nfxidentity/pkgs/httpx"

	"github.com/google/uuid"
)

type ListInput struct {
	AccountID uuid.UUID
}

func (s *Service) List(ctx context.Context, in ListInput) (httpx.Page[phoneQuery.PhoneItemVO], error) {
	items, err := s.phones.List.ByAccountID(ctx, in.AccountID)
	if err != nil {
		return httpx.Page[phoneQuery.PhoneItemVO]{}, err
	}
	return httpx.NewPage(items, int64(len(items))), nil
}
