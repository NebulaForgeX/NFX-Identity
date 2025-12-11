package permission

import (
	permissionErrors "nfxid/modules/permission/domain/permission/errors"
	"strings"
)

func (e *PermissionEditable) Validate() error {
	if strings.TrimSpace(e.Tag) == "" {
		return permissionErrors.ErrPermissionTagRequired
	}
	if len(e.Tag) < 2 {
		return permissionErrors.ErrPermissionTagInvalid
	}
	if len(e.Tag) > 100 {
		return permissionErrors.ErrPermissionTagInvalid
	}

	if strings.TrimSpace(e.Name) == "" {
		return permissionErrors.ErrPermissionNameRequired
	}
	if len(e.Name) > 255 {
		return permissionErrors.ErrPermissionNameInvalid
	}

	return nil
}

