package blocks

import (
	"company-name/constants"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"time"

	"company-name/entities"
	"company-name/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IContentBlockRepository interface {
	CreateContentBlock(ctx context.Context, contentBlock *entities.ContentBlocks) (*entities.ContentBlocks, error)
	UpdateContentBlock(ctx context.Context, contentBlock *entities.ContentBlocks) (*entities.ContentBlocks, error)
	DeleteContentBlock(ctx context.Context, key entities.BlockKey) error
	GetPageContentBlocks(ctx context.Context, page string) ([]*entities.ContentBlocks, error)
	GetContentBlock(ctx context.Context, key entities.BlockKey) (*entities.ContentBlocks, error)
}

type ContentBlockRepository struct {
	db database.IDatabase
}

func NewContentBlockRepository(db database.IDatabase) IContentBlockRepository {
	return &ContentBlockRepository{
		db: db,
	}
}

func (r *ContentBlockRepository) CreateContentBlock(ctx context.Context, contentBlock *entities.ContentBlocks) (*entities.ContentBlocks, error) {
	contentBlock.ID = primitive.NewObjectID()
	contentBlock.CreatedAt = time.Now()
	contentBlock.UpdatedAt = time.Now()

	if err := r.db.Create(ctx, constants.DbContentBlocksCollection, contentBlock); err != nil {
		return nil, err
	}

	return contentBlock, nil
}

func (r *ContentBlockRepository) UpdateContentBlock(ctx context.Context, contentBlock *entities.ContentBlocks) (*entities.ContentBlocks, error) {
	contentBlock.UpdatedAt = time.Now()

	filter := bson.M{"key.page": contentBlock.Key.Page, "key.section": contentBlock.Key.Section}
	update := bson.M{
		"$set": contentBlock,
	}

	if err := r.db.Update(ctx, constants.DbContentBlocksCollection, filter, update); err != nil {
		return nil, err
	}

	return contentBlock, nil
}

func (r *ContentBlockRepository) GetPageContentBlocks(ctx context.Context, page string) ([]*entities.ContentBlocks, error) {
	filter := bson.M{"key.page": page}
	var contentBlocks []*entities.ContentBlocks

	if err := r.db.Find(ctx, constants.DbContentBlocksCollection, filter, &contentBlocks); err != nil {
		return nil, err
	}

	return contentBlocks, nil
}

func (r *ContentBlockRepository) GetContentBlock(ctx context.Context, key entities.BlockKey) (*entities.ContentBlocks, error) {
	filter := bson.M{"key.page": key.Page, "key.section": key.Section}
	var contentBlock entities.ContentBlocks

	if err := r.db.FindOne(ctx, constants.DbContentBlocksCollection, filter, &contentBlock); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &contentBlock, nil
}

func (r *ContentBlockRepository) DeleteContentBlock(ctx context.Context, key entities.BlockKey) error {
	filter := bson.M{"key.page": key.Page, "key.section": key.Section}

	if err := r.db.Delete(ctx, constants.DbContentBlocksCollection, filter); err != nil {
		return err
	}

	return nil
}
