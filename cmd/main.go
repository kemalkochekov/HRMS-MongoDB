package main

import (
	"Human_Resources_Managament_System/config"
	"Human_Resources_Managament_System/internal/server"
	"Human_Resources_Managament_System/pkg/mongodb"
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("LoadConfig: %w", err)
	}

	mongoURI := fmt.Sprintf("mongodb://%s:%d", cfg.MongoDb.Host, cfg.MongoDb.Port)

	dbManager, err := mongodb.NewMongoDBManager(ctx, mongoURI, cfg.MongoDb.DBName, cfg)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB:", err)
	}

	defer dbManager.CloseConnection()

	app := server.NewServer(cfg, *dbManager)

	if err = app.Run(); err != nil {
		log.Printf("Cannot start server: %w", err)
	}
}
