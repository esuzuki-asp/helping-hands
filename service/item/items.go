package item

import (
	"helping-hands/service/location"

	"github.com/sirupsen/logrus"
)

type Service struct {
	repo Repository
}

type Repository interface {
	GetItems(string, string, string, string) ([]Item, error)
	AddToCart(string, string) error
	CreateItem(ItemWithLocationIDAndContact) (int64, error)
}

type Item struct {
	ID              int64  `db:"id" json:"id"`
	Category        string `db:"category" json:"category"`
	Type            string `db:"_type" json:"type"`
	Subtype         string `db:"subtype" json:"subtype"`
	CategoryFilters string `db:"tags" json:"category_filters"`
	AvailableStart  string `db:"available_start" json:"available_start"`
	AvailableEnd    string `db:"available_end" json:"available_end"`
}

type ContactInfo struct {
	Phone string `db:"phone" json:"phone"`
	Email string `db:"email" json:"email"`
}

type ItemWithLocationIDAndContact struct {
	Item
	ContactInfo
	LocationID int64 `db:"location_id"`
	UserID     int64 `db:"user_id"`
}

type ItemWithLocation struct {
	Item
	location.Location
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) GetItems(req *getItemsRequest, res *getItemsResponse) error {
	items, err := s.repo.GetItems(req.Category, req.Type, req.Subtype, req.CategoryFilters)
	if err != nil {
		logrus.Error("GetItems: ", err)
		return err
	}
	res.Items = items
	return nil
}

func (s Service) AddToCart(req *addToCartRequest, res *addToCartResponse) error {
	err := s.repo.AddToCart(req.ItemID, req.UserID)
	if err != nil {
		logrus.Error("AddToCart: ", err)
		return err
	}
	return nil
}

func (s Service) CreateItem(req *createItemRequest, res *createItemResponse) error {
	item := ItemWithLocationIDAndContact{
		Item: Item{
			Category:        req.Category,
			Type:            req.Type,
			Subtype:         req.Subtype,
			CategoryFilters: req.CategoryFilters,
			AvailableStart:  req.AvailableStart,
			AvailableEnd:    req.AvailableEnd,
		},
		ContactInfo: ContactInfo{
			Phone: req.Phone,
			Email: req.Email,
		},
		LocationID: req.LocationID,
		UserID:     req.UserID,
	}
	logrus.Println(item)
	itemID, err := s.repo.CreateItem(item)
	if err != nil {
		logrus.Error("CreateItem: ", err)
		return err
	}
	res.ID = itemID
	return nil
}
