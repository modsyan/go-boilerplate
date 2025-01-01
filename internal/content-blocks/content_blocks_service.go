package blocks

import (
	"company-name/constants/msgkey"
	"company-name/entities"
	"company-name/internal/content-blocks/dtos"
	"company-name/pkg/errors"
	loc "company-name/pkg/localization"
	"company-name/pkg/validators"
	"context"
)

// IContentBlocksService defines methods for managing content blocks, including creation, update, deletion, and retrieval.
// CreateBlock creates a new content block from the provided DTO and returns the created block response or an error.
// UpdateBlock updates an existing content block using the provided DTO and returns the updated block response or an error.
// DeleteBlock removes a content block identified by the given request DTO and returns an error if the operation fails.
// GetBlock retrieves a single content block based on the page and section specified in the request DTO and returns it or an error.
// GetPage fetches all content blocks for a given page specified in the request DTO and returns them or an error.
type IContentBlocksService interface {
	CreateBlock(ctx context.Context, dto *dtos.CreateContentBlockRequest) (*dtos.CreateContentBlockResponse, error)
	UpdateBlock(ctx context.Context, dto *dtos.UpdateContentBlockRequest) (*dtos.UpdateContentBlockResponse, error)
	DeleteBlock(ctx context.Context, dto *dtos.DeleteBlockRequest) error
	GetBlock(ctx context.Context, dto *dtos.GetContentBlockRequest) (*dtos.GetContentBlockResponse, error)
	GetPage(ctx context.Context, dto *dtos.GetPageContentBlocksRequest) (*dtos.GetPageContentBlocksResponse, error)
}

// ContentBlocksService provides methods for managing content blocks, including creation, update, deletion, and retrieval.
// It utilizes IContentBlockRepository for repository operations and IValidator for input validation.
type ContentBlocksService struct {
	repo      IContentBlockRepository
	validator validators.IValidator
}

// NewContentBlocksService initializes and returns a new IContentBlocksService instance with the provided repository and validator.
func NewContentBlocksService(repo IContentBlockRepository, validator validators.IValidator) IContentBlocksService {
	return &ContentBlocksService{
		repo:      repo,
		validator: validator,
	}
}

// CreateBlock creates a new content block using the provided DTO, validates the input, and stores it in the repository.
func (s *ContentBlocksService) CreateBlock(ctx context.Context, dto *dtos.CreateContentBlockRequest) (*dtos.CreateContentBlockResponse, error) {
	contentBlock := dto.ToEntity()

	if err := contentBlock.Validate(s.validator); err != nil {
		return nil, err
	}

	createdBlock, err := s.repo.CreateContentBlock(ctx, contentBlock)
	if err != nil {
		return nil, errors.InternalServerErrorM(loc.L(msgkey.ErrResourceCreated, msgkey.MsgContentBlockResource), err)
	}

	return dtos.CreateContentBlockResponseFromEntity(createdBlock), nil
}

// UpdateBlock updates an existing content block based on the provided request DTO. Validates input and returns an updated response.
func (s *ContentBlocksService) UpdateBlock(ctx context.Context, dto *dtos.UpdateContentBlockRequest) (*dtos.UpdateContentBlockResponse, error) {
	contentBlock := dto.ToEntity()
	if err := contentBlock.Validate(s.validator); err != nil {
		return nil, err
	}

	updatedBlock, err := s.repo.UpdateContentBlock(ctx, contentBlock)
	if err != nil {
		return nil, errors.InternalServerErrorM(loc.L(msgkey.ErrResourceUpdated, msgkey.MsgContentBlockResource), err)
	}

	return dtos.UpdateContentBlockResponseFromEntity(updatedBlock), nil
}

// DeleteBlock deletes a content block identified by the page and section in the provided request DTO. Returns an error if validation or deletion fails.
func (s *ContentBlocksService) DeleteBlock(ctx context.Context, dto *dtos.DeleteBlockRequest) error {
	blockKey := entities.BlockKey{
		Page:    dto.Page,
		Section: dto.Section,
	}

	if _, err := s.repo.GetContentBlock(ctx, blockKey); err == nil {
		return errors.NotFoundM(loc.L(msgkey.ErrResourceNotFound, msgkey.MsgContentBlockResource), err)
	}

	if err := s.repo.DeleteContentBlock(ctx, blockKey); err != nil {
		return errors.InternalServerErrorM(loc.L(msgkey.ErrResourceDeleted, msgkey.MsgContentBlockResource), err)
	}

	return nil
}

// GetPage retrieves all content blocks associated with a specified page from the repository. Returns an error if retrieval fails.
func (s *ContentBlocksService) GetPage(ctx context.Context, dto *dtos.GetPageContentBlocksRequest) (*dtos.GetPageContentBlocksResponse, error) {
	contentBlocks, err := s.repo.GetPageContentBlocks(ctx, dto.Page)
	if err != nil {
		return nil, errors.InternalServerErrorM(loc.L(msgkey.ErrResourceFetched, msgkey.MsgContentBlockResource), err)
	}

	return dtos.GetPageContentBlocksResponseFromEntity(contentBlocks), nil
}

// GetBlock retrieves a content block from the repository using the specified page and section identifiers.
func (s *ContentBlocksService) GetBlock(ctx context.Context, dto *dtos.GetContentBlockRequest) (*dtos.GetContentBlockResponse, error) {
	key := entities.BlockKey{
		Page:    dto.Page,
		Section: dto.Section,
	}

	contentBlocks, err := s.repo.GetContentBlock(ctx, key)
	if err != nil {
		return nil, errors.InternalServerErrorM(loc.L(msgkey.ErrResourceFetched, msgkey.MsgContentBlockResource), err)
	}

	return dtos.GetContentBlockResponseFromEntity(contentBlocks), nil
}
