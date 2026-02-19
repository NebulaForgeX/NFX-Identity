package http

import (
	"encoding/json"
	"time"

	accountLockoutApp "nfxid/modules/auth/application/account_lockouts"
	authApp "nfxid/modules/auth/application/auth"
	loginAttemptApp "nfxid/modules/auth/application/login_attempts"
	mfaFactorApp "nfxid/modules/auth/application/mfa_factors"
	passwordHistoryApp "nfxid/modules/auth/application/password_history"
	passwordResetApp "nfxid/modules/auth/application/password_resets"
	refreshTokenApp "nfxid/modules/auth/application/refresh_tokens"
	sessionApp "nfxid/modules/auth/application/sessions"
	trustedDeviceApp "nfxid/modules/auth/application/trusted_devices"
	userCredentialApp "nfxid/modules/auth/application/user_credentials"
	"nfxid/modules/auth/interfaces/http/handler"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/fiberx/middleware"
	"nfxid/pkgs/security/token"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
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
	AuthAppSvc() *authApp.Service
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
		Session:         handler.NewSessionHandler(d.SessionAppSvc()),
		UserCredential:  handler.NewUserCredentialHandler(d.UserCredentialAppSvc()),
		MFAFactor:       handler.NewMFAFactorHandler(d.MFAFactorAppSvc()),
		RefreshToken:    handler.NewRefreshTokenHandler(d.RefreshTokenAppSvc()),
		PasswordReset:   handler.NewPasswordResetHandler(d.PasswordResetAppSvc()),
		PasswordHistory: handler.NewPasswordHistoryHandler(d.PasswordHistoryAppSvc()),
		LoginAttempt:    handler.NewLoginAttemptHandler(d.LoginAttemptAppSvc()),
		AccountLockout:  handler.NewAccountLockoutHandler(d.AccountLockoutAppSvc()),
		TrustedDevice:   handler.NewTrustedDeviceHandler(d.TrustedDeviceAppSvc()),
		Auth:            handler.NewAuthHandler(d.AuthAppSvc()),
	}

	// 注册路由
	router := NewRouter(app, d.UserTokenVerifier(), reg)
	router.RegisterRoutes()

	return app
}
