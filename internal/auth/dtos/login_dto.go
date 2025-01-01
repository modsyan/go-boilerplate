package dtos

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token                    string `json:"token"`
	RefreshToken             string `json:"refresh_token,omitempty"`
	ExpirationInMilliseconds int64  `json:"expiration_in_milliseconds"`
}
