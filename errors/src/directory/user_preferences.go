package directory

import "nfxidentity/pkgs/errx"

const (
	CodeUserPreferenceNotFound      = "USER_PREFERENCE_NOT_FOUND"
	CodeUserPreferenceAlreadyExists = "USER_PREFERENCE_ALREADY_EXISTS"
)

var (
	ErrUserPreferenceNotFound      = errx.NotFound(CodeUserPreferenceNotFound, "user preference not found")
	ErrUserPreferenceAlreadyExists = errx.Conflict(CodeUserPreferenceAlreadyExists, "user preference already exists")
)

/*
!USER_PREFERENCE_NOT_FOUND
*en<user preference not found>
*zh<用户偏好不存在>
*fr<préférence utilisateur introuvable>

!USER_PREFERENCE_ALREADY_EXISTS
*en<user preference already exists>
*zh<用户偏好已存在>
*fr<préférence utilisateur existe déjà>

*/
