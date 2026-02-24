package tenants

import "nfxid/pkgs/errx"

const (
	CodeGroupNotFound        = "GROUP_NOT_FOUND"
	CodeGroupIDAlreadyExists = "GROUP_ID_ALREADY_EXISTS"
	CodeInvalidGroupType     = "INVALID_GROUP_TYPE"
)

var (
	ErrGroupNotFound        = errx.NotFound(CodeGroupNotFound, "group not found")
	ErrGroupIDAlreadyExists = errx.Conflict(CodeGroupIDAlreadyExists, "group id already exists")
	ErrInvalidGroupType     = errx.InvalidArg(CodeInvalidGroupType, "invalid group type")
)

/*
!GROUP_NOT_FOUND
*en<group not found>
*zh<组不存在>
*fr<groupe introuvable>

!GROUP_ID_ALREADY_EXISTS
*en<group id already exists>
*zh<组 ID 已存在>
*fr<id de groupe existe déjà>

!INVALID_GROUP_TYPE
*en<invalid group type>
*zh<无效的组类型>
*fr<type de groupe invalide>

*/
