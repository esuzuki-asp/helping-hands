package item

import (
	"errors"
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

func (r repository) GetItems(category, _type, subtype, tags string) ([]Item, error) {
	var items []Item
	statement := `
		SELECT
			i.id,
			i.category,
			i._type,
			i.subtype,
			i.available_start::text,
			i.available_end::text,
			i.tags
		FROM
			item i
		WHERE
			i.category = $1
		AND
		    i._type = $2  
		AND
			i.is_available = true
	`
	err := r.db.Select(&items, statement, category, _type)
	return items, err
}

func (r repository) AddToCart(itemID, userID string) error {
	statement := `
		UPDATE TABLE item SET is_available = false WHERE id = $1
	`
	_, err := r.db.Exec(statement, itemID)
	if err != nil {
		return err
	}

	statement = `
		INSERT INTO user_order
		(user_id, item_id, status)
		VALUES
		($1, $2, 'cart')
	`
	_, err = r.db.Exec(statement, userID, itemID)
	return err
}

func (r repository) CreateItem(item ItemWithLocationIDAndContact) (int64, error) {
	var itemID int64
	statement := `
		INSERT INTO item
		(category, _type, subtype, tags, available_start, available_end, pickup_location, is_available)
		VALUES
		($1, $2, $3, $4, $5::date, $6::date, $7, true)
		RETURNING id
	`
	err := r.db.Get(&itemID, statement,
		item.Category,
		item.Type,
		item.Subtype,
		item.CategoryFilters,
		item.AvailableStart,
		item.AvailableEnd,
		item.LocationID,
	)
	if err != nil {
		return itemID, err
	}
	if itemID == 0 {
		return itemID, errors.New("itemID is zero")
	}

	statement = `
		INSERT INTO user_item
		(user_id, item_id, phone, email)
		VALUES
		($1, $2, $3, $4)
	`
	_, err = r.db.Exec(statement, item.UserID, itemID, item.Phone, item.Email)
	return itemID, err
}
