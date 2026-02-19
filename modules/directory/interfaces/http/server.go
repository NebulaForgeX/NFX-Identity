package http

import (
	"encoding/json"
	"time"

	badgeApp "nfxid/modules/directory/application/badges"
	userAvatarApp "nfxid/modules/directory/application/user_avatars"
	userBadgeApp "nfxid/modules/directory/application/user_badges"
	userEducationApp "nfxid/modules/directory/application/user_educations"
	userEmailApp "nfxid/modules/directory/application/user_emails"
	userImageApp "nfxid/modules/directory/application/user_images"
	userOccupationApp "nfxid/modules/directory/application/user_occupations"
	userPhoneApp "nfxid/modules/directory/application/user_phones"
	userPreferenceApp "nfxid/modules/directory/application/user_preferences"
	userProfileApp "nfxid/modules/directory/application/user_profiles"
	userApp "nfxid/modules/directory/application/users"
	"nfxid/modules/directory/interfaces/http/handler"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/fiberx/middleware"
	"nfxid/pkgs/security/token"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
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
	UserAvatarAppSvc() *userAvatarApp.Service
	UserImageAppSvc() *userImageApp.Service
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
		UserAvatar:     handler.NewUserAvatarHandler(d.UserAvatarAppSvc()),
		UserImage:      handler.NewUserImageHandler(d.UserImageAppSvc()),
	}

	// 注册路由
	router := NewRouter(app, d.UserTokenVerifier(), reg)
	router.RegisterRoutes()

	return app
}
