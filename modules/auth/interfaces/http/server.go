package http

import (
	badgeApp "nfxid/modules/auth/application/badge"
	educationApp "nfxid/modules/auth/application/education"
	occupationApp "nfxid/modules/auth/application/occupation"
	profileApp "nfxid/modules/auth/application/profile"
	profileBadgeApp "nfxid/modules/auth/application/profile_badge"
	roleApp "nfxid/modules/auth/application/role"
	userApp "nfxid/modules/auth/application/user"
	"nfxid/modules/auth/interfaces/http/handler"
	"nfxid/pkgs/recover"
	"nfxid/pkgs/tokenx"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type httpDeps interface {
	UserAppSvc() *userApp.Service
	ProfileAppSvc() *profileApp.Service
	RoleAppSvc() *roleApp.Service
	BadgeAppSvc() *badgeApp.Service
	EducationAppSvc() *educationApp.Service
	OccupationAppSvc() *occupationApp.Service
	ProfileBadgeAppSvc() *profileBadgeApp.Service
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
		User:         handler.NewUserHandler(d.UserAppSvc()),
		Profile:      handler.NewProfileHandler(d.ProfileAppSvc()),
		Role:         handler.NewRoleHandler(d.RoleAppSvc()),
		Badge:        handler.NewBadgeHandler(d.BadgeAppSvc()),
		Education:    handler.NewEducationHandler(d.EducationAppSvc()),
		Occupation:   handler.NewOccupationHandler(d.OccupationAppSvc()),
		ProfileBadge: handler.NewProfileBadgeHandler(d.ProfileBadgeAppSvc()),
	}

	// 注册路由
	router := NewRouter(app, d.Tokenx(), reg)
	router.RegisterRoutes()

	return app
}
