package mapper

import (
	"encoding/json"
	"nfxid/modules/image/domain/images"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"

	"gorm.io/datatypes"
)

// ImageDomainToModel 将 Domain Image 转换为 Model Image
func ImageDomainToModel(i *images.Image) *models.Image {
	if i == nil {
		return nil
	}

	var metadata *datatypes.JSON
	if i.Metadata() != nil && len(i.Metadata()) > 0 {
		metaBytes, _ := json.Marshal(i.Metadata())
		jsonData := datatypes.JSON(metaBytes)
		metadata = &jsonData
	}

	return &models.Image{
		ID:               i.ID(),
		TypeID:           i.TypeID(),
		UserID:           i.UserID(),
		TenantID:         i.TenantID(),
		AppID:            i.AppID(),
		SourceDomain:     i.SourceDomain(),
		Filename:         i.Filename(),
		OriginalFilename: i.OriginalFilename(),
		MimeType:         i.MimeType(),
		Size:             i.Size(),
		Width:            i.Width(),
		Height:           i.Height(),
		StoragePath:      i.StoragePath(),
		URL:              i.URL(),
		IsPublic:         i.IsPublic(),
		Metadata:         metadata,
		CreatedAt:        i.CreatedAt(),
		UpdatedAt:        i.UpdatedAt(),
		DeletedAt:        timex.TimeToGormDeletedAt(i.DeletedAt()),
	}
}

// ImageModelToDomain 将 Model Image 转换为 Domain Image
func ImageModelToDomain(m *models.Image) *images.Image {
	if m == nil {
		return nil
	}

	var metadata map[string]interface{}
	if m.Metadata != nil {
		json.Unmarshal(*m.Metadata, &metadata)
	}

	state := images.ImageState{
		ID:               m.ID,
		TypeID:           m.TypeID,
		UserID:           m.UserID,
		TenantID:         m.TenantID,
		AppID:            m.AppID,
		SourceDomain:     m.SourceDomain,
		Filename:         m.Filename,
		OriginalFilename: m.OriginalFilename,
		MimeType:         m.MimeType,
		Size:             m.Size,
		Width:            m.Width,
		Height:           m.Height,
		StoragePath:      m.StoragePath,
		URL:              m.URL,
		IsPublic:         m.IsPublic,
		Metadata:         metadata,
		CreatedAt:        m.CreatedAt,
		UpdatedAt:        m.UpdatedAt,
		DeletedAt:        timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return images.NewImageFromState(state)
}

// ImageModelToUpdates 将 Model Image 转换为更新字段映射
func ImageModelToUpdates(m *models.Image) map[string]any {
	var metadata any
	if m.Metadata != nil {
		metadata = m.Metadata
	}

	return map[string]any{
		models.ImageCols.TypeID:           m.TypeID,
		models.ImageCols.UserID:           m.UserID,
		models.ImageCols.TenantID:         m.TenantID,
		models.ImageCols.AppID:            m.AppID,
		models.ImageCols.SourceDomain:     m.SourceDomain,
		models.ImageCols.Filename:         m.Filename,
		models.ImageCols.OriginalFilename: m.OriginalFilename,
		models.ImageCols.MimeType:         m.MimeType,
		models.ImageCols.Size:             m.Size,
		models.ImageCols.Width:            m.Width,
		models.ImageCols.Height:           m.Height,
		models.ImageCols.StoragePath:      m.StoragePath,
		models.ImageCols.URL:              m.URL,
		models.ImageCols.IsPublic:         m.IsPublic,
		models.ImageCols.Metadata:         metadata,
		models.ImageCols.UpdatedAt:        m.UpdatedAt,
		models.ImageCols.DeletedAt:        m.DeletedAt,
	}
}
