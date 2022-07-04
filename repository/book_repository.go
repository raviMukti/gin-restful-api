package repository

import (
	"context"
	"os"

	"github.com/raviMukti/gin-restful-api/helper"
	"github.com/raviMukti/gin-restful-api/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository interface {
	Save(ctx context.Context, book domain.Book)
	Update(ctx context.Context, book domain.Book) domain.Book
	Delete(ctx context.Context, bookId string) error
	FindBydId(ctx context.Context, bookId string) (domain.Book, error)
	FindAll(ctx context.Context) []domain.Book
}

type BookRepositoryImpl struct {
	MongoClient *mongo.Client
}

func NewBookRepository(mongoClient *mongo.Client) BookRepository {
	return &BookRepositoryImpl{
		MongoClient: mongoClient,
	}
}

// Save implements BookRepository
func (repository *BookRepositoryImpl) Save(ctx context.Context, book domain.Book) {
	// Get Collection Names
	collection := repository.MongoClient.Database(os.Getenv("MONGO_INSTANCE")).Collection("book")
	// Get Inserted Object
	_, errInsert := collection.InsertOne(ctx, book)
	helper.PanicIfError(errInsert)
}

// FindBydId implements BookRepository
func (repository *BookRepositoryImpl) FindBydId(ctx context.Context, bookId string) (domain.Book, error) {
	var result domain.Book
	// Get Collection Names
	collection := repository.MongoClient.Database(os.Getenv("MONGO_INSTANCE")).Collection("book")
	// Find The Object
	objId, _ := primitive.ObjectIDFromHex(bookId)
	errFind := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&result)
	// Return Value
	return result, errFind
}

// FindAll implements BookRepository
func (repository *BookRepositoryImpl) FindAll(ctx context.Context) []domain.Book {
	// Get Collection Names
	collection := repository.MongoClient.Database(os.Getenv("MONGO_INSTANCE")).Collection("book")
	// Find The Object
	cursor, err := collection.Find(ctx, bson.M{})
	helper.PanicIfError(err)
	defer cursor.Close(ctx)
	// Return Value
	result := make([]domain.Book, 0)
	for cursor.Next(ctx) {
		var row domain.Book
		err := cursor.Decode(&row)
		helper.PanicIfError(err)
		result = append(result, row)
	}
	return result
}

// Update implements BookRepository
func (repository *BookRepositoryImpl) Update(ctx context.Context, book domain.Book) domain.Book {
	var result domain.Book
	// Get Collection Names
	collection := repository.MongoClient.Database(os.Getenv("MONGO_INSTANCE")).Collection("book")
	// Update Object If Exist
	objId := book.Id
	filter := bson.M{"_id": objId}
	update := bson.M{"$set": bson.M{"title": book.Title, "author": book.Author, "year": book.Year}}
	_, err := collection.UpdateOne(ctx, filter, update)
	helper.PanicIfError(err)
	// Create New Object
	bookResult := collection.FindOne(ctx, bson.M{"_id": objId})
	// Return Value
	err = bookResult.Decode(&result)
	helper.PanicIfError(err)
	return result
}

// Delete implements BookRepository
func (repository *BookRepositoryImpl) Delete(ctx context.Context, bookId string) error {
	// Get Collection Names
	collection := repository.MongoClient.Database(os.Getenv("MONGO_INSTANCE")).Collection("book")
	// Delete Object If Exist
	objId, _ := primitive.ObjectIDFromHex(bookId)
	_, errDelete := collection.DeleteOne(ctx, bson.M{"_id": objId})
	return errDelete
}
