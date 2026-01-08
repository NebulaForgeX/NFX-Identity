package images

import (
	"time"
)

func (i *Image) UpdateURL(url string) error {
	if i.DeletedAt() != nil {
		return ErrImageNotFound
	}
	i.state.URL = &url
	i.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (i *Image) UpdatePublic(isPublic bool) error {
	if i.DeletedAt() != nil {
		return ErrImageNotFound
	}
	i.state.IsPublic = isPublic
	i.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (i *Image) UpdateMetadata(metadata map[string]interface{}) error {
	if i.DeletedAt() != nil {
		return ErrImageNotFound
	}
	if metadata == nil {
		return nil
	}
	i.state.Metadata = metadata
	i.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (i *Image) Delete() error {
	if i.DeletedAt() != nil {
		return nil // idempotent
	}

	now := time.Now().UTC()
	i.state.DeletedAt = &now
	i.state.UpdatedAt = now
	return nil
}
