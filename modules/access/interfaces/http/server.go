package http

import (
	"encoding/json"

	grantApp "nfxid/modules/access/application/grants"
	permissionApp "nfxid/modules/access/application/permissions"
	roleApp "nfxid/modules/access/application/roles"
	rolePermissionApp "nfxid/modules/access/application/role_permissions"
	scopeApp "nfxid/modules/access/application/scopes"
	scopePermissionApp "nfxid/modules/access/application/scope_permissions"
	"nfxid/modules/access/interfaces/http/handler"
	"nfxid/pkgs/recover"
	"nfxid/pkgs/security/token"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type httpDeps interface {
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
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// CORS 中间件 - 必须在其他中间件之前
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*", // 开发环境允许所有源
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: false, // 使用通配符时不能为 true，JWT token 通过 Authorization header 传递
		ExposeHeaders:    "Content-Length",
	}))

	app.Use(recover.RecoverMiddleware(), logger.New())

	// 创建handlers
	reg := &Registry{
		Role:            handler.NewRoleHandler(d.RoleAppSvc()),
		Permission:      handler.NewPermissionHandler(d.PermissionAppSvc()),
		Scope:           handler.NewScopeHandler(d.ScopeAppSvc()),
		Grant:           handler.NewGrantHandler(d.GrantAppSvc()),
		RolePermission:  handler.NewRolePermissionHandler(d.RolePermissionAppSvc()),
		ScopePermission: handler.NewScopePermissionHandler(d.ScopePermissionAppSvc()),
	}

	// 注册路由
	router := NewRouter(app, d.UserTokenVerifier(), reg)
	router.RegisterRoutes()

	return app
}
