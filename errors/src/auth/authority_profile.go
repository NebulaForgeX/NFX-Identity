package auth

import "nfxidentity/pkgs/errx"

var (
	ErrAuthorityProfileNotFound             = errx.NotFound("AUTHORITY_PROFILE_NOT_FOUND", "authority profile not found")
	ErrAuthorityProfileAccountIDInvalid     = errx.InvalidArg("AUTHORITY_PROFILE_ACCOUNT_ID_INVALID", "invalid authority profile account id")
	ErrAuthorityProfileLanguageInvalid      = errx.InvalidArg("AUTHORITY_PROFILE_LANGUAGE_INVALID", "invalid authority profile language")
	ErrAuthorityProfileAuthorityRoleInvalid = errx.InvalidArg("AUTHORITY_PROFILE_AUTHORITY_ROLE_INVALID", "invalid authority profile authority role")
	ErrAuthorityProfilePatchEmpty           = errx.InvalidArg("AUTHORITY_PROFILE_PATCH_EMPTY", "no authority profile fields to update")
	ErrAuthorityProfileFieldInvalid         = errx.InvalidArg("AUTHORITY_PROFILE_FIELD_INVALID", "invalid authority profile field")
	ErrAuthorityProfileUpdateFailed         = errx.Internal("AUTHORITY_PROFILE_UPDATE_FAILED", "failed to update authority profile")
	ErrAuthorityProfileCreateFailed         = errx.Internal("AUTHORITY_PROFILE_CREATE_FAILED", "failed to create authority profile")
	ErrAuthorityProfileDeleteFailed         = errx.Internal("AUTHORITY_PROFILE_DELETE_FAILED", "failed to delete authority profile")
	ErrAuthorityProfileCountFailed          = errx.Internal("AUTHORITY_PROFILE_COUNT_FAILED", "failed to count authority profiles")
	ErrAuthorityProfileLimitReached         = errx.InvalidArg("AUTHORITY_PROFILE_LIMIT_REACHED", "authority profile limit reached")
	ErrAuthorityProfileCannotDeleteLast     = errx.InvalidArg("AUTHORITY_PROFILE_CANNOT_DELETE_LAST", "cannot delete the last authority profile")
	ErrAuthorityProfileCannotDeleteCurrent  = errx.InvalidArg("AUTHORITY_PROFILE_CANNOT_DELETE_CURRENT", "cannot delete the currently active authority profile")

	ErrAuthorityProfileAvatarAccountIDInvalid = errx.InvalidArg("AUTHORITY_PROFILE_AVATAR_ACCOUNT_ID_INVALID", "invalid authority profile avatar account id")
	ErrAuthorityProfileAvatarProfileIDInvalid = errx.InvalidArg("AUTHORITY_PROFILE_AVATAR_PROFILE_ID_INVALID", "invalid authority profile avatar profile id")
	ErrAuthorityProfileAvatarImageIDInvalid   = errx.InvalidArg("AUTHORITY_PROFILE_AVATAR_IMAGE_ID_INVALID", "invalid authority profile avatar image id")

	ErrAuthorityProfileBackgroundAccountIDInvalid = errx.InvalidArg(
		"AUTHORITY_PROFILE_BACKGROUND_ACCOUNT_ID_INVALID",
		"invalid authority profile background account id",
	)
	ErrAuthorityProfileBackgroundProfileIDInvalid = errx.InvalidArg(
		"AUTHORITY_PROFILE_BACKGROUND_PROFILE_ID_INVALID",
		"invalid authority profile background profile id",
	)
	ErrAuthorityProfileBackgroundImageIDInvalid = errx.InvalidArg(
		"AUTHORITY_PROFILE_BACKGROUND_IMAGE_ID_INVALID",
		"invalid authority profile background image id",
	)
	ErrAuthorityProfileBackgroundSortOrderInvalid = errx.InvalidArg(
		"AUTHORITY_PROFILE_BACKGROUND_SORT_ORDER_INVALID",
		"invalid authority profile background sort order",
	)
	ErrAuthorityProfileBackgroundConfirmFailed = errx.Internal(
		"AUTHORITY_PROFILE_BACKGROUND_CONFIRM_FAILED",
		"failed to confirm authority profile backgrounds",
	)
	ErrAuthorityProfileBackgroundDeleteFailed = errx.Internal(
		"AUTHORITY_PROFILE_BACKGROUND_DELETE_FAILED",
		"failed to delete authority profile background image",
	)

	ErrAuthorityProfileSearchFailed   = errx.Internal("AUTHORITY_PROFILE_SEARCH_FAILED", "failed to search authority profiles")
	ErrAuthorityProfileBatchGetFailed = errx.Internal("AUTHORITY_PROFILE_BATCH_GET_FAILED", "failed to batch get authority profiles")

	ErrAuthorityProfileSettingsNotFound         = errx.NotFound("AUTHORITY_PROFILE_SETTINGS_NOT_FOUND", "authority profile settings not found")
	ErrAuthorityProfileSettingsProfileIDInvalid = errx.InvalidArg(
		"AUTHORITY_PROFILE_SETTINGS_PROFILE_ID_INVALID",
		"invalid authority profile settings profile id",
	)
	ErrAuthorityProfileSettingsPatchEmpty   = errx.InvalidArg("AUTHORITY_PROFILE_SETTINGS_PATCH_EMPTY", "no authority profile settings fields to update")
	ErrAuthorityProfileSettingsUpdateFailed = errx.Internal("AUTHORITY_PROFILE_SETTINGS_UPDATE_FAILED", "failed to update authority profile settings")
	ErrAuthorityProfileSettingsCreateFailed = errx.Internal("AUTHORITY_PROFILE_SETTINGS_CREATE_FAILED", "failed to create authority profile settings")
	ErrAuthorityProfileSettingsGetFailed    = errx.Internal("AUTHORITY_PROFILE_SETTINGS_GET_FAILED", "failed to get authority profile settings")

	ErrAuthorityProfileInsufficientRole   = errx.Forbidden("AUTHORITY_PROFILE_INSUFFICIENT_ROLE", "insufficient authority profile role")
	ErrAuthorityProfileScopeRequired      = errx.Forbidden("AUTHORITY_PROFILE_SCOPE_REQUIRED", "authority profile session required")
	ErrAuthorityProfileOwnerRoleImmutable = errx.InvalidArg(
		"AUTHORITY_PROFILE_OWNER_ROLE_IMMUTABLE",
		"owner authority role can only be changed manually in the database",
	)
)

