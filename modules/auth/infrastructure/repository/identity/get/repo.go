package get
import ("nfxidentity/modules/auth/domain/identity"; "gorm.io/gorm")
type Handler struct{ db *gorm.DB }
func NewHandler(db *gorm.DB) identity.Get { return &Handler{db: db} }
