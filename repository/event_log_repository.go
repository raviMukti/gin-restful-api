package repository

import (
	"context"
	"os"

	"github.com/raviMukti/gin-restful-api/model/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventLogRepository interface {
	Save(ctx context.Context, eventLog domain.EventLog) error
}

type EventLogRepositoryImpl struct {
	MongoClient *mongo.Client
}

func NewEventLogRepository(mongoClient *mongo.Client) EventLogRepository {
	return &EventLogRepositoryImpl{
		MongoClient: mongoClient,
	}
}

// Save implements EventLogRepository
func (repository *EventLogRepositoryImpl) Save(ctx context.Context, eventLog domain.EventLog) error {
	// Get Collection Names
	collection := repository.MongoClient.Database(os.Getenv("MONGO_INSTANCE")).Collection("event_log")
	// Get Inserted Object
	_, errInsert := collection.InsertOne(ctx, eventLog)
	return errInsert
}
