package client_scopes

import "github.com/google/uuid"

func (cs *ClientScope) Validate() error {
	if cs.AppID() == uuid.Nil {
		return ErrAppIDRequired
	}
	if cs.Scope() == "" {
		return ErrScopeRequired
	}
	return nil
}
