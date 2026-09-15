package get
import ("nfxidentity/modules/auth/domain/authorityprofile"; "gorm.io/gorm")
type Handler struct{ db *gorm.DB }
func NewHandler(db *gorm.DB) authorityprofile.Get { return &Handler{db: db} }
