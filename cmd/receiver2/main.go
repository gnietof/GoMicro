package main

import (
	"context"
	"encoding/json"
	"gomicro/internal/model"
	"gomicro/internal/shared/mongo"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	ctx := context.Background()

	receiveLocations(ctx)

}

func receiveLocations(ctx context.Context) error {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{os.Getenv("KAFKA_BROKER")},
		GroupID:  "gnf-group",
		Topic:    "gnftest",
		MaxBytes: 10e3,
		// StartOffset: kafka.FirstOffset,
	})
	r.SetOffset(0)

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			log.Fatal("Error serializing object.", err)
		}
		log.Printf("Message at %d: %s", m.Offset, string(m.Value))

		var location model.Location
		if err := json.Unmarshal(m.Value, &location); err != nil {
			log.Fatal("Error unmarshaling object:", err)
		}

		go addLocation(ctx, location)
	}

	// if err := r.Close(); err != nil {
	// 	log.Fatal("Failed to close writer.")
	// }

	// return nil
}

func addLocation(ctx context.Context, location model.Location) error {

	client, err := mongo.NewMongoClient(ctx)

	if err != nil {
		log.Fatalf("Error connecting to Mongo: %s.", err)
		return err
	}
	defer client.Client.Disconnect(ctx)

	db := client.Client.Database("ibm")
	log.Printf("Connected to DB: %s\n", db.Name())
	collection := db.Collection("locations")

	result, err := collection.InsertOne(ctx, location)
	if err != nil {
		log.Fatalf("Error inserting location: %s.", err)
		return err
	}
	log.Printf("Inserted ID: %s\n", result.InsertedID)

	return nil

}
