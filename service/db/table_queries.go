package db

const (
	locationTableQuery = `
	CREATE TABLE location (
		id bigint primary key,
		country text not null,
		city text not null,
		meeting_point text not null
	);`
	userTableQuery = `
	CREATE TABLE user_ (
		id bigint primary key,
		first_name text not null,
		last_name text not null,
		username text not null,
		password text not null,
		location text not null,
		email text not null,
		preferred_pickup_location bigint references location(id) default null,
		preferred_dropoff_location bigint references location(id) default null
	);`
	itemTableQuery = `
	CREATE TABLE item (
		id bigint primary key,
		category text not null,
		type text not null,
		subtype text not null default '',
		is_available bool not null,
		available_start	date not null,
		available_end date not null,
		tags jsonb not null default '{}'::jsonb,
		image text default null,
		pickup_location bigint references location(id)
	);`
	userItemTableQuery = `
	CREATE TABLE user_item (
		user_id bigint references user_(id),
		item_id bigint references user_(id),
		phone text not null,
		email text not null
	);`
	userOrderTableQuery = `
	CREATE TABLE user_order (
		user_id bigint references user_(id),
		item_id bigint references user_(id),
		status text not null,
		pickup_date date,
		pickup_time time
	);`
)
