package tenants

import "nfxidentity/pkgs/errx"

const (
	CodeMemberNotFound      = "MEMBER_NOT_FOUND"
	CodeUserIDRequired      = "USER_ID_REQUIRED"
	CodeMemberAlreadyExists = "MEMBER_ALREADY_EXISTS"
	CodeInvalidMemberStatus = "INVALID_MEMBER_STATUS"
	CodeInvalidMemberSource = "INVALID_MEMBER_SOURCE"
)

var (
	ErrMemberNotFound      = errx.NotFound(CodeMemberNotFound, "member not found")
	ErrMemberAlreadyExists = errx.Conflict(CodeMemberAlreadyExists, "member already exists")
	ErrUserIDRequired      = errx.InvalidArg(CodeUserIDRequired, "user id is required")
	ErrInvalidMemberStatus = errx.InvalidArg(CodeInvalidMemberStatus, "invalid member status")
	ErrInvalidMemberSource = errx.InvalidArg(CodeInvalidMemberSource, "invalid member source")
)

/*
!MEMBER_NOT_FOUND
*en<member not found>
*zh<成员不存在>
*fr<membre introuvable>

!USER_ID_REQUIRED
*en<user id required>
*zh<用户 ID 为必填>
*fr<identifiant utilisateur requis>

!MEMBER_ALREADY_EXISTS
*en<member already exists>
*zh<成员已存在>
*fr<membre existe déjà>

!INVALID_MEMBER_STATUS
*en<invalid member status>
*zh<无效的成员状态>
*fr<statut de membre invalide>

!INVALID_MEMBER_SOURCE
*en<invalid member source>
*zh<无效的成员来源>
*fr<source du membre invalide>

*/
