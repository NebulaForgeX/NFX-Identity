package image_type

import (
	"time"
)

func (it *ImageType) Update(editable ImageTypeEditable) {
	it.state.Key = editable.Key
	it.state.Description = editable.Description
	it.state.MaxWidth = editable.MaxWidth
	it.state.MaxHeight = editable.MaxHeight
	it.state.AspectRatio = editable.AspectRatio
	it.state.IsSystem = editable.IsSystem
	it.state.UpdatedAt = time.Now()
}

func (it *ImageType) SetDescription(description *string) {
	it.state.Description = description
	it.state.UpdatedAt = time.Now()
}

func (it *ImageType) SetDimensions(maxWidth, maxHeight *int) {
	it.state.MaxWidth = maxWidth
	it.state.MaxHeight = maxHeight
	it.state.UpdatedAt = time.Now()
}

func (it *ImageType) SetAspectRatio(aspectRatio *string) {
	it.state.AspectRatio = aspectRatio
	it.state.UpdatedAt = time.Now()
}

