package auth

import (
	accountlockoutpb "nfxid/protos/gen/auth/account_lockout"
	loginattemptpb "nfxid/protos/gen/auth/login_attempt"
	mfafactorpb "nfxid/protos/gen/auth/mfa_factor"
	passwordhistorypb "nfxid/protos/gen/auth/password_history"
	passwordresetpb "nfxid/protos/gen/auth/password_reset"
	refreshtokenpb "nfxid/protos/gen/auth/refresh_token"
	sessionpb "nfxid/protos/gen/auth/session"
	trusteddevicepb "nfxid/protos/gen/auth/trusted_device"
	usercredentialpb "nfxid/protos/gen/auth/user_credential"
)

// Client Auth 服务客户端
type Client struct {
	UserCredential  *UserCredentialClient
	Session         *SessionClient
	TrustedDevice   *TrustedDeviceClient
	MfaFactor       *MfaFactorClient
	RefreshToken    *RefreshTokenClient
	PasswordReset   *PasswordResetClient
	PasswordHistory *PasswordHistoryClient
	LoginAttempt    *LoginAttemptClient
	AccountLockout  *AccountLockoutClient
}

// NewClient 创建 Auth 客户端
func NewClient(
	userCredentialClient usercredentialpb.UserCredentialServiceClient,
	sessionClient sessionpb.SessionServiceClient,
	trustedDeviceClient trusteddevicepb.TrustedDeviceServiceClient,
	mfaFactorClient mfafactorpb.MfaFactorServiceClient,
	refreshTokenClient refreshtokenpb.RefreshTokenServiceClient,
	passwordResetClient passwordresetpb.PasswordResetServiceClient,
	passwordHistoryClient passwordhistorypb.PasswordHistoryServiceClient,
	loginAttemptClient loginattemptpb.LoginAttemptServiceClient,
	accountLockoutClient accountlockoutpb.AccountLockoutServiceClient,
) *Client {
	return &Client{
		UserCredential:  NewUserCredentialClient(userCredentialClient),
		Session:         NewSessionClient(sessionClient),
		TrustedDevice:   NewTrustedDeviceClient(trustedDeviceClient),
		MfaFactor:       NewMfaFactorClient(mfaFactorClient),
		RefreshToken:    NewRefreshTokenClient(refreshTokenClient),
		PasswordReset:   NewPasswordResetClient(passwordResetClient),
		PasswordHistory: NewPasswordHistoryClient(passwordHistoryClient),
		LoginAttempt:    NewLoginAttemptClient(loginAttemptClient),
		AccountLockout:  NewAccountLockoutClient(accountLockoutClient),
	}
}
