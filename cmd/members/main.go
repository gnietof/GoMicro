package main

import (
	members "gomicro/internal/members"
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

	// locations, err := fetchLocations()
	controller := members.NewMembersController()

	http.HandleFunc("/members", controller.GetMembers)
	http.HandleFunc("/member/", controller.GetMemberById)

	log.Println("Locations service running on :8081!")
	http.ListenAndServe(":8081", nil)
}

// func fetchLocations() ([]model.Location, error) {
// 	resp, err := http.Get(os.Getenv(("LOCATIONS_SERVICE")))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var locations []model.Location
// 	err = json.NewDecoder(resp.Body).Decode(&locations)
// 	return locations, nil

// }
