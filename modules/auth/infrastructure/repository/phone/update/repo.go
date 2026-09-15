package update
import ("nfxidentity/modules/auth/domain/phone"; "gorm.io/gorm")
type Handler struct{ db *gorm.DB }
func NewHandler(db *gorm.DB) phone.Update { return &Handler{db: db} }
