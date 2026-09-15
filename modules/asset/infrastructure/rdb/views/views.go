package views

import (
	"time"

	"github.com/google/uuid"
)

type ImagesActiveView struct {
	ID         uuid.UUID  `gorm:"column:id"`
	FilePath   string     `gorm:"column:file_path"`
	FileName   string     `gorm:"column:file_name"`
	FileSize   int64      `gorm:"column:file_size"`
	MimeType   string     `gorm:"column:mime_type"`
	Width      *int       `gorm:"column:width"`
	Height     *int       `gorm:"column:height"`
	AltText    *string    `gorm:"column:alt_text"`
	UploaderID uuid.UUID  `gorm:"column:uploader_id"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at"`
}

func (ImagesActiveView) TableName() string { return "asset.ImagesActiveView" }

type FilesActiveView struct {
	ID         uuid.UUID `gorm:"column:id"`
	FilePath   string    `gorm:"column:file_path"`
	FileName   string    `gorm:"column:file_name"`
	FileSize   int64     `gorm:"column:file_size"`
	MimeType   string    `gorm:"column:mime_type"`
	UploaderID uuid.UUID `gorm:"column:uploader_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (FilesActiveView) TableName() string { return "asset.FilesActiveView" }

type VideosActiveView struct {
	ID              uuid.UUID `gorm:"column:id"`
	FilePath        string    `gorm:"column:file_path"`
	FileName        string    `gorm:"column:file_name"`
	FileSize        int64     `gorm:"column:file_size"`
	MimeType        string    `gorm:"column:mime_type"`
	DurationSeconds *float64  `gorm:"column:duration_seconds"`
	Width           *int      `gorm:"column:width"`
	Height          *int      `gorm:"column:height"`
	UploaderID      uuid.UUID `gorm:"column:uploader_id"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}

func (VideosActiveView) TableName() string { return "asset.VideosActiveView" }

type AudiosActiveView struct {
	ID              uuid.UUID `gorm:"column:id"`
	FilePath        string    `gorm:"column:file_path"`
	FileName        string    `gorm:"column:file_name"`
	FileSize        int64     `gorm:"column:file_size"`
	MimeType        string    `gorm:"column:mime_type"`
	DurationSeconds *float64  `gorm:"column:duration_seconds"`
	UploaderID      uuid.UUID `gorm:"column:uploader_id"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}

func (AudiosActiveView) TableName() string { return "asset.AudiosActiveView" }
