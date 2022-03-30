package location

import (
	"helping-hands/service/db"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository() Repository {
	return repository{
		db: db.DB,
	}
}

func (r repository) GetLocationByID(locationID int64) (Location, error) {
	var location Location
	statement := `
		SELECT
			l.id,
		    l.country,
		    l.city,
		    l.meeting_point
		FROM
			location l
		WHERE
			id = $1
	`
	err := r.db.Get(&location, statement, locationID)
	return location, err
}

func (r repository) GetLocationsByCountryAndCity(country, city string) ([]Location, error) {
	var locations []Location
	statement := `
		SELECT
			l.id,
		    l.country,
		    l.city,
		    l.meeting_point
		FROM
			location l
		WHERE
			l.country = $1
		AND
		    l.city = $2  	
	`
	err := r.db.Select(&locations, statement, country, city)
	return locations, err
}

func (r repository) CreateLocation(location Location) (int64, error) {
	var locationID int64
	statement := `
		INSERT INTO location 
		(city, country, meeting_point)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := r.db.Get(&locationID, statement,
		location.City,
		location.Country,
		location.MeetingPoint,
	)
	return locationID, err
}
