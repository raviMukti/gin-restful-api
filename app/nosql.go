package app

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/raviMukti/gin-restful-api/helper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient() *mongo.Client {
	err := godotenv.Load()
	helper.PanicIfError(err)
	connectionString := "mongodb://" + os.Getenv("MONGO_USER") + ":" + os.Getenv("MONGO_PASSWORD") + "@" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT")
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	helper.PanicIfError(err)
	return client
}
