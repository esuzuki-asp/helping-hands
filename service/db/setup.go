package db

import (
	"database/sql"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	fixtures "gopkg.in/testfixtures.v2"
)

var tableNames = []string{"location", "user_", "item", "user_item", "user_order"}

var DB *sqlx.DB

func init() {
	conn := "user=postgres password=postgres host=127.0.0.1 port=7777 sslmode=disable"
	db, err := sqlx.Open("postgres", conn)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := recover(); err != nil {
			RemoveTables(db)
			log.Fatal("failed to setup db")
		}
	}()

	createTables(db)
	loadFixtures()

	DB = db
}

func createTables(db *sqlx.DB) {
	for _, query := range []string{
		locationTableQuery, userTableQuery, itemTableQuery, userItemTableQuery, userOrderTableQuery,
	} {
		_, err := db.Exec(query)
		if err != nil {
			log.Panic(err)
		}
	}
	logrus.Info("Create tables")
}

func loadFixtures() {
	conn := "user=postgres password=postgres host=127.0.0.1 port=7777 sslmode=disable"
	db, err := sql.Open("postgres", conn)

	fixtures.SkipDatabaseNameCheck(true)

	fixtures, err := fixtures.NewFolder(db, &fixtures.PostgreSQL{}, "./service/db/fixtures")
	if err != nil {
		log.Panic(err)
	}

	if err = fixtures.Load(); err != nil {
		log.Panic(err)
	}
	logrus.Info("Load fixtures")
}

func RemoveTables(db *sqlx.DB) {
	tableString := strings.Join(tableNames, ", ")
	_, err := db.Exec("DROP TABLE " + tableString)
	if err != nil {
		log.Println("failed to drop tables")
		return
	}
	logrus.Info("Delete tables")

}
