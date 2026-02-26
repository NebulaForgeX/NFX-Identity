package user_emails

import (
	userEmailDomain "nfxidentity/modules/directory/domain/user_emails"
)

type Service struct {
	userEmailRepo *userEmailDomain.Repo
}

func NewService(
	userEmailRepo *userEmailDomain.Repo,
) *Service {
	return &Service{
		userEmailRepo: userEmailRepo,
	}
}
