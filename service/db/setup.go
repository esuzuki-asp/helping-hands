package db

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	fixtures "gopkg.in/testfixtures.v2"
)

var tableNames = []string{"location", "user_", "item", "user_item", "user_order"}

var DB *sqlx.DB

func init() {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load(".env")
		if err != nil {
			logrus.Panic(err)
		}
	}
	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))

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
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

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
