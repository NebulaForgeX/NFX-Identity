package respdto

import (
	imageTypeAppViews "nebulaid/modules/image/application/image_type/views"
)

func ImageTypeViewToDTO(v *imageTypeAppViews.ImageTypeView) map[string]interface{} {
	result := map[string]interface{}{
		"id":         v.ID,
		"key":        v.Key,
		"is_system":  v.IsSystem,
		"created_at": v.CreatedAt,
		"updated_at": v.UpdatedAt,
	}

	if v.Description != "" {
		result["description"] = v.Description
	}
	if v.MaxWidth != nil {
		result["max_width"] = *v.MaxWidth
	}
	if v.MaxHeight != nil {
		result["max_height"] = *v.MaxHeight
	}
	if v.AspectRatio != "" {
		result["aspect_ratio"] = v.AspectRatio
	}

	return result
}

func ImageTypeListViewToDTO(items []imageTypeAppViews.ImageTypeView) []map[string]interface{} {
	result := make([]map[string]interface{}, len(items))
	for i, item := range items {
		result[i] = ImageTypeViewToDTO(&item)
	}
	return result
}
