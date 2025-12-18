package http

import (
	authApp "nfxid/modules/permission/application/auth"
	authorizationCodeApp "nfxid/modules/permission/application/authorization_code"
	permissionApp "nfxid/modules/permission/application/permission"
	userPermissionApp "nfxid/modules/permission/application/user_permission"
	"nfxid/modules/permission/interfaces/http/handler"
	"nfxid/pkgs/recover"
	"nfxid/pkgs/tokenx"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type httpDeps interface {
	AuthAppSvc() *authApp.Service
	PermissionAppSvc() *permissionApp.Service
	UserPermissionAppSvc() *userPermissionApp.Service
	AuthorizationCodeAppSvc() *authorizationCodeApp.Service
	Tokenx() *tokenx.Tokenx
}

func NewHTTPServer(d httpDeps) *fiber.App {
	app := fiber.New()

	// CORS 中间件
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: false,
		ExposeHeaders:    "Content-Length",
	}))

	app.Use(recover.RecoverMiddleware(), logger.New())

	// 创建 handlers
	reg := &Registry{
		Auth:              handler.NewAuthHandler(d.AuthAppSvc()),
		Permission:        handler.NewPermissionHandler(d.PermissionAppSvc()),
		UserPermission:    handler.NewUserPermissionHandler(d.UserPermissionAppSvc()),
		AuthorizationCode: handler.NewAuthorizationCodeHandler(d.AuthorizationCodeAppSvc()),
	}

	// 注册路由
	router := NewRouter(app, d.Tokenx(), reg)
	router.RegisterRoutes()

	return app
}
