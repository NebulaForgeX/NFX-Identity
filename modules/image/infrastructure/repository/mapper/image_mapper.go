package mapper

import (
	"encoding/json"
	imageDomain "nebulaid/modules/image/domain/image"
	"nebulaid/modules/image/infrastructure/rdb/models"
	"nebulaid/pkgs/utils/timex"

	"gorm.io/datatypes"
)

func ImageDomainToModel(img *imageDomain.Image) *models.Image {
	if img == nil {
		return nil
	}

	var metadataJSON *datatypes.JSON
	if img.Metadata() != nil {
		data, _ := json.Marshal(img.Metadata())
		jsonData := datatypes.JSON(data)
		metadataJSON = &jsonData
	}

	return &models.Image{
		ID:               img.ID(),
		TypeID:           img.TypeID(),
		UserID:           img.UserID(),
		SourceDomain:     img.SourceDomain(),
		Filename:         img.Filename(),
		OriginalFilename: img.OriginalFilename(),
		MimeType:         img.MimeType(),
		Size:             img.Size(),
		Width:            img.Width(),
		Height:           img.Height(),
		StoragePath:      img.StoragePath(),
		URL:              img.URL(),
		IsPublic:         img.IsPublic(),
		Metadata:         metadataJSON,
		CreatedAt:        img.CreatedAt(),
		UpdatedAt:        img.UpdatedAt(),
		DeletedAt:        timex.TimeToGormDeletedAt(img.DeletedAt()),
	}
}

func ImageModelToDomain(m *models.Image) *imageDomain.Image {
	if m == nil {
		return nil
	}

	var metadata map[string]interface{}
	if m.Metadata != nil {
		_ = json.Unmarshal([]byte(*m.Metadata), &metadata)
	}

	editable := imageDomain.ImageEditable{
		TypeID:           m.TypeID,
		UserID:           m.UserID,
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
	}

	state := imageDomain.ImageState{
		ID:               m.ID,
		TypeID:           editable.TypeID,
		UserID:           editable.UserID,
		SourceDomain:     editable.SourceDomain,
		Filename:         editable.Filename,
		OriginalFilename: editable.OriginalFilename,
		MimeType:         editable.MimeType,
		Size:             editable.Size,
		Width:            editable.Width,
		Height:           editable.Height,
		StoragePath:      editable.StoragePath,
		URL:              editable.URL,
		IsPublic:         editable.IsPublic,
		Metadata:         editable.Metadata,
		CreatedAt:        m.CreatedAt,
		UpdatedAt:        m.UpdatedAt,
		DeletedAt:        timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return imageDomain.NewImageFromState(state)
}

func ImageModelsToUpdates(m *models.Image) map[string]any {
	var metadataJSON *datatypes.JSON
	if m.Metadata != nil {
		metadataJSON = m.Metadata
	}

	return map[string]any{
		models.ImageCols.TypeID:           m.TypeID,
		models.ImageCols.UserID:           m.UserID,
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
		models.ImageCols.Metadata:         metadataJSON,
		models.ImageCols.DeletedAt:        m.DeletedAt,
	}
}
