package auth

import "nfxidentity/pkgs/errx"

var (
	ErrProfileNotOwned    = errx.Forbidden("PROFILE_NOT_OWNED", "profile is not owned by account")
	ErrInvalidProfileKind = errx.InvalidArg("INVALID_PROFILE_KIND", "kind must be community or authority")
	ErrInvalidProfileID   = errx.InvalidArg("INVALID_PROFILE_ID", "invalid profile id")
	ErrInvalidImageID     = errx.InvalidArg("INVALID_IMAGE_ID", "invalid image id")
	ErrInvalidEmailID     = errx.InvalidArg("INVALID_EMAIL_ID", "invalid email id")
	ErrInvalidPhoneID     = errx.InvalidArg("INVALID_PHONE_ID", "invalid phone id")
	ErrInvalidEmail       = errx.InvalidArg("INVALID_EMAIL", "email required")
	ErrInvalidPhone       = errx.InvalidArg("INVALID_PHONE", "phone required")
	ErrHashFailed         = errx.Internal("HASH_FAILED", "failed to hash password")
	ErrTokenFailed        = errx.Internal("TOKEN_FAILED", "token operation failed")
	ErrPhoneNotDeletable  = errx.FailedPrecond("PHONE_NOT_DELETABLE", "cannot delete primary or missing phone")
	ErrOwnerRoleImmutable = errx.Forbidden("OWNER_ROLE_IMMUTABLE", "owner role cannot be assigned via API")
)

/*
!PROFILE_NOT_OWNED
*en<Profile is not owned by account>
*zh<该资料不属于当前账号>
*fr<Le profil n'appartient pas au compte>

!INVALID_PROFILE_KIND
*en<Kind must be community or authority>
*zh<资料类型必须是 community 或 authority>
*fr<Le type doit être community ou authority>

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
*/
