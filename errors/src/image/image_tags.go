package image

import "nfxid/pkgs/errx"

const (
	CodeImageTagNotFound      = "IMAGE_TAG_NOT_FOUND"
	CodeTagRequired           = "TAG_REQUIRED"
	CodeImageTagAlreadyExists = "IMAGE_TAG_ALREADY_EXISTS"
)

var (
	ErrImageTagNotFound      = errx.NotFound(CodeImageTagNotFound, "image tag not found")
	ErrTagRequired           = errx.InvalidArg(CodeTagRequired, "tag is required")
	ErrImageTagAlreadyExists = errx.Conflict(CodeImageTagAlreadyExists, "image tag already exists")
)

/*
!IMAGE_TAG_NOT_FOUND
*en<image tag not found>
*zh<图片标签不存在>
*fr<étiquette d'image introuvable>

!TAG_REQUIRED
*en<tag required>
*zh<标签为必填>
*fr<étiquette requise>

!IMAGE_TAG_ALREADY_EXISTS
*en<image tag already exists>
*zh<图片标签已存在>
*fr<étiquette d'image existe déjà>

*/
