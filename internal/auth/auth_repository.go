package auth

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"time"

	"company-name/constants"
	"company-name/entities"
	"company-name/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAuthRepository interface {
	CreateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetUserById(ctx context.Context, id string) (*entities.User, error)
	UpdatePassword(ctx context.Context, id string, hashedPassword string) error
	UpdateUser(ctx context.Context, user *entities.User) error
}

type Repository struct {
	db database.IDatabase
}

func NewAuthRepository(db database.IDatabase) IAuthRepository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	if err := r.db.Create(ctx, constants.DbUsersCollection, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	filter := bson.M{"email": email}
	err := r.db.FindOne(ctx, constants.DbUsersCollection, filter, &user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetUserById(ctx context.Context, id string) (*entities.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	var user entities.User
	filter := bson.M{"_id": objectId}
	err = r.db.FindOne(ctx, constants.DbUsersCollection, filter, &user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *Repository) UpdatePassword(ctx context.Context, id string, hashedPassword string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user ID")
	}

	update := bson.M{"$set": bson.M{"password": hashedPassword, "updatedAt": time.Now()}}
	filter := bson.M{"_id": objectId}

	err = r.db.Update(ctx, constants.DbUsersCollection, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateUser(ctx context.Context, user *entities.User) error {
	user.UpdatedAt = time.Now()

	filter := bson.M{"_id": user.ID}
	update := bson.M{
		"$set": user,
	}

	if err := r.db.Update(ctx, constants.DbUsersCollection, filter, update); err != nil {
		return err
	}

	return nil
}
