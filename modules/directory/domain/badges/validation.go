package badges

import (
	dirErr "nfxidentity/errors/src/directory"
)

func (b *Badge) Validate() error {
	if b.Name() == "" {
		return dirErr.ErrNameRequired
	}
	return nil
}
