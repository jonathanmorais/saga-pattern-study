package initializers

import (
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"log"
)

func ConnectToMongo() {
	mongoUrl := os.Getenv("DB_URL")
	ctx := context.TODO()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil { 
		log.Fatal("Failed to connect with Mongo", err)
	}
	log.Println("Connect succeful")
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()	
}