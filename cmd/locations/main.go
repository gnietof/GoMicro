package main

import (
	"gomicro/internal/locations"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	log.Println("URL: " + os.Getenv("DB2_HOST"))

	controller := locations.NewLocationsController()

	http.HandleFunc("/locations", controller.GetLocations)
	http.HandleFunc("/location/", controller.GetLocationsById)

	log.Println("Locations service running on :8080!")
	http.ListenAndServe(":8080", nil)
}
