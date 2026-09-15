package create
import ("nfxidentity/modules/auth/domain/identity"; "gorm.io/gorm")
type Handler struct{ db *gorm.DB }
func NewHandler(db *gorm.DB) identity.Create { return &Handler{db: db} }
