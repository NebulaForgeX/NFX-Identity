package image

import "nfxidentity/pkgs/errx"

const (
	CodeImageVariantNotFound      = "IMAGE_VARIANT_NOT_FOUND"
	CodeVariantKeyRequired        = "VARIANT_KEY_REQUIRED"
	CodeImageVariantAlreadyExists = "IMAGE_VARIANT_ALREADY_EXISTS"
)

var (
	ErrImageVariantNotFound      = errx.NotFound(CodeImageVariantNotFound, "image variant not found")
	ErrVariantKeyRequired        = errx.InvalidArg(CodeVariantKeyRequired, "variant key is required")
	ErrImageVariantAlreadyExists = errx.Conflict(CodeImageVariantAlreadyExists, "image variant already exists")
)

/*
!IMAGE_VARIANT_NOT_FOUND
*en<image variant not found>
*zh<图片变体不存在>
*fr<variante d'image introuvable>

!VARIANT_KEY_REQUIRED
*en<variant key required>
*zh<变体键为必填>
*fr<clé de variante requise>

!IMAGE_VARIANT_ALREADY_EXISTS
*en<image variant already exists>
*zh<图片变体已存在>
*fr<variante d'image existe déjà>

*/
