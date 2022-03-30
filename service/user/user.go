package user

import (
	"database/sql"
	"helping-hands/service/item"
	"helping-hands/service/location"
)

type Service struct {
	repo Repository
}

type Repository interface {
	GetCart(int64) ([]item.ItemWithLocation, error)
	GetOrders(int64) ([]OrderWithFullDetails, error)
	CreateUser(User) (int64, error)
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

type User struct {
	ID                       int64         `db:"id"`
	Username                 string        `db:"username" json:"username"`
	Password                 string        `db:"password" json:"password"`
	FirstName                string        `db:"first_name" json:"first_name"`
	LastName                 string        `db:"last_name" json:"last_name"`
	Location                 string        `db:"location" json:"location"`
	Email                    string        `db:"email" json:"email"`
	PreferredPickupLocation  sql.NullInt64 `db:"preferred_pickup_location"`
	PreferredDropoffLocation sql.NullInt64 `db:"preferred_dropoff_location"`
}

type OrderWithFullDetails struct {
	item.Item
	location.Location
	item.ContactInfo
	PickupDate string `db:"pickup_date" json:"pickup_date"`
	PickupTime string `db:"pickup_time" json:"pickup_time"`
	Status     string `db:"status" json:"status"`
}

func (s Service) GetCart(req *getCartRequest, res *getCartResponse) error {
	cart, err := s.repo.GetCart(req.UserID)
	if err != nil {
		return err
	}
	res.Items = cart
	return nil
}

func (s Service) GetOrders(req *getOrdersRequest, res *getOrdersResponse) error {
	orders, err := s.repo.GetOrders(req.UserID)
	if err != nil {
		return err
	}
	res.Items = orders
	return nil
}

func (s Service) CreateUser(req *createUserRequest, res *createUserResponse) error {
	user := User{
		Username:  req.Username,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Location:  req.Location,
		Email:     req.Email,
	}
	if req.PreferredDropoffLocation != 0 {
		user.PreferredDropoffLocation = sql.NullInt64{
			Int64: req.PreferredDropoffLocation,
			Valid: true,
		}
	}
	if req.PreferredPickupLocation != 0 {
		user.PreferredDropoffLocation = sql.NullInt64{
			Int64: req.PreferredPickupLocation,
			Valid: true,
		}
	}

	userID, err := s.repo.CreateUser(user)
	if err != nil {
		return err
	}
	res.UserID = userID
	return nil
}
