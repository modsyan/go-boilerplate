package user

import (
	"context"
	"errors"
	"log"
	"time"

	"company-name/constants"
	"company-name/entities"
	"company-name/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// IUserRepository defines the interface for user repository
type IUserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	Update(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id string) error
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	FindByID(ctx context.Context, id string) (*entities.User, error)
	FindAll(ctx context.Context) ([]*entities.User, error)
	FindAllPaginated(ctx context.Context, filter string, page, pageSize int, sortBy, sortOrder string) ([]*entities.User, int64, error)
}

type Repository struct {
	db database.IDatabase
}

// NewUserRepository initializes a new UserRepository
func NewUserRepository(db database.IDatabase) IUserRepository {
	return &Repository{db: db}
}

// Create adds a new user to the database
func (r *Repository) Create(ctx context.Context, user *entities.User) error {
	if err := r.db.Create(ctx, constants.DbUsersCollection, user); err != nil {
		return err
	}
	return nil
}

// Update modifies an existing user's details in the database
func (r *Repository) Update(ctx context.Context, user *entities.User) error {
	user.UpdatedAt = time.Now()
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}

	if err := r.db.Update(ctx, constants.DbUsersCollection, filter, update); err != nil {
		return err
	}
	return nil
}

// Delete removes a user from the database by ID
func (r *Repository) Delete(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user ID")
	}

	filter := bson.M{"_id": objectId}
	if err := r.db.Delete(ctx, constants.DbUsersCollection, filter); err != nil {
		return err
	}
	return nil
}

// FindByEmail retrieves a user from the database by their email
func (r *Repository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
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

// FindByID retrieves a user from the database by their ID
func (r *Repository) FindByID(ctx context.Context, id string) (*entities.User, error) {
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

func (r *Repository) FindAll(ctx context.Context) ([]*entities.User, error) {
	var users []*entities.User

	// Query all users from the collection
	cursor, err := r.db.GetDB().Collection(constants.DbUsersCollection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := cursor.Close(ctx); closeErr != nil {
			log.Printf("Error closing cursor: %v", closeErr)
		}
	}()

	// Iterate through the cursor
	for cursor.Next(ctx) {
		var user entities.User
		if err := cursor.Decode(&user); err != nil {
			log.Printf("Error decoding user document: %v", err)
			continue
		}
		users = append(users, &user)
	}

	// Check if cursor encountered any error during iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// FindAllPaginated retrieves a paginated list of users from the database
func (r *Repository) FindAllPaginated(ctx context.Context, filter string, page, pageSize int, sortBy, sortOrder string) ([]*entities.User, int64, error) {
	searchFilter := bson.M{}
	if filter != "" {
		searchFilter["$or"] = bson.A{
			bson.M{"firstName": bson.M{"$regex": filter, "$options": "i"}},
			bson.M{"lastName": bson.M{"$regex": filter, "$options": "i"}},
			bson.M{"email": bson.M{"$regex": filter, "$options": "i"}},
			bson.M{"phoneNumber": bson.M{"$regex": filter, "$options": "i"}},
		}
	}

	offset := (page - 1) * pageSize

	var users []*entities.User
	totalCount, err := r.db.Count(ctx, constants.DbUsersCollection, searchFilter)
	if err != nil {
		return nil, 0, err
	}

	if err := r.db.FindWithPagination(ctx, constants.DbUsersCollection, searchFilter, sortBy, sortOrder, int64(offset), int64(pageSize), &users); err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}
