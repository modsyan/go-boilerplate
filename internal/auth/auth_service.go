package auth

import (
	"company-name/configs"
	"company-name/constants"
	"company-name/entities"
	"company-name/internal/auth/dtos"
	"company-name/pkg/email"
	errors2 "company-name/pkg/errors"
	"company-name/pkg/hasher"
	"company-name/pkg/jwttoken"
	"company-name/pkg/localization"
	"company-name/pkg/validators"
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

type IAuthService interface {
	GetToken(ctx context.Context, request *dtos.LoginRequest) (*dtos.LoginResponse, error)
	Register(ctx context.Context, request *dtos.RegisterRequest) (*dtos.RegisterResponse, error)
	VerifyEmail(ctx context.Context, req *dtos.VerifyEmailRequest) error
}

type Service struct {
	repository   IAuthRepository
	config       *configs.Config
	validator    validators.IValidator
	emailService email.IEmailService
}

func NewAuthService(
	repository IAuthRepository,
	config *configs.Config,
	validator validators.IValidator,
	emailService email.IEmailService,
) IAuthService {
	return &Service{
		repository:   repository,
		config:       config,
		validator:    validator,
		emailService: emailService,
	}
}

func (s *Service) GetToken(ctx context.Context, req *dtos.LoginRequest) (*dtos.LoginResponse, error) {
	// Fetch user by email
	user, err := s.repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors2.UnauthorizedM(localization.L("error_invalid_credentials"), err)
	}

	if err := hasher.CheckPasswordHash(req.Password, user.HashedPassword); err != nil {
		return nil, err
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, errors.New(localization.L("error_token_generation"))
	}

	response := &dtos.LoginResponse{
		Token:                    token,
		ExpirationInMilliseconds: s.config.JWT.Expiration,
	}

	return response, nil
}

func (s *Service) Register(ctx context.Context, req *dtos.RegisterRequest) (*dtos.RegisterResponse, error) {

	userExisted, err := s.repository.GetUserByEmail(ctx, req.Email)
	//if err != nil {
	//	return nil, errors2.BadRequestM("error_invalid_credentials", err)
	//}
	if err != nil && userExisted != nil {
		return nil, errors2.BadRequestM("error_email_already_used", err)
	}

	hashedPassword, err := hasher.HashPassword(req.Password)
	if err != nil {
		return nil, errors2.InternalServerErrorM("error_password_hashing", err)
	}

	// Convert DTO to Entity
	createdUser := req.ToEntity(hashedPassword)

	// Persist the user in the database
	user, err := s.repository.CreateUser(ctx, &createdUser)
	if err != nil {
		return nil, errors2.InternalServerErrorM("error_user_creation", err)
	}

	// Convert Entity to Response DTO
	response := &dtos.RegisterResponse{}
	response.FromEntity(user, s.generateVerificationLink(user.ID.Hex()))

	return response, nil
}

func (s *Service) VerifyEmail(ctx context.Context, req *dtos.VerifyEmailRequest) error {
	// Validate the token
	claims, err := jwttoken.ValidateAccessToken(req.Token)
	if err != nil {
		return errors2.BadRequestM("error_invalid_token", err)
	}

	// Fetch user by ID
	user, err := s.repository.GetUserById(ctx, claims["sub"].(string))
	if err != nil {
		return errors2.BadRequestM("error_invalid_user", err)
	}

	// Update the user's email verification status
	user.Status = constants.UserStatusActivated
	if err := s.repository.UpdateUser(ctx, user); err != nil {
		return errors2.InternalServerErrorM("error_user_update", err)
	}

	return nil
}

func (s *Service) generateToken(user *entities.User) (string, error) {
	jwtExp := s.config.JWT.Expiration
	jwtSec := s.config.JWT.Secret

	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"exp":   jwtExp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSec))
}

func (s *Service) generateVerificationLink(userID string) string {
	token := jwttoken.GenerateAccessToken(userID)

	baseURL := s.config.App.VerificationUrl

	// Construct the full verification link with token
	verificationLink := fmt.Sprintf("%s?token=%s", baseURL, token)

	return verificationLink
}
