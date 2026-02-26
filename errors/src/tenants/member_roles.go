package tenants

import "nfxidentity/pkgs/errx"

const (
	CodeMemberRoleNotFound       = "MEMBER_ROLE_NOT_FOUND"
	CodeMemberRoleAlreadyExists  = "MEMBER_ROLE_ALREADY_EXISTS"
	CodeMemberRoleAlreadyRevoked = "MEMBER_ROLE_ALREADY_REVOKED"
	CodeMemberRoleExpired        = "MEMBER_ROLE_EXPIRED"
)

var (
	ErrMemberRoleNotFound       = errx.NotFound(CodeMemberRoleNotFound, "member role not found")
	ErrMemberRoleAlreadyExists  = errx.Conflict(CodeMemberRoleAlreadyExists, "member role already exists")
	ErrMemberRoleAlreadyRevoked = errx.Conflict(CodeMemberRoleAlreadyRevoked, "member role already revoked")
	ErrMemberRoleExpired        = errx.Expired(CodeMemberRoleExpired, "member role expired")
)

/*
!MEMBER_ROLE_NOT_FOUND
*en<member role not found>
*zh<成员角色不存在>
*fr<rôle du membre introuvable>

!MEMBER_ROLE_ALREADY_EXISTS
*en<member role already exists>
*zh<成员角色已存在>
*fr<rôle du membre existe déjà>

!MEMBER_ROLE_ALREADY_REVOKED
*en<member role already revoked>
*zh<成员角色已撤销>
*fr<rôle du membre déjà révoqué>

!MEMBER_ROLE_EXPIRED
*en<member role expired>
*zh<成员角色已过期>
*fr<rôle du membre expiré>

*/
