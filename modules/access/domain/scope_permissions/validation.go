package scope_permissions

import "github.com/google/uuid"

func (sp *ScopePermission) Validate() error {
	if sp.Scope() == "" {
		return ErrScopeRequired
	}
	if sp.PermissionID() == uuid.Nil {
		return ErrPermissionIDRequired
	}
	return nil
}
