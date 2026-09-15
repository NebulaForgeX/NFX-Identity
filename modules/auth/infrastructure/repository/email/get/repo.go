package get
import ("nfxidentity/modules/auth/domain/email"; "gorm.io/gorm")
type Handler struct{ db *gorm.DB }
func NewHandler(db *gorm.DB) email.Get { return &Handler{db: db} }
