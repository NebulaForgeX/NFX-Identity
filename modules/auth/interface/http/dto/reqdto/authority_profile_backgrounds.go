package reqdto

import (
	"strings"

	authErr "nfxidentity/errors/src/auth"
	account "nfxidentity/modules/auth/application/account"

	"github.com/google/uuid"
)

type ConfirmAuthorityProfileBackgroundItem struct {
	ImageID   string `json:"image_id"`
	SortOrder int    `json:"sort_order"`
}

type ConfirmAuthorityProfileBackgrounds struct {
	Images []ConfirmAuthorityProfileBackgroundItem `json:"images"`
}

func (r ConfirmAuthorityProfileBackgrounds) Validate() error {
	for _, image := range r.Images {
		imageID, err := uuid.Parse(strings.TrimSpace(image.ImageID))
		if err != nil || imageID == uuid.Nil {
			return authErr.ErrAuthorityProfileBackgroundImageIDInvalid.WithCause(err)
		}
		if image.SortOrder < 0 {
			return authErr.ErrAuthorityProfileBackgroundSortOrderInvalid
		}
	}
	return nil
}

func (r ConfirmAuthorityProfileBackgrounds) ToInput(accountID, profileID uuid.UUID) (account.ConfirmAuthorityProfileBackgroundsInput, error) {
	images := make([]account.ConfirmAuthorityProfileBackgroundItemInput, 0, len(r.Images))
	for _, image := range r.Images {
		imageID, err := uuid.Parse(strings.TrimSpace(image.ImageID))
		if err != nil || imageID == uuid.Nil {
			return account.ConfirmAuthorityProfileBackgroundsInput{}, authErr.ErrAuthorityProfileBackgroundImageIDInvalid.WithCause(err)
		}
		images = append(images, account.ConfirmAuthorityProfileBackgroundItemInput{
			ImageID:   imageID,
			SortOrder: image.SortOrder,
		})
	}
	return account.ConfirmAuthorityProfileBackgroundsInput{
		AccountID: accountID,
		ProfileID: profileID,
		Images:    images,
	}, nil
}
