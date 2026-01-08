package images

import (
	"time"

	"github.com/google/uuid"
)

type Image struct {
	state ImageState
}

type ImageState struct {
	ID              uuid.UUID
	TypeID          *uuid.UUID
	UserID          *uuid.UUID
	TenantID        *uuid.UUID
	AppID           *uuid.UUID
	SourceDomain    *string
	Filename        string
	OriginalFilename string
	MimeType        string
	Size            int64
	Width           *int
	Height          *int
	StoragePath     string
	URL             *string
	IsPublic        bool
	Metadata        map[string]interface{}
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

func (i *Image) ID() uuid.UUID                    { return i.state.ID }
func (i *Image) TypeID() *uuid.UUID               { return i.state.TypeID }
func (i *Image) UserID() *uuid.UUID               { return i.state.UserID }
func (i *Image) TenantID() *uuid.UUID             { return i.state.TenantID }
func (i *Image) AppID() *uuid.UUID                { return i.state.AppID }
func (i *Image) SourceDomain() *string            { return i.state.SourceDomain }
func (i *Image) Filename() string                 { return i.state.Filename }
func (i *Image) OriginalFilename() string         { return i.state.OriginalFilename }
func (i *Image) MimeType() string                 { return i.state.MimeType }
func (i *Image) Size() int64                      { return i.state.Size }
func (i *Image) Width() *int                      { return i.state.Width }
func (i *Image) Height() *int                     { return i.state.Height }
func (i *Image) StoragePath() string              { return i.state.StoragePath }
func (i *Image) URL() *string                     { return i.state.URL }
func (i *Image) IsPublic() bool                   { return i.state.IsPublic }
func (i *Image) Metadata() map[string]interface{} { return i.state.Metadata }
func (i *Image) CreatedAt() time.Time             { return i.state.CreatedAt }
func (i *Image) UpdatedAt() time.Time             { return i.state.UpdatedAt }
func (i *Image) DeletedAt() *time.Time            { return i.state.DeletedAt }
