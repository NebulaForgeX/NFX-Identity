package respdto

import (
	imageAppViews "nfxid/modules/image/application/image/views"
)

func ImageViewToDTO(v *imageAppViews.ImageView) map[string]interface{} {
	result := map[string]interface{}{
		"id":                v.ID,
		"filename":          v.Filename,
		"original_filename": v.OriginalFilename,
		"mime_type":         v.MimeType,
		"size":              v.Size,
		"storage_path":      v.StoragePath,
		"is_public":         v.IsPublic,
		"created_at":        v.CreatedAt,
		"updated_at":        v.UpdatedAt,
	}

	if v.TypeID != "" {
		result["type_id"] = v.TypeID
	}
	if v.UserID != "" {
		result["user_id"] = v.UserID
	}
	if v.SourceDomain != "" {
		result["source_domain"] = v.SourceDomain
	}
	if v.Width != nil {
		result["width"] = *v.Width
	}
	if v.Height != nil {
		result["height"] = *v.Height
	}
	if v.URL != "" {
		result["url"] = v.URL
	}
	if v.Metadata != nil {
		result["metadata"] = v.Metadata
	}

	return result
}

func ImageListViewToDTO(items []imageAppViews.ImageView) []map[string]interface{} {
	result := make([]map[string]interface{}, len(items))
	for i, item := range items {
		result[i] = ImageViewToDTO(&item)
	}
	return result
}
