package dtos

import (
	"company-name/entities"
	"time"
)

type UpdateContentBlockRequest struct {
	Page    string `json:"page" validate:"required"`
	Section string `json:"section" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (req *UpdateContentBlockRequest) ToEntity() *entities.ContentBlocks {
	return &entities.ContentBlocks{
		Key: entities.BlockKey{
			Page:    req.Page,
			Section: req.Section,
		},
		Content:   req.Content,
		UpdatedAt: time.Now(),
	}
}

type UpdateContentBlockResponse struct {
	ContentBlockDto
}

func UpdateContentBlockResponseFromEntity(block *entities.ContentBlocks) *UpdateContentBlockResponse {
	return &UpdateContentBlockResponse{
		ContentBlockDto: *ContentBlockDtoFromEntity(block),
	}
}
