package reqdto

type LoginWithEmail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	DeviceID string `json:"device_id"`
}

type LoginWithPhone struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	DeviceID string `json:"device_id"`
}

type SignupWithEmail struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	VerificationCode string `json:"verification_code"`
	Lang             string `json:"lang"`
	DeviceID         string `json:"device_id"`
	SignupPlatform   string `json:"signup_platform"`
}

type SendSignupCode struct {
	Email string `json:"email"`
	Lang  string `json:"lang"`
}

type Refresh struct {
	RefreshToken string `json:"refresh_token"`
	DeviceID     string `json:"device_id"`
}

type Logout struct {
	RefreshToken string `json:"refresh_token"`
}

type SelectProfile struct {
	ProfileID string `json:"profile_id"`
	Kind      string `json:"kind"`
	DeviceID  string `json:"device_id"`
}

type PatchProfileSettings struct {
	LoginNotification *bool `json:"login_notification"`
}

type ConfirmAvatar struct {
	ImageID string `json:"image_id"`
}

type BackgroundImage struct {
	ImageID   string `json:"image_id"`
	SortOrder int    `json:"sort_order"`
}

type ConfirmBackgrounds struct {
	Images []BackgroundImage `json:"images"`
}

type Preference struct {
	Preference string `json:"preference"`
}

type CreateProfile struct {
	DisplayName     string `json:"display_name"`
	ProfileLanguage string `json:"profile_language"`
}

type SearchProfiles struct {
	Query  string `json:"query"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type Email struct {
	Email string `json:"email"`
}

type SendEmailCode struct {
	Lang string `json:"lang"`
}

type VerifyCode struct {
	VerificationCode string `json:"verification_code"`
}

type Phone struct {
	Phone string `json:"phone"`
}

type SendPasswordCode struct {
	Lang string `json:"lang"`
}

type ChangePassword struct {
	CurrentPassword  string `json:"current_password"`
	NewPassword      string `json:"new_password"`
	VerificationCode string `json:"verification_code"`
}

type UpdateAuthorityRoles struct {
	AuthorityRoles []string `json:"authority_roles"`
}
