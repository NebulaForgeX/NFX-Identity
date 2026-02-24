package tenants

import "nfxid/pkgs/errx"

const (
	CodeInvitationNotFound        = "INVITATION_NOT_FOUND"
	CodeInviteIDRequired          = "INVITE_ID_REQUIRED"
	CodeEmailRequired             = "EMAIL_REQUIRED"
	CodeTokenHashRequired         = "TOKEN_HASH_REQUIRED"
	CodeInvitedByRequired         = "INVITED_BY_REQUIRED"
	CodeInviteIDAlreadyExists     = "INVITE_ID_ALREADY_EXISTS"
	CodeInvalidInvitationStatus   = "INVALID_INVITATION_STATUS"
	CodeInvitationExpired         = "INVITATION_EXPIRED"
	CodeInvitationAlreadyAccepted = "INVITATION_ALREADY_ACCEPTED"
	CodeInvitationAlreadyRevoked  = "INVITATION_ALREADY_REVOKED"
)

var (
	ErrInvitationNotFound        = errx.NotFound(CodeInvitationNotFound, "invitation not found")
	ErrInviteIDRequired          = errx.InvalidArg(CodeInviteIDRequired, "invite id is required")
	ErrEmailRequired             = errx.InvalidArg(CodeEmailRequired, "email is required")
	ErrTokenHashRequired         = errx.InvalidArg(CodeTokenHashRequired, "token hash is required")
	ErrInvitedByRequired         = errx.InvalidArg(CodeInvitedByRequired, "invited by is required")
	ErrInviteIDAlreadyExists     = errx.Conflict(CodeInviteIDAlreadyExists, "invite id already exists")
	ErrInvalidInvitationStatus   = errx.InvalidArg(CodeInvalidInvitationStatus, "invalid invitation status")
	ErrInvitationExpired         = errx.Expired(CodeInvitationExpired, "invitation expired")
	ErrInvitationAlreadyAccepted = errx.Conflict(CodeInvitationAlreadyAccepted, "invitation already accepted")
	ErrInvitationAlreadyRevoked  = errx.Conflict(CodeInvitationAlreadyRevoked, "invitation already revoked")
)

/*
!INVITATION_NOT_FOUND
*en<invitation not found>
*zh<邀请不存在>
*fr<invitation introuvable>

!INVITE_ID_REQUIRED
*en<invite id required>
*zh<邀请 ID 为必填>
*fr<id d'invitation requis>

!EMAIL_REQUIRED
*en<email required>
*zh<邮箱为必填>
*fr<email requis>

!TOKEN_HASH_REQUIRED
*en<token hash required>
*zh<令牌哈希为必填>
*fr<hachage du jeton requis>

!INVITED_BY_REQUIRED
*en<invited by required>
*zh<邀请人为必填>
*fr<invité par requis>

!INVITE_ID_ALREADY_EXISTS
*en<invite id already exists>
*zh<邀请 ID 已存在>
*fr<id d'invitation existe déjà>

!INVALID_INVITATION_STATUS
*en<invalid invitation status>
*zh<无效的邀请状态>
*fr<statut d'invitation invalide>

!INVITATION_EXPIRED
*en<invitation expired>
*zh<邀请已过期>
*fr<invitation expirée>

!INVITATION_ALREADY_ACCEPTED
*en<invitation already accepted>
*zh<邀请已接受>
*fr<invitation déjà acceptée>

!INVITATION_ALREADY_REVOKED
*en<invitation already revoked>
*zh<邀请已撤销>
*fr<invitation déjà révoquée>

*/
