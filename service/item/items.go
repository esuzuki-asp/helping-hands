package item

type Service struct {
	repo Repository
}

type Repository interface {
	GetItems() ([]Item, error)
}

type Item struct {
	ID              int64  `db:"id" json:"id"`
	Category        string `db:"category" json:"category"`
	Type            string `db:"type" json:"type"`
	Subtype         string `db:"subtype" json:"subtype"`
	CategoryFilters string `db:"tags" json:"category_filters"`
	AvailableStart  string `db:"available_start" json:"available_start"`
	AvailableEnd    string `db:"available_end" json:"available_end"`
}

type Location struct {
	City         string `db:"city" json:"city"`
	Country      string `db:"country" json:"country"`
	MeetingPoint string `db:"meeting_point" json:"meeting_point"`
}

type ItemWithLocation struct {
	Item
	Location
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) GetItems() {

}
