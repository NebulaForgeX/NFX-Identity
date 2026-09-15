package rdb

import "time"

type Image struct {
	ID         string     `gorm:"column:id;type:uuid;primaryKey"`
	FilePath   string     `gorm:"column:file_path"`
	FileName   string     `gorm:"column:file_name"`
	FileSize   int64      `gorm:"column:file_size"`
	MimeType   string     `gorm:"column:mime_type"`
	Width      *int       `gorm:"column:width"`
	Height     *int       `gorm:"column:height"`
	AltText    *string    `gorm:"column:alt_text"`
	UploaderID string     `gorm:"column:uploader_id;type:uuid"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at"`
}

func (Image) TableName() string { return "asset.Images" }

type File struct {
	ID         string     `gorm:"column:id;type:uuid;primaryKey"`
	FilePath   string     `gorm:"column:file_path"`
	FileName   string     `gorm:"column:file_name"`
	FileSize   int64      `gorm:"column:file_size"`
	MimeType   string     `gorm:"column:mime_type"`
	UploaderID string     `gorm:"column:uploader_id;type:uuid"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at"`
}

func (File) TableName() string { return "asset.Files" }

type Video struct {
	ID              string     `gorm:"column:id;type:uuid;primaryKey"`
	FilePath        string     `gorm:"column:file_path"`
	FileName        string     `gorm:"column:file_name"`
	FileSize        int64      `gorm:"column:file_size"`
	MimeType        string     `gorm:"column:mime_type"`
	DurationSeconds *float64   `gorm:"column:duration_seconds"`
	Width           *int       `gorm:"column:width"`
	Height          *int       `gorm:"column:height"`
	UploaderID      string     `gorm:"column:uploader_id;type:uuid"`
	CreatedAt       time.Time  `gorm:"column:created_at"`
	UpdatedAt       time.Time  `gorm:"column:updated_at"`
	DeletedAt       *time.Time `gorm:"column:deleted_at"`
}

func (Video) TableName() string { return "asset.Videos" }

type Audio struct {
	ID              string     `gorm:"column:id;type:uuid;primaryKey"`
	FilePath        string     `gorm:"column:file_path"`
	FileName        string     `gorm:"column:file_name"`
	FileSize        int64      `gorm:"column:file_size"`
	MimeType        string     `gorm:"column:mime_type"`
	DurationSeconds *float64   `gorm:"column:duration_seconds"`
	UploaderID      string     `gorm:"column:uploader_id;type:uuid"`
	CreatedAt       time.Time  `gorm:"column:created_at"`
	UpdatedAt       time.Time  `gorm:"column:updated_at"`
	DeletedAt       *time.Time `gorm:"column:deleted_at"`
}

func (Audio) TableName() string { return "asset.Audios" }
