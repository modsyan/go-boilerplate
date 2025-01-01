package dtos

import "company-name/entities"

type GetPageContentBlocksRequest struct {
	Page string `form:"page" validate:"required"`
}

type GetPageContentBlocksResponse struct {
	Blocks []ContentBlockDto
}

func GetPageContentBlocksResponseFromEntity(blocks []*entities.ContentBlocks) *GetPageContentBlocksResponse {
	var blocksDto []ContentBlockDto

	for _, block := range blocks {
		blocksDto = append(blocksDto, ContentBlockDto{
			Page:      block.Key.Page,
			Section:   block.Key.Section,
			Content:   block.Content,
			CreatedAt: block.CreatedAt,
			UpdatedAt: block.UpdatedAt,
		})
	}

	return &GetPageContentBlocksResponse{
		Blocks: blocksDto,
	}
}
