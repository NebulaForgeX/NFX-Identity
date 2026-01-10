package http

import (
	"encoding/json"

	badgeApp "nfxid/modules/directory/application/badges"
	userApp "nfxid/modules/directory/application/users"
	userBadgeApp "nfxid/modules/directory/application/user_badges"
	userEducationApp "nfxid/modules/directory/application/user_educations"
	userEmailApp "nfxid/modules/directory/application/user_emails"
	userOccupationApp "nfxid/modules/directory/application/user_occupations"
	userPhoneApp "nfxid/modules/directory/application/user_phones"
	userPreferenceApp "nfxid/modules/directory/application/user_preferences"
	userProfileApp "nfxid/modules/directory/application/user_profiles"
	"nfxid/modules/directory/interfaces/http/handler"
	"nfxid/pkgs/recover"
	"nfxid/pkgs/security/token"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type httpDeps interface {
	UserAppSvc() *userApp.Service
	BadgeAppSvc() *badgeApp.Service
	UserBadgeAppSvc() *userBadgeApp.Service
	UserEducationAppSvc() *userEducationApp.Service
	UserEmailAppSvc() *userEmailApp.Service
	UserOccupationAppSvc() *userOccupationApp.Service
	UserPhoneAppSvc() *userPhoneApp.Service
	UserPreferenceAppSvc() *userPreferenceApp.Service
	UserProfileAppSvc() *userProfileApp.Service
	UserTokenVerifier() token.Verifier
}

func NewHTTPServer(d httpDeps) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// CORS 中间件 - 必须在其他中间件之前
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: false,
		ExposeHeaders:    "Content-Length",
	}))

	app.Use(recover.RecoverMiddleware(), logger.New())

	// 创建handlers
	reg := &Registry{
		User:           handler.NewUserHandler(d.UserAppSvc()),
		Badge:          handler.NewBadgeHandler(d.BadgeAppSvc()),
		UserBadge:      handler.NewUserBadgeHandler(d.UserBadgeAppSvc()),
		UserEducation:  handler.NewUserEducationHandler(d.UserEducationAppSvc()),
		UserEmail:      handler.NewUserEmailHandler(d.UserEmailAppSvc()),
		UserOccupation: handler.NewUserOccupationHandler(d.UserOccupationAppSvc()),
		UserPhone:      handler.NewUserPhoneHandler(d.UserPhoneAppSvc()),
		UserPreference: handler.NewUserPreferenceHandler(d.UserPreferenceAppSvc()),
		UserProfile:    handler.NewUserProfileHandler(d.UserProfileAppSvc()),
	}

	// 注册路由
	router := NewRouter(app, d.UserTokenVerifier(), reg)
	router.RegisterRoutes()

	return app
}
