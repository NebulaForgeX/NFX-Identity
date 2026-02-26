package image

import "nfxidentity/pkgs/errx"

const (
	CodeImageNotFound            = "IMAGE_NOT_FOUND"
	CodeFilenameRequired         = "FILENAME_REQUIRED"
	CodeOriginalFilenameRequired = "ORIGINAL_FILENAME_REQUIRED"
	CodeMimeTypeRequired         = "MIME_TYPE_REQUIRED"
	CodeSizeRequired             = "SIZE_REQUIRED"
)

var (
	ErrImageNotFound            = errx.NotFound(CodeImageNotFound, "image not found")
	ErrFilenameRequired         = errx.InvalidArg(CodeFilenameRequired, "filename is required")
	ErrOriginalFilenameRequired = errx.InvalidArg(CodeOriginalFilenameRequired, "original filename is required")
	ErrMimeTypeRequired         = errx.InvalidArg(CodeMimeTypeRequired, "mime type is required")
	ErrSizeRequired             = errx.InvalidArg(CodeSizeRequired, "size is required")
)

/*
!IMAGE_NOT_FOUND
*en<image not found>
*zh<图片不存在>
*fr<image introuvable>

!FILENAME_REQUIRED
*en<filename required>
*zh<文件名为必填>
*fr<nom de fichier requis>

!ORIGINAL_FILENAME_REQUIRED
*en<original filename required>
*zh<原文件名为必填>
*fr<nom de fichier d'origine requis>

!MIME_TYPE_REQUIRED
*en<mime type required>
*zh<MIME 类型为必填>
*fr<type MIME requis>

!SIZE_REQUIRED
*en<size required>
*zh<大小为必填>
*fr<taille requise>

*/
