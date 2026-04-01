package locations

import (
	"gomicro/internal/model"
	"gomicro/internal/shared/db2"
)

type LocationsRepository struct {
	DB *db2.DB2Client
}

func NewLocationsRepository() *LocationsRepository {
	db, _ := db2.NewDB2Client()
	return &LocationsRepository{DB: db}
}

func (m *LocationsRepository) GetLocations() ([]model.Location, error) {

	rows, err := m.DB.DB.Query("SELECT WLC,CAMPUS_NAME,CAMPUS_ID,COUNTRY,CITY FROM ER_POPULATIONS.LOCS")
	if err != nil {
		print(err)
	}

	// var members []map[string]interface{}
	var locations []model.Location

	for rows.Next() {
		var location model.Location
		rows.Scan(&location.WLC, &location.CampusId, &location.CampusName, &location.Country, &location.City)

		locations = append(locations, location)

	}

	return locations, nil
}

func (m *LocationsRepository) GetLocationById(wlc string) (model.Location, error) {

	row := m.DB.DB.QueryRow("SELECT WLC,CAMPUSNAME,CAMPUSID,COUNTRY,CITY FROM ER_POPULATIONS.LOCS WHERE WLC=?", wlc)

	var location model.Location
	err := row.Scan(&location.CampusId, &location.CampusName, &location.Country, &location.City)
	if err != nil {
		return location, err
	}

	// result := map[string]interface{}{
	// 	"id":        id,
	// 	"firstname": firstname,
	// 	"lastname":  lastname,
	// 	"email":     email,
	// }

	// json.NewEncoder(w).Encode(result)
	return location, nil
}
