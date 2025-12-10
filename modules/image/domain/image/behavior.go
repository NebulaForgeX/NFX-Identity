package image

import (
	"time"
)

func (i *Image) Update(editable ImageEditable) {
	i.state.TypeID = editable.TypeID
	i.state.UserID = editable.UserID
	i.state.SourceDomain = editable.SourceDomain
	i.state.Filename = editable.Filename
	i.state.OriginalFilename = editable.OriginalFilename
	i.state.MimeType = editable.MimeType
	i.state.Size = editable.Size
	i.state.Width = editable.Width
	i.state.Height = editable.Height
	i.state.StoragePath = editable.StoragePath
	i.state.URL = editable.URL
	i.state.IsPublic = editable.IsPublic
	i.state.Metadata = editable.Metadata
	i.state.UpdatedAt = time.Now()
}

func (i *Image) Delete() {
	now := time.Now()
	i.state.DeletedAt = &now
	i.state.UpdatedAt = now
}

func (i *Image) SetPublic(isPublic bool) {
	i.state.IsPublic = isPublic
	i.state.UpdatedAt = time.Now()
}

func (i *Image) SetURL(url string) {
	i.state.URL = &url
	i.state.UpdatedAt = time.Now()
}

func (i *Image) SetMetadata(metadata map[string]interface{}) {
	i.state.Metadata = metadata
	i.state.UpdatedAt = time.Now()
}

