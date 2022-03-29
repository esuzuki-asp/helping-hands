package user

import (
	"helping-hands/service/db"
	"helping-hands/service/item"

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

func (r repository) GetCart(userID int64) ([]item.ItemWithLocation, error) {
	var items []item.ItemWithLocation
	statement := `
		SELECT
			i.category,
			i.type,
			i.subtype,
			i.available_start::text,
			i.available_end::text,
			l.city,
			l.country,
			l.meeting_point
		FROM
			user_order uo
		JOIN
			item i ON uo.item_id = i.id
		JOIN
			location l on i.pickup_location = l.id
		WHERE
			user_id = $1
			AND status = 'cart'
	`
	err := r.db.Select(&items, statement, userID)
	return items, err
}

func (r repository) GetOrders(userID int64) ([]OrderWithFullDetails, error) {
	var orders []OrderWithFullDetails
	statement := `
		SELECT
			uo.status,
			uo.pickup_date::text,
			uo.pickup_time::text,
			i.id,
			i.category,
			i.type,
			i.subtype,
			i.tags,
			i.available_start,
			i.available_end,
			l.city,
			l.country,
			l.meeting_point,
			ui.phone,
			ui.email
	FROM
		user_order uo
	JOIN
		item i on uo.item_id = i.id
	JOIN
		user_item ui on i.id = ui.item_id
	JOIN
		location l on i.pickup_location = l.id
	WHERE 
		uo.user_id = $1
	AND
		(status = 'pending' OR status = 'complete')
	`
	err := r.db.Select(&orders, statement, userID)
	return orders, err
}
