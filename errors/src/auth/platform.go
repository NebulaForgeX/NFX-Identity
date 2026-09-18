package auth

import "nfxidentity/pkgs/errx"

var (
	ErrProfileNotOwned      = errx.Forbidden("PROFILE_NOT_OWNED", "profile is not owned by account")
	ErrInvalidProfileKind   = errx.InvalidArg("INVALID_PROFILE_KIND", "kind must be forger or authority")
	ErrInvalidProfileID     = errx.InvalidArg("INVALID_PROFILE_ID", "invalid profile id")
	ErrInvalidImageID       = errx.InvalidArg("INVALID_IMAGE_ID", "invalid image id")
	ErrInvalidEmailID       = errx.InvalidArg("INVALID_EMAIL_ID", "invalid email id")
	ErrInvalidPhoneID       = errx.InvalidArg("INVALID_PHONE_ID", "invalid phone id")
	ErrInvalidEmail         = errx.InvalidArg("INVALID_EMAIL", "email required")
	ErrInvalidPhone         = errx.InvalidArg("INVALID_PHONE", "phone required")
	ErrHashFailed           = errx.Internal("HASH_FAILED", "failed to hash password")
	ErrTokenFailed          = errx.Internal("TOKEN_FAILED", "token operation failed")
	ErrPhoneNotDeletable    = errx.FailedPrecond("PHONE_NOT_DELETABLE", "cannot delete primary or missing phone")
	ErrOwnerRoleImmutable   = errx.Forbidden("OWNER_ROLE_IMMUTABLE", "owner role cannot be assigned via API")
	ErrLastIdentity         = errx.FailedPrecond("LAST_IDENTITY", "cannot unlink the last identity")
	ErrGitHubNotConfigured  = errx.FailedPrecond("GITHUB_NOT_CONFIGURED", "configure GITHUB_CLIENT_ID to enable GitHub login")
	ErrGitHubExchangeFailed = errx.Unauthorized("GITHUB_EXCHANGE_FAILED", "github oauth exchange failed")
	ErrGitHubUserFailed     = errx.Internal("GITHUB_USER_FAILED", "failed to load github user")
	ErrGitHubLookupFailed   = errx.Internal("GITHUB_LOOKUP_FAILED", "failed to look up github identity")
	ErrGitHubSignupFailed   = errx.Internal("GITHUB_SIGNUP_FAILED", "github signup failed")
	ErrGitHubTaken          = errx.Conflict("GITHUB_TAKEN", "github account already linked")
	ErrInvalidOAuthState    = errx.Unauthorized("INVALID_OAUTH_STATE", "invalid oauth state")
)

/*
!PROFILE_NOT_OWNED
*en<Profile is not owned by account>
*zh<该资料不属于当前账号>
*fr<Le profil n'appartient pas au compte>

!INVALID_PROFILE_KIND
*en<Kind must be forger or authority>
*zh<资料类型必须是 forger 或 authority>
*fr<Le type doit être forger ou authority>

!INVALID_PROFILE_ID
*en<Invalid profile id>
*zh<资料 ID 无效>
*fr<Identifiant de profil invalide>

!INVALID_IMAGE_ID
*en<Invalid image id>
*zh<图片 ID 无效>
*fr<Identifiant d'image invalide>

!INVALID_EMAIL_ID
*en<Invalid email id>
*zh<邮箱 ID 无效>
*fr<Identifiant email invalide>

!INVALID_PHONE_ID
*en<Invalid phone id>
*zh<手机号 ID 无效>
*fr<Identifiant téléphone invalide>

!INVALID_EMAIL
*en<Email required>
*zh<邮箱必填>
*fr<Email requis>

!INVALID_PHONE
*en<Phone required>
*zh<手机号必填>
*fr<Téléphone requis>

!HASH_FAILED
*en<Failed to hash password>
*zh<密码哈希失败>
*fr<Échec du hachage du mot de passe>

!TOKEN_FAILED
*en<Token operation failed>
*zh<令牌操作失败>
*fr<Échec de l'opération sur le jeton>

!PHONE_NOT_DELETABLE
*en<Cannot delete primary or missing phone>
*zh<不能删除主手机号或不存在的手机号>
*fr<Impossible de supprimer le téléphone principal ou inexistant>

!OWNER_ROLE_IMMUTABLE
*en<Owner role cannot be assigned via API>
*zh<不能通过 API 分配所有者角色>
*fr<Le rôle propriétaire ne peut pas être attribué via l'API>

!LAST_IDENTITY
*en<Cannot unlink the last identity>
*zh<不能解绑最后一个登录身份>
*fr<Impossible de détacher la dernière identité>

!GITHUB_NOT_CONFIGURED
*en<Configure GITHUB_CLIENT_ID to enable GitHub login>
*zh<请配置 GITHUB_CLIENT_ID 以启用 GitHub 登录>
*fr<Configurez GITHUB_CLIENT_ID pour activer la connexion GitHub>

!GITHUB_EXCHANGE_FAILED
*en<GitHub OAuth exchange failed>
*zh<GitHub OAuth 交换失败>
*fr<Échec de l'échange OAuth GitHub>

!GITHUB_USER_FAILED
*en<Failed to load GitHub user>
*zh<获取 GitHub 用户失败>
*fr<Échec du chargement de l'utilisateur GitHub>

!GITHUB_LOOKUP_FAILED
*en<Failed to look up GitHub identity>
*zh<查找 GitHub 身份失败>
*fr<Échec de la recherche de l'identité GitHub>

!GITHUB_SIGNUP_FAILED
*en<GitHub signup failed>
*zh<GitHub 注册失败>
*fr<Inscription GitHub échouée>

!GITHUB_TAKEN
*en<GitHub account already linked>
*zh<该 GitHub 账号已被绑定>
*fr<Compte GitHub déjà lié>

!INVALID_OAUTH_STATE
*en<Invalid OAuth state>
*zh<OAuth state 无效>
*fr<État OAuth invalide>
*/
