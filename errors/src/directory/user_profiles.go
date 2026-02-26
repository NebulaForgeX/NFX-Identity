package directory

import "nfxidentity/pkgs/errx"

const (
	CodeUserProfileNotFound      = "USER_PROFILE_NOT_FOUND"
	CodeUserProfileAlreadyExists = "USER_PROFILE_ALREADY_EXISTS"
)

var (
	ErrUserProfileNotFound      = errx.NotFound(CodeUserProfileNotFound, "user profile not found")
	ErrUserProfileAlreadyExists = errx.Conflict(CodeUserProfileAlreadyExists, "user profile already exists")
)

/*
!USER_PROFILE_NOT_FOUND
*en<user profile not found>
*zh<用户资料不存在>
*fr<profil utilisateur introuvable>

!USER_PROFILE_ALREADY_EXISTS
*en<user profile already exists>
*zh<用户资料已存在>
*fr<profil utilisateur existe déjà>

*/
