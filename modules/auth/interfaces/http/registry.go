package http

import (
	"nfxid/modules/auth/interfaces/http/handler"
)

type Registry struct {
	Session        *handler.SessionHandler
	UserCredential *handler.UserCredentialHandler
	MFAFactor      *handler.MFAFactorHandler
	RefreshToken   *handler.RefreshTokenHandler
	PasswordReset  *handler.PasswordResetHandler
	PasswordHistory *handler.PasswordHistoryHandler
	LoginAttempt   *handler.LoginAttemptHandler
	AccountLockout *handler.AccountLockoutHandler
	TrustedDevice  *handler.TrustedDeviceHandler
	Auth           *handler.AuthHandler
}
