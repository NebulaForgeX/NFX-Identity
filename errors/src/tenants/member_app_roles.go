package tenants

import "nfxid/pkgs/errx"

const (
	CodeMemberAppRoleNotFound       = "MEMBER_APP_ROLE_NOT_FOUND"
	CodeMemberAppRoleAlreadyExists  = "MEMBER_APP_ROLE_ALREADY_EXISTS"
	CodeMemberAppRoleAlreadyRevoked = "MEMBER_APP_ROLE_ALREADY_REVOKED"
	CodeMemberAppRoleExpired        = "MEMBER_APP_ROLE_EXPIRED"
)

var (
	ErrMemberAppRoleNotFound       = errx.NotFound(CodeMemberAppRoleNotFound, "member app role not found")
	ErrMemberAppRoleAlreadyExists  = errx.Conflict(CodeMemberAppRoleAlreadyExists, "member app role already exists")
	ErrMemberAppRoleAlreadyRevoked = errx.Conflict(CodeMemberAppRoleAlreadyRevoked, "member app role already revoked")
	ErrMemberAppRoleExpired        = errx.Expired(CodeMemberAppRoleExpired, "member app role expired")
)

/*
!MEMBER_APP_ROLE_NOT_FOUND
*en<member app role not found>
*zh<成员应用角色不存在>
*fr<rôle d'application du membre introuvable>

!MEMBER_APP_ROLE_ALREADY_EXISTS
*en<member app role already exists>
*zh<成员应用角色已存在>
*fr<rôle d'application du membre existe déjà>

!MEMBER_APP_ROLE_ALREADY_REVOKED
*en<member app role already revoked>
*zh<成员应用角色已撤销>
*fr<rôle d'application du membre déjà révoqué>

!MEMBER_APP_ROLE_EXPIRED
*en<member app role expired>
*zh<成员应用角色已过期>
*fr<rôle d'application du membre expiré>

*/
