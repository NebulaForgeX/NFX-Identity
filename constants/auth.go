package constants

import (
	"nfxidentity/enums"
	"nfxidentity/pkgs/constantx"
)

const (
	MaxLoginAttempts              = 5
	LockoutDurationMinutes        = 30
	DefaultAccessTokenTTLSeconds  = 900
	DefaultRefreshTokenTTLSeconds = 7 * 24 * 3600
	AuthProfileMaxBackgrounds     = 6
	AuthAccountMaxProfiles        = 5
	AuthAccountMinProfiles        = 1
	AuthAccountMinVerifiedEmails  = 1
)

var AuthProfileScope = constantx.NewStringEnumSet(
	enums.AuthProfileScopeCommunity,
	enums.AuthProfileScopeAuthority,
)

var AuthForgerRole = constantx.NewStringEnumSet(
	enums.AuthForgerRoleForger,
)

var AuthAuthorityRole = constantx.NewStringEnumSet(
	enums.AuthAuthorityRoleAuditor,
	enums.AuthAuthorityRoleAdministrator,
	enums.AuthAuthorityRoleOwner,
)

var AuthAccountStatus = constantx.NewStringEnumSet(
	enums.AuthAccountStatusActive,
	enums.AuthAccountStatusSuspended,
	enums.AuthAccountStatusDeleted,
)

var AuthSignupPlatform = constantx.NewStringEnumSet(
	enums.AuthSignupPlatformNfxidentity,
	enums.AuthSignupPlatformNfxnews,
	enums.AuthSignupPlatformNfxstorages,
	enums.AuthSignupPlatformNfxedge,
)

var AuthIdentityProvider = constantx.NewStringEnumSet(
	enums.AuthIdentityProviderPassword,
)

var AuthLanguage = constantx.NewStringEnumSet(
	enums.AuthProfileLanguageEn,
	enums.AuthProfileLanguageZh,
	enums.AuthProfileLanguageFr,
)
