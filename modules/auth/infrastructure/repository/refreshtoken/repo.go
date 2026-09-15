package refreshtoken
import ("nfxidentity/modules/auth/domain/refreshtoken"; "nfxidentity/modules/auth/infrastructure/repository/refreshtoken/create"; "nfxidentity/modules/auth/infrastructure/repository/refreshtoken/get"; "nfxidentity/modules/auth/infrastructure/repository/refreshtoken/update"; "gorm.io/gorm")
func NewRepo(db *gorm.DB) *refreshtoken.Repo {
	return &refreshtoken.Repo{Create: create.NewHandler(db), Get: get.NewHandler(db), Update: update.NewHandler(db)}
}
