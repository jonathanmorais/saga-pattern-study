package models

import (
	"log"
	"context"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"


)

type FlightPost struct {
	Place        string  `json:"place"`
}


func PersistRequest(){
	
	dbm := os.Getenv("DB_URL")
	ctx := context.TODO()
	opt := options.Client().ApplyURI(dbm)
	client, err := mongo.NewClient(opt)
	if err != nil {
		panic(err)
	}

	autoresponderDB     := client.Database("sagadatabase")
	flightCollection := autoresponderDB.Collection("flightdatabase")


	if err != nil {
		panic(err)
	}
	var flightPost FlightPost

	flight := FlightPost{
		Place: flightPost.Place,
	}

	insertResult, err := flightCollection.InsertMany(ctx, flight)
	if err != nil {
	panic(err)
	}

	log.Println(insertResult)

}
