package client_scopes

import (
	clientsErr "nfxidentity/errors/src/clients"

	"github.com/google/uuid"
)

func (cs *ClientScope) Validate() error {
	if cs.AppID() == uuid.Nil {
		return clientsErr.ErrAppIDRequired
	}
	if cs.Scope() == "" {
		return clientsErr.ErrScopeRequired
	}
	return nil
}
