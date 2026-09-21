package http

import (
	"encoding/json"
	"time"

	accountSvc "nfxidentity/modules/auth/application/account"
	emailSvc "nfxidentity/modules/auth/application/email"
	loginSvc "nfxidentity/modules/auth/application/login"
	phoneSvc "nfxidentity/modules/auth/application/phone"
	signupSvc "nfxidentity/modules/auth/application/signup"
	"nfxidentity/modules/auth/interface/http/handler"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/fiberx/middleware"
	"nfxidentity/pkgs/httpx"
	"nfxidentity/pkgs/security/token"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

type Registry struct {
	Login            *handler.LoginHandler
	Signup           *handler.SignupHandler
	Account          *handler.AccountHandler
	AuthorityAccount *handler.AuthorityAccountHandler
	Owner            *handler.OwnerHandler
	Locales          *handler.LocalesHandler
}

type httpDeps interface {
	LoginService() *loginSvc.Service
	SignupService() *signupSvc.Service
	AccountService() *accountSvc.Service
	EmailService() *emailSvc.Service
	PhoneService() *phoneSvc.Service
	UserTokenVerifier() token.Verifier
}

func NewHTTPServer(d httpDeps, accessLog httpx.AccessLogConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: fiberx.ErrorHandler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "X-Api-Key", "X-Request-ID"},
		AllowCredentials: false,
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		MaxAge:           3600,
	}))
	app.Use(middleware.Logger(), middleware.AccessLog(accessLog), middleware.Recover())

	reg := &Registry{
		Login:            handler.NewLoginHandler(d.LoginService()),
		Signup:           handler.NewSignupHandler(d.SignupService()),
		Account:          handler.NewAccountHandler(d.AccountService(), d.EmailService(), d.PhoneService()),
		AuthorityAccount: handler.NewAuthorityAccountHandler(d.AccountService()),
		Owner:            handler.NewOwnerHandler(d.AccountService()),
		Locales:          handler.NewLocalesHandler(),
	}
	NewRouter(app, reg, d.UserTokenVerifier()).RegisterRoutes()
	return app
}
