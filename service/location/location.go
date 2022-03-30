package location

import "github.com/sirupsen/logrus"

type Service struct {
	repo Repository
}

type Repository interface {
	GetLocationByID(int64) (Location, error)
	GetLocationsByCountryAndCity(string, string) ([]Location, error)
	CreateLocation(Location) (int64, error)
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

type Location struct {
	ID           int64  `db:"id" json:"id"`
	City         string `db:"city" json:"city"`
	Country      string `db:"country" json:"country"`
	MeetingPoint string `db:"meeting_point" json:"meeting_point"`
}

func (s Service) GetLocation(req *getLocationRequest, res *getLocationResponse) error {
	location, err := s.repo.GetLocationByID(req.ID)
	if err != nil {
		logrus.Error("GetLocation: ", err)
		return err
	}
	res.ID = location.ID
	res.City = location.City
	res.Country = location.Country
	res.MeetingPoint = location.MeetingPoint
	return nil
}

func (s Service) GetLocations(req *getLocationsRequest, res *getLocationsResponse) error {
	locations, err := s.repo.GetLocationsByCountryAndCity(req.Country, req.City)
	if err != nil {
		logrus.Error("GetLocations: ", err)
		return err
	}
	res.Locations = locations
	return nil
}

func (s Service) CreateLocation(req *createLocationRequest, res *createLocationResponse) error {
	location := Location{
		City:         req.City,
		Country:      req.Country,
		MeetingPoint: req.MeetingPoint,
	}

	locationID, err := s.repo.CreateLocation(location)
	if err != nil {
		logrus.Error("CreateLocation: ", err)
		return err
	}
	res.LocationID = locationID
	return nil
}
