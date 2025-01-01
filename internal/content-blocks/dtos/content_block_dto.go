package dtos

import (
	"company-name/entities"
	"time"
)

type ContentBlockDto struct {
	Page      string    `json:"page"`
	Section   string    `json:"section"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ContentBlockDtoFromEntity(entity *entities.ContentBlocks) *ContentBlockDto {
	return &ContentBlockDto{
		Page:      entity.Key.Page,
		Section:   entity.Key.Section,
		Content:   entity.Content,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func (dto *ContentBlockDto) ToEntity() *entities.ContentBlocks {
	return &entities.ContentBlocks{
		Key: entities.BlockKey{
			Page:    dto.Page,
			Section: dto.Section,
		},
		Content:   dto.Content,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}
