package http

import (
	"encoding/json"

	accountLockoutApp "nfxid/modules/auth/application/account_lockouts"
	loginAttemptApp "nfxid/modules/auth/application/login_attempts"
	mfaFactorApp "nfxid/modules/auth/application/mfa_factors"
	passwordHistoryApp "nfxid/modules/auth/application/password_history"
	passwordResetApp "nfxid/modules/auth/application/password_resets"
	refreshTokenApp "nfxid/modules/auth/application/refresh_tokens"
	sessionApp "nfxid/modules/auth/application/sessions"
	trustedDeviceApp "nfxid/modules/auth/application/trusted_devices"
	userCredentialApp "nfxid/modules/auth/application/user_credentials"
	"nfxid/modules/auth/interfaces/http/handler"
	"nfxid/pkgs/recover"
	"nfxid/pkgs/security/token"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type httpDeps interface {
	SessionAppSvc() *sessionApp.Service
	UserCredentialAppSvc() *userCredentialApp.Service
	MFAFactorAppSvc() *mfaFactorApp.Service
	RefreshTokenAppSvc() *refreshTokenApp.Service
	PasswordResetAppSvc() *passwordResetApp.Service
	PasswordHistoryAppSvc() *passwordHistoryApp.Service
	LoginAttemptAppSvc() *loginAttemptApp.Service
	AccountLockoutAppSvc() *accountLockoutApp.Service
	TrustedDeviceAppSvc() *trustedDeviceApp.Service
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
		Session:         handler.NewSessionHandler(d.SessionAppSvc()),
		UserCredential:  handler.NewUserCredentialHandler(d.UserCredentialAppSvc()),
		MFAFactor:       handler.NewMFAFactorHandler(d.MFAFactorAppSvc()),
		RefreshToken:    handler.NewRefreshTokenHandler(d.RefreshTokenAppSvc()),
		PasswordReset:   handler.NewPasswordResetHandler(d.PasswordResetAppSvc()),
		PasswordHistory: handler.NewPasswordHistoryHandler(d.PasswordHistoryAppSvc()),
		LoginAttempt:    handler.NewLoginAttemptHandler(d.LoginAttemptAppSvc()),
		AccountLockout:  handler.NewAccountLockoutHandler(d.AccountLockoutAppSvc()),
		TrustedDevice:   handler.NewTrustedDeviceHandler(d.TrustedDeviceAppSvc()),
	}

	// 注册路由
	router := NewRouter(app, d.UserTokenVerifier(), reg)
	router.RegisterRoutes()

	return app
}
