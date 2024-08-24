package database

import (
	"database/sql"
	"log"

	"github.com/Gustavopnhro/api-golang/configs"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {

	var err error
	config, _ := configs.LoadConfig()

	db, err = sql.Open("mysql", config.DBUser+":"+config.DBPassword+"@/"+config.DBName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
}

func GetDB() *sql.DB {
	return db
}
