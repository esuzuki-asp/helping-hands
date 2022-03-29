package item

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

func (r repository) GetItems() ([]Item, error) {
	var items []Item
	//	statement := ``
	//	_, err := db.execute(statement, agrs ...)

	return items, nil

}
