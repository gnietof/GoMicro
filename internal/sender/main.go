package main

import (
	"encoding/json"
	"fmt"
	"gomicro/internal/model"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	// controller := locations.NewLocationsController()
	// locations := controller.GetLocations()
	// print(len(locations))

	ticker := time.NewTicker(10 * time.Second)
	// done := make(chan bool)
	defer ticker.Stop()

	// for {
	// 	select {
	// 	case t := <-ticker.C:
	// 		fmt.Println("Sender at:", t)
	// 		// doTask()
	// 	case <-done:
	// 		return
	// 	}
	// }

	for t := range ticker.C {
		fmt.Println("Sender at:", t)
		// doTask()
	}
}

func doTask() {

	resp, err := http.Get("http://localhost:8081/hello")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func fetchLocations() ([]model.Location, error) {
	resp, err := http.Get(os.Getenv(("LOCATIONS_SERVICE")))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var locations []model.Location
	err = json.NewDecoder(resp.Body).Decode(&locations)
	return locations, nil

}
