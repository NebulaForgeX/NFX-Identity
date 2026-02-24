package badges

import (
	dirErr "nfxid/errors/src/directory"
)

func (b *Badge) Validate() error {
	if b.Name() == "" {
		return dirErr.ErrNameRequired
	}
	return nil
}
