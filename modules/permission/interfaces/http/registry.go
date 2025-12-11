package http

import (
	"nfxid/modules/permission/interfaces/http/handler"
)

type Registry struct {
	Auth           *handler.AuthHandler
	Permission     *handler.PermissionHandler
	UserPermission *handler.UserPermissionHandler
}

