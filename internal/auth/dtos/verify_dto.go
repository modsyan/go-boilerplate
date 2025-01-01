package dtos

//VerifyEmailRequest

type VerifyEmailRequest struct {
	Token string `json:"token" validate:"required"`
}
