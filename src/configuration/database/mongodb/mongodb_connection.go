package mongodb

import (
	"context"
	"os"

	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL           = "MONGODB_URL"
	MONGODB_USER_DATABASE = "MONGODB_USER_DATABASE"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodbURI := os.Getenv(MONGODB_URL)
	mongodbDatabase := os.Getenv(MONGODB_USER_DATABASE)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbURI))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Conexão com o banco de dados MongoDB realizada")

	return client.Database(mongodbDatabase), nil
}
