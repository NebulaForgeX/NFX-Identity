package update
import ("nfxidentity/modules/auth/domain/email"; "gorm.io/gorm")
type Handler struct{ db *gorm.DB }
func NewHandler(db *gorm.DB) email.Update { return &Handler{db: db} }
