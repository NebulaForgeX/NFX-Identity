package get
import ("nfxidentity/modules/auth/domain/refreshtoken"; "gorm.io/gorm")
type Handler struct{ db *gorm.DB }
func NewHandler(db *gorm.DB) refreshtoken.Get { return &Handler{db: db} }
