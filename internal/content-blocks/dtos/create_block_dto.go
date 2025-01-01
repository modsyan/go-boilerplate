package dtos

import (
	"company-name/entities"
	"time"
)

type CreateContentBlockRequest struct {
	Page    string `json:"page" validate:"required"`
	Section string `json:"section" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (req *CreateContentBlockRequest) ToEntity() *entities.ContentBlocks {
	return &entities.ContentBlocks{
		Key: entities.BlockKey{
			Page:    req.Page,
			Section: req.Section,
		},
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type CreateContentBlockResponse struct {
	ContentBlockDto
}

func CreateContentBlockResponseFromEntity(entity *entities.ContentBlocks) *CreateContentBlockResponse {
	return &CreateContentBlockResponse{
		ContentBlockDto: *ContentBlockDtoFromEntity(entity),
	}
}
