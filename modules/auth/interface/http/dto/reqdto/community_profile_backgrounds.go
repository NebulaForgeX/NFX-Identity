package reqdto

import (
	"strings"

	authErr "nfxidentity/errors/src/auth"
	account "nfxidentity/modules/auth/application/account"

	"github.com/google/uuid"
)

type ConfirmCommunityProfileBackgroundItem struct {
	ImageID   string `json:"image_id"`
	SortOrder int    `json:"sort_order"`
}

type ConfirmCommunityProfileBackgrounds struct {
	Images []ConfirmCommunityProfileBackgroundItem `json:"images"`
}

func (r ConfirmCommunityProfileBackgrounds) Validate() error {
	for _, image := range r.Images {
		imageID, err := uuid.Parse(strings.TrimSpace(image.ImageID))
		if err != nil || imageID == uuid.Nil {
			return authErr.ErrForgerProfileBackgroundImageIDInvalid.WithCause(err)
		}
		if image.SortOrder < 0 {
			return authErr.ErrForgerProfileBackgroundSortOrderInvalid
		}
	}
	return nil
}

func (r ConfirmCommunityProfileBackgrounds) ToInput(accountID, profileID uuid.UUID) (account.ConfirmCommunityProfileBackgroundsInput, error) {
	images := make([]account.ConfirmCommunityProfileBackgroundItemInput, 0, len(r.Images))
	for _, image := range r.Images {
		imageID, err := uuid.Parse(strings.TrimSpace(image.ImageID))
		if err != nil || imageID == uuid.Nil {
			return account.ConfirmCommunityProfileBackgroundsInput{}, authErr.ErrForgerProfileBackgroundImageIDInvalid.WithCause(err)
		}
		images = append(images, account.ConfirmCommunityProfileBackgroundItemInput{
			ImageID:   imageID,
			SortOrder: image.SortOrder,
		})
	}
	return account.ConfirmCommunityProfileBackgroundsInput{
		AccountID: accountID,
		ProfileID: profileID,
		Images:    images,
	}, nil
}
