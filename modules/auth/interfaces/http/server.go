package http

import (
	badgeApp "nebulaid/modules/auth/application/badge"
	educationApp "nebulaid/modules/auth/application/education"
	occupationApp "nebulaid/modules/auth/application/occupation"
	profileApp "nebulaid/modules/auth/application/profile"
	profileBadgeApp "nebulaid/modules/auth/application/profile_badge"
	roleApp "nebulaid/modules/auth/application/role"
	userApp "nebulaid/modules/auth/application/user"
	"nebulaid/modules/auth/interfaces/http/handler"
	"nebulaid/pkgs/recover"
	"nebulaid/pkgs/tokenx"

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
