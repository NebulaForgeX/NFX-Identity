package constants

// 图片存储路径常量
const (
	// StorageBasePath 存储根目录（磁盘路径）
	StorageBasePath = "./data"

	// StoragePathImages 图片存储根目录
	StoragePathImages = "images"

	// StoragePathTmp 临时上传目录
	StoragePathTmp = "images/tmp"

	// StoragePathAvatar 头像存储目录
	StoragePathAvatar = "images/avatar"

	// StoragePathBackground 背景图存储目录
	StoragePathBackground = "images/background"
)

// ImageStorageType 图片存储类型
type ImageStorageType string

const (
	ImageStorageTypeTmp        ImageStorageType = "tmp"
	ImageStorageTypeAvatar     ImageStorageType = "avatar"
	ImageStorageTypeBackground ImageStorageType = "background"
)

// GetStoragePath 根据类型获取存储路径
func GetStoragePath(storageType ImageStorageType) string {
	switch storageType {
	case ImageStorageTypeTmp:
		return StoragePathTmp
	case ImageStorageTypeAvatar:
		return StoragePathAvatar
	case ImageStorageTypeBackground:
		return StoragePathBackground
	default:
		return StoragePathTmp
	}
}
