package get
import ("nfxidentity/modules/auth/domain/phone"; "gorm.io/gorm")
type Handler struct{ db *gorm.DB }
func NewHandler(db *gorm.DB) phone.Get { return &Handler{db: db} }
