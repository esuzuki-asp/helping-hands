package user

import (
	"helping-hands/service/item"
)

type Service struct {
	repo Repository
}

type Repository interface {
	GetCart(int64) ([]item.ItemWithLocation, error)
	GetOrders(int64) ([]OrderWithFullDetails, error)
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

type OrderWithFullDetails struct {
	item.Item
	item.Location
	ContactInfo
	PickupDate string `db:"pickup_date" json:"pickup_date"`
	PickupTime string `db:"pickup_time" json:"pickup_time"`
	Status     string `db:"status" json:"status"`
}

type ContactInfo struct {
	Phone string `db:"phone" json:"phone"`
	Email string `db:"email" json:"email"`
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
