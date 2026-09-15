package forgerprofile
import ("nfxidentity/modules/auth/domain/forgerprofile"; "nfxidentity/modules/auth/infrastructure/repository/forgerprofile/create"; "nfxidentity/modules/auth/infrastructure/repository/forgerprofile/get"; "nfxidentity/modules/auth/infrastructure/repository/forgerprofile/update"; "gorm.io/gorm")
func NewRepo(db *gorm.DB) *forgerprofile.Repo {
	return &forgerprofile.Repo{Create: create.NewHandler(db), Get: get.NewHandler(db), Update: update.NewHandler(db)}
}
