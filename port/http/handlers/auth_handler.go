package handlers

import (
	"company-name/constants/msgkey"
	"company-name/internal/auth"
	"company-name/internal/auth/dtos"
	"company-name/pkg/errors"
	loc "company-name/pkg/localization"
	"company-name/pkg/responses"
	"company-name/pkg/validators"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service   auth.IAuthService
	validator validators.IValidator
}

func NewAuthHandler(service auth.IAuthService, validator validators.IValidator) *AuthHandler {
	return &AuthHandler{
		service:   service,
		validator: validator,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginRequest dtos.LoginRequest

	if !validators.BindJsonAndValidateRequest(c, &loginRequest, h.validator) {
		return
	}

	result, err := h.service.GetToken(c, &loginRequest)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	responses.Ok(c, loc.L(msgkey.MsgLoginSuccessful), result)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var registerRequest dtos.RegisterRequest

	if !validators.BindJsonAndValidateRequest(c, &registerRequest, h.validator) {
		return
	}

	result, err := h.service.Register(c, &registerRequest)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	responses.Created(c, loc.L(msgkey.MsgUserRegistered), result)
}

// verify-email
func (h *AuthHandler) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	var verifyEmailRequest = dtos.VerifyEmailRequest{Token: token}

	if !validators.ValidateRequestOnly(c, &verifyEmailRequest, h.validator) {
		return
	}

	err := h.service.VerifyEmail(c, &verifyEmailRequest)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	responses.Ok(c, loc.L(msgkey.MsgEmailVerified), nil)
}
