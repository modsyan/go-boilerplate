package user

import (
	"company-name/constants/msgkey"
	"company-name/internal/user/dtos"
	"company-name/pkg/errors"
	"company-name/pkg/hasher"
	loc "company-name/pkg/localization"
	"company-name/pkg/validators"
	"context"
)

type IUserService interface {
	CreateUser(ctx context.Context, req *dtos.CreateUserRequest) (*dtos.CreateUserResponse, error)
	UpdateUser(ctx context.Context, req *dtos.UpdateUserRequest) (*dtos.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req *dtos.DeleteUserRequest) error
	GetUserDetailsById(ctx context.Context, req *dtos.GetUserDetailsRequest) (*dtos.GetUserDetailsResponse, error)
	GetPaginatedUsers(ctx context.Context, dto *dtos.GetPaginatedUsersRequest) (*dtos.GetPaginatedUsersResponse, error)
}

type Service struct {
	repo      IUserRepository
	validator validators.IValidator
}

func NewUserService(repo IUserRepository, validator validators.IValidator) IUserService {
	return &Service{repo, validator}
}

func (s *Service) CreateUser(ctx context.Context, req *dtos.CreateUserRequest) (*dtos.CreateUserResponse, error) {
	user, password := req.ToEntity()

	_, err := s.repo.FindByEmail(ctx, req.Email)
	if err == nil {
		return nil, errors.Conflict(err)
	}

	hashedPassword, err := hasher.HashPassword(password)
	if err != nil {
		return nil, errors.InternalServerError(err)
	}
	user.HashedPassword = hashedPassword

	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, errors.InternalServerErrorM("Failed to create user", err)
	}

	response := dtos.CreateUserResponseFromEntity(user)
	return response, nil
	panic("not implemented")
}

func (s *Service) UpdateUser(ctx context.Context, req *dtos.UpdateUserRequest) (*dtos.UpdateUserResponse, error) {
	user, password, err := req.ToEntity()
	if err != nil {
		return nil, err
	}

	if password != "" {
		hashedPassword, err := hasher.HashPassword(password)
		if err != nil {
			return nil, errors.InternalServerError(err)
		}
		user.HashedPassword = hashedPassword
	}

	err = s.repo.Update(ctx, user)
	if err != nil {
		return nil, errors.InternalServerErrorM("Failed to update user", err)
	}

	return dtos.UpdateUserResponseFromEntity(user), nil
}

func (s *Service) DeleteUser(ctx context.Context, req *dtos.DeleteUserRequest) error {
	if _, err := s.repo.FindByID(ctx, req.ID); err != nil {
		return errors.NotFound(err)
	}

	if err := s.repo.Delete(ctx, req.ID); err != nil {
		return errors.InternalServerErrorM(loc.L(msgkey.ErrResourceDeleted, msgkey.MsgUserResource), err)
	}
	return nil
}

func (s *Service) GetUserDetailsById(ctx context.Context, req *dtos.GetUserDetailsRequest) (*dtos.GetUserDetailsResponse, error) {
	user, err := s.repo.FindByID(ctx, req.ID)
	if err != nil {
		return nil, errors.InternalServerErrorM("Failed to get user details", err)
	}

	return dtos.GetUserDetailsResponseFromEntity(user), nil
}

func (s *Service) GetPaginatedUsers(ctx context.Context, dto *dtos.GetPaginatedUsersRequest) (*dtos.GetPaginatedUsersResponse, error) {
	users, totalCount, err := s.repo.FindAllPaginated(ctx, dto.FilterSearch, dto.Page, dto.PageSize, dto.SortBy, dto.SortOrder)
	if err != nil {
		return nil, errors.InternalServerErrorM("Failed to get paginated users", err)
	}

	return dtos.GetPaginatedUsersResponseFromEntity(users, int(totalCount)), nil
}
