package dtos

import "company-name/entities"

type GetContentBlockRequest struct {
	Page    string `form:"page" validate:"required"`
	Section string `form:"section" validate:"required"`
}

type GetContentBlockResponse struct {
	ContentBlockDto
}

func GetContentBlockResponseFromEntity(entity *entities.ContentBlocks) *GetContentBlockResponse {
	return &GetContentBlockResponse{
		ContentBlockDto: *ContentBlockDtoFromEntity(entity),
	}
}
