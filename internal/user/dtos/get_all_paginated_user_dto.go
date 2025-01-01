package dtos

import (
	"company-name/entities"
	"company-name/pkg/models/results"
)

type GetPaginatedUsersRequest struct {
	results.PaginationRequest
}

type GetPaginatedUsersResponse struct {
	Users []UserDto `json:"users"`
	Total int       `json:"total"`
}

func GetPaginatedUsersResponseFromEntity(users []*entities.User, totalCount int) *GetPaginatedUsersResponse {
	userDtos := make([]UserDto, len(users))
	for _, user := range users {
		userDtos = append(userDtos, *UserDtoFromEntity(user))
	}

	return &GetPaginatedUsersResponse{Users: userDtos, Total: totalCount}
}
