package image

import "nfxidentity/pkgs/errx"

const (
	CodeImageTypeNotFound      = "IMAGE_TYPE_NOT_FOUND"
	CodeKeyRequired            = "KEY_REQUIRED"
	CodeKeyAlreadyExists       = "KEY_ALREADY_EXISTS"
	CodeCannotDeleteSystemType = "CANNOT_DELETE_SYSTEM_TYPE"
)

var (
	ErrImageTypeNotFound      = errx.NotFound(CodeImageTypeNotFound, "image type not found")
	ErrKeyRequired            = errx.InvalidArg(CodeKeyRequired, "key is required")
	ErrKeyAlreadyExists       = errx.Conflict(CodeKeyAlreadyExists, "key already exists")
	ErrCannotDeleteSystemType = errx.FailedPrecond(CodeCannotDeleteSystemType, "cannot delete system type")
)

/*
!IMAGE_TYPE_NOT_FOUND
*en<image type not found>
*zh<图片类型不存在>
*fr<type d'image introuvable>

!KEY_REQUIRED
*en<key required>
*zh<键为必填>
*fr<clé requise>

!KEY_ALREADY_EXISTS
*en<key already exists>
*zh<键已存在>
*fr<clé existe déjà>

!CANNOT_DELETE_SYSTEM_TYPE
*en<cannot delete system type>
*zh<无法删除系统类型>
*fr<impossible de supprimer le type système>

*/
