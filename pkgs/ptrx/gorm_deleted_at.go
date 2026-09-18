package ptrx

import (
	"time"

	"gorm.io/gorm"
)

// DeletedAtToTimePtr converts GORM soft-delete metadata to domain *time.Time.
// When the row is not deleted (DeletedAt invalid), it returns nil.
func DeletedAtToTimePtr(d gorm.DeletedAt) *time.Time {
	if d.Valid {
		return &d.Time
	}
	return nil
}

// TimePtrToDeletedAt converts domain *time.Time to gorm.DeletedAt for persistence.
// A nil pointer means the row is not soft-deleted.
func TimePtrToDeletedAt(t *time.Time) gorm.DeletedAt {
	if t == nil {
		return gorm.DeletedAt{}
	}
	return gorm.DeletedAt{Time: *t, Valid: true}
}
