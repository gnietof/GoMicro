package mongo

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Client *mongo.Client
}

func NewMongoClient(ctx context.Context) (*MongoClient, error) {

	connStr := buildConnString()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	if err != nil {
		return nil, err
	}
	// defer conn.Close(ctx)

	return &MongoClient{Client: client}, nil
}

func buildConnString() string {
	host := os.Getenv("MONGO_HOST")
	user := os.Getenv("MONGO_USER")
	pwd := os.Getenv("MONGO_PWD")

	connStr := fmt.Sprintf("mongodb://%s:%s@%s:27017",
		user, pwd, host)

	return connStr
}
