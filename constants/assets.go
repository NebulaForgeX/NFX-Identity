package constants

import (
	"path/filepath"
	"strings"
	"time"

	"nfxidentity/enums"
	"nfxidentity/pkgs/constantx"

	"github.com/google/uuid"
)

// AssetType S3 路径类型段：image | video | file | audio
type AssetType string

const (
	AssetTypeImage AssetType = "image"
	AssetTypeVideo AssetType = "video"
	AssetTypeAudio AssetType = "audio"
	AssetTypeFile  AssetType = "file"
)

// AssetCategory S3 路径 category 段（用途/槽位）
type AssetCategory string

const (
	AssetCategoryAvatar     AssetCategory = "avatar"
	AssetCategoryBackground AssetCategory = "background"
	AssetCategoryMessage    AssetCategory = "message"

	AssetCategoryNews      AssetCategory = "news"
	AssetCategoryEmergency AssetCategory = "emergency"
	AssetCategoryTraffic   AssetCategory = "traffic"
	AssetCategoryActivity  AssetCategory = "activity"
	AssetCategoryLife      AssetCategory = "life"
	AssetCategoryProduct   AssetCategory = "product"
	AssetCategoryDiscovery AssetCategory = "discovery"
	AssetCategoryOther     AssetCategory = "other"

	AssetCategoryBadminton AssetCategory = "badminton"
)

var AssetTypes = constantx.NewStringEnumSet(
	AssetTypeImage,
	AssetTypeVideo,
	AssetTypeAudio,
	AssetTypeFile,
)

var AssetCategories = constantx.NewStringEnumSet(
	AssetCategoryAvatar,
	AssetCategoryBackground,
	AssetCategoryMessage,
	AssetCategoryNews,
	AssetCategoryEmergency,
	AssetCategoryTraffic,
	AssetCategoryActivity,
	AssetCategoryLife,
	AssetCategoryProduct,
	AssetCategoryDiscovery,
	AssetCategoryOther,
	AssetCategoryBadminton,
)

const assetPathSegmentTmp = "tmp"

const (
	StaleTmpMaxAge    = 24 * time.Hour
	StaleTmpBatchSize = 200
	StaleTmpPathLike  = "%/tmp/%"
)

// IsTmpFilePath reports whether filePath is under an asset tmp segment.
func IsTmpFilePath(filePath string) bool {
	p := filepath.ToSlash(strings.TrimPrefix(filePath, "/"))
	parts := strings.Split(p, "/")
	for _, part := range parts {
		if part == assetPathSegmentTmp {
			return true
		}
	}
	return false
}

// BuildAssetTmpPath 暂存：{account_id}/{profile_scope}/{profile_id}/{asset_type}/tmp/{asset_category}/{filename}
func BuildAssetTmpPath(accountID, profileID uuid.UUID, scope enums.AuthProfileScope, assetType AssetType, category AssetCategory, filename string) string {
	parts := []string{
		accountID.String(),
		string(scope),
		profileID.String(),
		string(assetType),
		assetPathSegmentTmp,
		string(category),
		filename,
	}
	return filepath.Join(parts...)
}

// BuildAssetPath 正式：{account_id}/{profile_scope}/{profile_id}/{asset_type}/{asset_category}/{filename}
func BuildAssetPath(accountID, profileID uuid.UUID, scope enums.AuthProfileScope, assetType AssetType, category AssetCategory, filename string) string {
	parts := []string{
		accountID.String(),
		string(scope),
		profileID.String(),
		string(assetType),
		string(category),
		filename,
	}
	return filepath.Join(parts...)
}

// MatchesTmpPath 判断 filePath 是否为指定 account + scope + profile + asset_type + category 的 tmp 路径。
func MatchesTmpPath(scope enums.AuthProfileScope, filePath string, accountID, profileID uuid.UUID, assetType AssetType, category AssetCategory) bool {
	p := filepath.ToSlash(strings.TrimPrefix(filePath, "/"))
	parts := strings.Split(p, "/")
	if len(parts) < 7 {
		return false
	}
	return parts[0] == accountID.String() &&
		parts[1] == string(scope) &&
		parts[2] == profileID.String() &&
		parts[3] == string(assetType) &&
		parts[4] == assetPathSegmentTmp &&
		parts[5] == string(category)
}

// MatchesPermanentPath reports whether filePath is a confirmed asset path for the given account/scope/profile/type/category.
func MatchesPermanentPath(scope enums.AuthProfileScope, filePath string, accountID, profileID uuid.UUID, assetType AssetType, category AssetCategory) bool {
	p := filepath.ToSlash(strings.TrimPrefix(filePath, "/"))
	parts := strings.Split(p, "/")
	if len(parts) < 6 {
		return false
	}
	return parts[0] == accountID.String() &&
		parts[1] == string(scope) &&
		parts[2] == profileID.String() &&
		parts[3] == string(assetType) &&
		parts[4] == string(category) &&
		!IsTmpFilePath(filePath)
}
