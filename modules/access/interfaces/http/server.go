package http

import (
	"encoding/json"
	"time"

	actionRequirementApp "nfxid/modules/access/application/action_requirements"
	actionApp "nfxid/modules/access/application/actions"
	grantApp "nfxid/modules/access/application/grants"
	permissionApp "nfxid/modules/access/application/permissions"
	rolePermissionApp "nfxid/modules/access/application/role_permissions"
	roleApp "nfxid/modules/access/application/roles"
	scopePermissionApp "nfxid/modules/access/application/scope_permissions"
	scopeApp "nfxid/modules/access/application/scopes"
	"nfxid/modules/access/interfaces/http/handler"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/fiberx/middleware"
	"nfxid/pkgs/security/token"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

type httpDeps interface {
	ActionAppSvc() *actionApp.Service
	ActionRequirementAppSvc() *actionRequirementApp.Service
	RoleAppSvc() *roleApp.Service
	PermissionAppSvc() *permissionApp.Service
	ScopeAppSvc() *scopeApp.Service
	GrantAppSvc() *grantApp.Service
	RolePermissionAppSvc() *rolePermissionApp.Service
	ScopePermissionAppSvc() *scopePermissionApp.Service
	UserTokenVerifier() token.Verifier
}

func NewHTTPServer(d httpDeps) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		ErrorHandler:  fiberx.ErrorHandler,
		ReadTimeout:   30 * time.Second,
		WriteTimeout:  30 * time.Second,
		IdleTimeout:   120 * time.Second,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "X-Api-Key", "X-Request-ID"},
		AllowCredentials: false,
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		MaxAge:           3600,
	}))

	app.Use(middleware.Logger(), middleware.Recover())

	reg := &Registry{
		Role:              handler.NewRoleHandler(d.RoleAppSvc()),
		Permission:        handler.NewPermissionHandler(d.PermissionAppSvc()),
		Scope:             handler.NewScopeHandler(d.ScopeAppSvc()),
		Grant:             handler.NewGrantHandler(d.GrantAppSvc()),
		RolePermission:    handler.NewRolePermissionHandler(d.RolePermissionAppSvc()),
		ScopePermission:   handler.NewScopePermissionHandler(d.ScopePermissionAppSvc()),
		Action:            handler.NewActionHandler(d.ActionAppSvc()),
		ActionRequirement: handler.NewActionRequirementHandler(d.ActionRequirementAppSvc()),
	}

	router := NewRouter(app, d.UserTokenVerifier(), reg)
	router.RegisterRoutes()

	return app
}
