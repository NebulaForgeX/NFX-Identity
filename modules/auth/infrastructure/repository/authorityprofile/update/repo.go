package update
import ("nfxidentity/modules/auth/domain/authorityprofile"; "gorm.io/gorm")
type Handler struct{ db *gorm.DB }
func NewHandler(db *gorm.DB) authorityprofile.Update { return &Handler{db: db} }