/*
!AUTHORITY_PROFILE_NOT_FOUND
*en<User profile not found>
*zh<用户资料不存在>
*fr<Profil utilisateur introuvable>

!AUTHORITY_PROFILE_ACCOUNT_ID_INVALID
*en<Invalid authority profile account id>
*zh<用户资料账号 ID 无效>
*fr<Identifiant de compte du profil invalide>

!AUTHORITY_PROFILE_LANGUAGE_INVALID
*en<Invalid authority profile language>
*zh<用户资料语言无效>
*fr<Langue du profil invalide>

!AUTHORITY_PROFILE_AUTHORITY_ROLE_INVALID
*en<Invalid authority profile authority role>
*zh<官方资料角色无效>
*fr<Rôle officiel du profil invalide>

!AUTHORITY_PROFILE_PATCH_EMPTY
*en<No authority profile fields to update>
*zh<未提供要更新的资料字段>
*fr<Aucun champ de profil à mettre à jour>

!AUTHORITY_PROFILE_FIELD_INVALID
*en<Invalid authority profile field>
*zh<用户资料字段无效>
*fr<Champ de profil invalide>

!AUTHORITY_PROFILE_UPDATE_FAILED
*en<Failed to update authority profile>
*zh<更新用户资料失败>
*fr<Échec de la mise à jour du profil>
*p*en<Sorry, we couldn't update your profile right now. Please try again later.>
*p*zh<抱歉，暂时无法更新资料，请稍后再试。>
*p*fr<Désolé, le profil n'a pas pu être mis à jour. Veuillez réessayer plus tard.>

!AUTHORITY_PROFILE_CREATE_FAILED
*en<Failed to create authority profile>
*zh<创建用户资料失败>
*fr<Échec de la création du profil>
*p*en<Sorry, we couldn't create your profile right now. This is on us—please try again later.>
*p*zh<抱歉，暂时无法创建资料，这是我们的问题，请稍后再试。>
*p*fr<Désolé, le profil n'a pas pu être créé. Veuillez réessayer plus tard.>

!AUTHORITY_PROFILE_DELETE_FAILED
*en<Failed to delete authority profile>
*zh<删除用户资料失败>
*fr<Échec de la suppression du profil>
*p*en<Sorry, we couldn't delete this profile right now. Please try again later.>
*p*zh<抱歉，暂时无法删除资料，请稍后再试。>
*p*fr<Désolé, le profil n'a pas pu être supprimé. Veuillez réessayer plus tard.>

!AUTHORITY_PROFILE_COUNT_FAILED
*en<Failed to count authority profiles>
*zh<统计用户资料数量失败>
*fr<Échec du comptage des profils>
*p*en<Sorry, we couldn't complete this count right now. Please try again later.>
*p*zh<抱歉，暂时无法完成统计，请稍后再试。>
*p*fr<Désolé, le décompte n'a pas pu aboutir. Veuillez réessayer plus tard.>

!AUTHORITY_PROFILE_LIMIT_REACHED
*en<User profile limit reached>
*zh<已达到资料数量上限>
*fr<Limite de profils atteinte>

!AUTHORITY_PROFILE_CANNOT_DELETE_LAST
*en<Cannot delete the last authority profile>
*zh<无法删除最后一个资料>
*fr<Impossible de supprimer le dernier profil>

!AUTHORITY_PROFILE_CANNOT_DELETE_CURRENT
*en<Cannot delete the currently active authority profile>
*zh<无法删除当前正在使用的资料>
*fr<Impossible de supprimer le profil actif>

!AUTHORITY_PROFILE_OWNER_ROLE_IMMUTABLE
*en<Owner authority role can only be changed manually in the database>
*zh<owner 角色仅可通过数据库手动修改>
*fr<Le rôle owner ne peut être modifié que manuellement dans la base de données>

!AUTHORITY_PROFILE_AVATAR_ACCOUNT_ID_INVALID
*en<Invalid authority profile avatar account id>
*zh<用户头像关联账号 ID 无效>
*fr<Identifiant de compte de l'avatar invalide>

!AUTHORITY_PROFILE_AVATAR_PROFILE_ID_INVALID
*en<Invalid authority profile avatar profile id>
*zh<用户头像关联资料 ID 无效>
*fr<Identifiant de profil de l'avatar invalide>

!AUTHORITY_PROFILE_AVATAR_IMAGE_ID_INVALID
*en<Invalid authority profile avatar image id>
*zh<用户头像图片 ID 无效>
*fr<Identifiant d'image de l'avatar invalide>

!AUTHORITY_PROFILE_BACKGROUND_ACCOUNT_ID_INVALID
*en<Invalid authority profile background account id>
*zh<用户背景图关联账号 ID 无效>
*fr<Identifiant de compte du fond invalide>

!AUTHORITY_PROFILE_BACKGROUND_PROFILE_ID_INVALID
*en<Invalid authority profile background profile id>
*zh<用户背景图关联资料 ID 无效>
*fr<Identifiant de profil du fond invalide>

!AUTHORITY_PROFILE_BACKGROUND_IMAGE_ID_INVALID
*en<Invalid authority profile background image id>
*zh<用户背景图图片 ID 无效>
*fr<Identifiant d'image du fond invalide>

!AUTHORITY_PROFILE_BACKGROUND_SORT_ORDER_INVALID
*en<Invalid authority profile background sort order>
*zh<用户背景图排序无效>
*fr<Ordre de tri du fond invalide>

!AUTHORITY_PROFILE_BACKGROUND_CONFIRM_FAILED
*en<Failed to confirm authority profile backgrounds>
*zh<确认用户背景图失败>
*fr<Échec de la confirmation des fonds du profil>
*p*en<Sorry, we couldn't save your background images right now. Please try again later.>
*p*zh<抱歉，暂时无法保存背景图，请稍后再试。>
*p*fr<Désolé, les images de fond n'ont pas pu être enregistrées. Veuillez réessayer plus tard.>

!AUTHORITY_PROFILE_BACKGROUND_DELETE_FAILED
*en<Failed to delete authority profile background image>
*zh<删除用户背景图失败>
*fr<Échec de la suppression de l'image de fond>
*p*en<Sorry, we couldn't remove this background image right now. Please try again later.>
*p*zh<抱歉，暂时无法删除背景图，请稍后再试。>
*p*fr<Désolé, l'image de fond n'a pas pu être supprimée. Veuillez réessayer plus tard.>

!AUTHORITY_PROFILE_SEARCH_FAILED
*en<Failed to search authority profiles>
*zh<搜索用户资料失败>
*fr<Échec de la recherche de profils>
*p*en<Sorry, search isn't working right now. Please try again later.>
*p*zh<抱歉，暂时无法搜索资料，请稍后再试。>
*p*fr<Désolé, la recherche de profils a échoué. Veuillez réessayer plus tard.>

!AUTHORITY_PROFILE_BATCH_GET_FAILED
*en<Failed to batch get authority profiles>
*zh<批量获取用户资料失败>
*fr<Échec de la récupération groupée des profils>
*p*en<Sorry, we couldn't load these profiles right now. Please try again later.>
*p*zh<抱歉，暂时无法加载资料，请稍后再试。>
*p*fr<Désolé, les profils n'ont pas pu être chargés. Veuillez réessayer plus tard.>

!AUTHORITY_PROFILE_SETTINGS_NOT_FOUND
*en<User profile settings not found>
*zh<用户资料设置不存在>
*fr<Paramètres du profil introuvables>

!AUTHORITY_PROFILE_SETTINGS_PROFILE_ID_INVALID
*en<Invalid authority profile settings profile id>
*zh<用户资料设置的资料 ID 无效>
*fr<Identifiant de profil des paramètres invalide>

!AUTHORITY_PROFILE_SETTINGS_PATCH_EMPTY
*en<No authority profile settings fields to update>
*zh<未提供要更新的资料设置字段>
*fr<Aucun paramètre de profil à mettre à jour>

!AUTHORITY_PROFILE_SETTINGS_UPDATE_FAILED
*en<Failed to update authority profile settings>
*zh<更新用户资料设置失败>
*fr<Échec de la mise à jour des paramètres du profil>
*p*en<Sorry, we couldn't update your settings right now. Please try again later.>
*p*zh<抱歉，暂时无法更新设置，请稍后再试。>
*p*fr<Désolé, les paramètres n'ont pas pu être mis à jour. Veuillez réessayer plus tard.>

!AUTHORITY_PROFILE_SETTINGS_CREATE_FAILED
*en<Failed to create authority profile settings>
*zh<创建用户资料设置失败>
*fr<Échec de la création des paramètres du profil>
*p*en<Sorry, we couldn't create your settings right now. This is on us—please try again later.>
*p*zh<抱歉，暂时无法创建设置，这是我们的问题，请稍后再试。>
*p*fr<Désolé, les paramètres n'ont pas pu être créés. Veuillez réessayer plus tard.>

!AUTHORITY_PROFILE_SETTINGS_GET_FAILED
*en<Failed to get authority profile settings>
*zh<获取用户资料设置失败>
*fr<Échec de la récupération des paramètres du profil>
*p*en<Sorry, we couldn't load your settings right now. Please try again later.>
*p*zh<抱歉，暂时无法加载设置，请稍后再试。>
*p*fr<Désolé, les paramètres n'ont pas pu être chargés. Veuillez réessayer plus tard.>

!AUTHORITY_PROFILE_INSUFFICIENT_ROLE
*en<Insufficient authority profile role>
*zh<官方资料权限不足>
*fr<Rôle du profil officiel insuffisant>

!AUTHORITY_PROFILE_SCOPE_REQUIRED
*en<This action requires an authority profile session. Switch profile and sign in again.>
*zh<此操作需要权限资料会话，请切换至权限资料后重新登录。>
*fr<Cette action nécessite une session profil officiel. Changez de profil puis reconnectez-vous.>
*p*en<You're signed in with a community profile. Switch to an authority profile to continue.>
*p*zh<当前登录的是社区资料，请切换到权限资料后再访问此页面。>
*p*fr<Vous êtes connecté avec un profil communautaire. Passez à un profil officiel pour continuer.>
*/
