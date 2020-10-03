package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	dBConn *gorm.DB
)

func CleanUpDb(dbName string) {
	path := fmt.Sprintf("./%s.db", dbName)
	e := os.Remove(path)
	if e != nil {
		log.Println(e)
	}
}

func InitDatabase(dbName string) {
	log.Println("initDatabase")
	var err error

	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", dbName)), &gorm.Config{DryRun: false})
	if err != nil {
		log.Panic(fmt.Sprintf("connection to database failed: %s", err))
	}
	log.Println("Connection Opened to Database")

	dBConn = db
}

func MigrateDatabase(dbObject interface{}) {
	log.Println("MigrateDatabase")
	err := dBConn.AutoMigrate(dbObject)
	if err != nil {
		log.Panic(fmt.Sprintf("database migration failed: %s", err))
	}
	log.Println("database was migrated successfully")
}

func CloseDatabase() {
	log.Println("close database")
}

func GetDb() *gorm.DB {
	return dBConn
}
