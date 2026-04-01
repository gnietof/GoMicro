package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gomicro/internal/model"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	ctx := context.Background()
	var wg sync.WaitGroup

	locs := make(chan string, 5)

	locations, err := fetchLocations()

	if err != nil {
		log.Fatal("Locations service did not provide any data!")
		return
	}

	for _, location := range locations {
		wg.Add(1)
		go sendLocation(ctx, location, &wg, locs)
	}

	go func() {
		wg.Wait()
		log.Printf("All threads completed (%d)", len(locations))
		close(locs)
	}()

	for loc := range locs {
		fmt.Printf("Processed: %s\n", loc)
	}

}

func sendLocation(ctx context.Context, location model.Location, wg *sync.WaitGroup, locs chan<- string) error {

	defer wg.Done()

	log.Printf("Sending location: %s", location.WLC)

	w := &kafka.Writer{
		Addr:     kafka.TCP(os.Getenv("KAFKA_BROKER")),
		Topic:    "gnftest",
		Balancer: &kafka.LeastBytes{},
	}

	payload, err := json.Marshal(location)
	if err != nil {
		log.Fatal("Error serializing object.")
	}
	err = w.WriteMessages(ctx,
		kafka.Message{
			// Key: []byte(location.WLC)
			Value: payload,
		},
	)
	if err != nil {
		log.Fatalf("Failed to write message: %s.", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("Failed to close writer.")
	}

	log.Printf("Thread for location %s completed", location.WLC)
	locs <- location.WLC
	return nil
}

func fetchLocations() ([]model.Location, error) {
	url := os.Getenv("LOCATIONS_SERVICE")
	if url == "" {
		return nil, fmt.Errorf("Locations service not properly configured.")
	}
	fmt.Printf("%s\n", url)

	// resp, err := http.Get(url + "/locations")
	// if err != nil {
	// 	return nil, err
	// }
	// defer resp.Body.Close()

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url + "/locations")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	var locations []model.Location
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		return nil, err
	}

	return locations, nil

}
