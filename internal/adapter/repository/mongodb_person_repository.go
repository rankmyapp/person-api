package repository

import (
	"context"
	"time"

	"github.com/rpuglielli/person-api/internal/domain/person/entity"
	"github.com/rpuglielli/person-api/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBPersonRepository struct {
	collection *mongo.Collection
}

func NewMongoDBPersonRepository(db *mongo.Database) *MongoDBPersonRepository {
	return &MongoDBPersonRepository{
		collection: db.Collection("persons"),
	}
}

func (r *MongoDBPersonRepository) Create(ctx context.Context, person *entity.Person) error {
	result, err := r.collection.InsertOne(ctx, person)
	if err != nil {
		return errors.NewInternalError("Failed to create person")
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	person.ID = id

	return nil
}

func (r *MongoDBPersonRepository) Update(ctx context.Context, person *entity.Person) error {
	objectID, err := primitive.ObjectIDFromHex(person.ID)
	if err != nil {
		return errors.NewValidationError("Invalid ID")
	}

	person.Updated = time.Now()

	updateData, err := bson.Marshal(person)
	if err != nil {
		return err
	}

	var updateDoc bson.M
	err = bson.Unmarshal(updateData, &updateDoc)
	if err != nil {
		return err
	}
	delete(updateDoc, "_id")
	delete(updateDoc, "created")

	update := bson.M{
		"$set": updateDoc,
	}

	_, err = r.collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, update)
	if err != nil {
		return err
	}

	var updatedPerson entity.Person
	err = r.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&updatedPerson)
	if err != nil {
		return err
	}

	return nil
}

func (r *MongoDBPersonRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.NewValidationError("Invalid ID")
	}

	filter := bson.M{"_id": objectID}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.NewInternalError("Failed to delete person")
	}

	if result.DeletedCount == 0 {
		return errors.NewNotFoundError("Person not found")
	}

	return nil
}

func (r *MongoDBPersonRepository) FindByID(ctx context.Context, id string) (*entity.Person, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewValidationError("Invalid ID")
	}

	filter := bson.M{"_id": objectID}

	var person entity.Person
	err = r.collection.FindOne(ctx, filter).Decode(&person)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.NewNotFoundError("Person not found")
		}
		return nil, errors.NewInternalError("Failed to find person")
	}

	return &person, nil
}

func (r *MongoDBPersonRepository) FindAll(ctx context.Context, page, pageSize int) ([]*entity.Person, int64, error) {
	skip := (page - 1) * pageSize

	findOptions := options.Find()
	findOptions.SetLimit(int64(pageSize))
	findOptions.SetSkip(int64(skip))

	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, 0, errors.NewInternalError("Failed to fetch persons")
	}
	defer cursor.Close(ctx)

	var persons []*entity.Person
	for cursor.Next(ctx) {
		var person entity.Person
		if err := cursor.Decode(&person); err != nil {
			return nil, 0, errors.NewInternalError("Failed to decode person")
		}
		persons = append(persons, &person)
	}

	if err := cursor.Err(); err != nil {
		return nil, 0, errors.NewInternalError("Cursor error")
	}

	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, errors.NewInternalError("Failed to count persons")
	}

	return persons, total, nil
}

func (r *MongoDBPersonRepository) FindByEmail(ctx context.Context, email string) (*entity.Person, error) {
	filter := bson.M{"email": email}

	var person entity.Person
	err := r.collection.FindOne(ctx, filter).Decode(&person)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.NewNotFoundError("Person not found")
		}
		return nil, errors.NewInternalError("Failed to find person by email")
	}

	return &person, nil
}
