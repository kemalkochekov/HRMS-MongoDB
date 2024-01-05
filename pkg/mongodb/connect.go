package mongodb

import (
	"Human_Resources_Managament_System/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBManager struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongoDBManager(ctx context.Context, connectionString, dbName string, cfg *config.Config) (*MongoDBManager, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	clientOptions.SetAuth(options.Credential{
		Username: cfg.MongoDb.User,
		Password: cfg.MongoDb.Password,
	})
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	database := client.Database(dbName)

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &MongoDBManager{
		Client:   client,
		Database: database,
	}, nil
}

// CloseConnection closes the MongoDB connection
func (m *MongoDBManager) CloseConnection() error {
	return m.Client.Disconnect(context.Background())
}
