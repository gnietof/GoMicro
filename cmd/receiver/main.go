package main

import (
	"context"
	"encoding/json"
	"gomicro/internal/model"
	"gomicro/internal/shared/postgres"
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
			log.Fatal("Error serializing object.")
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

	conn, err := postgres.NewPostgresClient(ctx)

	if err != nil {
		log.Fatalf("Error connecting to Postgres: %s.", err)
		return err
	}
	_, err = conn.Conn.Exec(ctx, `INSERT INTO IBM.LOCATIONS 
		(WLC,CAMPUS_NAME,CAMPUS_ID,COUNTRY,CITY) 
		VALUES ($1,$2,$3,$4,$5)`,
		location.WLC, location.CampusName, location.CampusId,
		location.Country, location.City)

	conn.Conn.Close(ctx)

	if err != nil {
		log.Fatalf("Error inserting location: %s.", err)
		return err
	}

	return nil

}
