package reqdto

type SendVerificationCodeRequestDTO struct {
	Email string `json:"email" validate:"required,email"`
}
