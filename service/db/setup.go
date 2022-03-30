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
			//RemoveTables(db)
			log.Fatal("failed to setup db")
		}
	}()

	//addSequences(db)
	//useSequences(db)
	//createTables(db)
	//loadFixtures()
	//alterTable(db)

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

	logrus.Println("Try to skip db name check")
	fixtures.SkipDatabaseNameCheck(true)
	logrus.Println("Try to get fixtures")

	fixtures, err := fixtures.NewFolder(db, &fixtures.PostgreSQL{}, "./service/db/fixtures")
	if err != nil {
		logrus.Println("fail to retrieve fixtures")
		log.Panic(err)
	}

	logrus.Println("try to load fixtures")
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

func addSequences(db *sqlx.DB) {
	for _, tableName := range []string{"location", "user_", "item"} {
		_, err := db.Exec("CREATE SEQUENCE " + tableName + "_id_seq")
		if err != nil {
			log.Println("failed to add sequences")
			return
		}
	}

}

func useSequences(db *sqlx.DB) {
	for _, tableName := range []string{"location", "user_", "item"} {
		_, err := db.Exec("ALTER TABLE " + tableName + " ALTER COLUMN id SET DEFAULT nextval('" + tableName + "_id_seq')")
		if err != nil {
			log.Println("failed to add sequences to table " + tableName)
			return
		}
		_, err = db.Exec("ALTER SEQUENCE " + tableName + "_id_seq OWNED BY " + tableName + ".id")
		if err != nil {
			log.Println("failed to alter sequence for " + tableName)
			return
		}
	}

}

func alterTable(db *sqlx.DB) {
	_, err := db.Exec("ALTER TABLE item RENAME COLUMN type TO _type")
	if err != nil {
		log.Println("failed to adjust column name")
		return
	}
}
