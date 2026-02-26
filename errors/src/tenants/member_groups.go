package tenants

import "nfxidentity/pkgs/errx"

const (
	CodeMemberGroupNotFound       = "MEMBER_GROUP_NOT_FOUND"
	CodeMemberGroupAlreadyExists  = "MEMBER_GROUP_ALREADY_EXISTS"
	CodeMemberGroupAlreadyRevoked = "MEMBER_GROUP_ALREADY_REVOKED"
)

var (
	ErrMemberGroupNotFound       = errx.NotFound(CodeMemberGroupNotFound, "member group not found")
	ErrMemberGroupAlreadyExists  = errx.Conflict(CodeMemberGroupAlreadyExists, "member group already exists")
	ErrMemberGroupAlreadyRevoked = errx.Conflict(CodeMemberGroupAlreadyRevoked, "member group already revoked")
)

/*
!MEMBER_GROUP_NOT_FOUND
*en<member group not found>
*zh<成员组不存在>
*fr<groupe du membre introuvable>

!MEMBER_GROUP_ALREADY_EXISTS
*en<member group already exists>
*zh<成员组已存在>
*fr<groupe du membre existe déjà>

!MEMBER_GROUP_ALREADY_REVOKED
*en<member group already revoked>
*zh<成员组已撤销>
*fr<groupe du membre déjà révoqué>

*/
