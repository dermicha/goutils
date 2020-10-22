package database

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strings"
)

var (
	dBConn *gorm.DB
)

func CleanUpDb(dbName string) {
	if !strings.Contains(dbName, ":memory:") {
		e := os.Remove(getDbPath(dbName))
		if e != nil {
			log.Println(e)
		}
	}
}

func InitDatabase(dbName string, conf *gorm.Config) {
	log.Debug("initDatabase")
	var err error
	//logger := log.New()
	var gConf gorm.Config
	if conf != nil {
		gConf = *conf
	} else {
		gConf = gorm.Config{
			DryRun:            false,
			PrepareStmt:       false,
			AllowGlobalUpdate: true,
			//Logger: logger.,
			Logger: logger.Default.LogMode(logger.Warn),
		}
	}

	fullDbName := getDbPath(dbName)

	sqllDb := sqlite.Open(fullDbName)

	db, err := gorm.Open(sqllDb, &gConf)
	if err != nil {
		log.Panic(fmt.Sprintf("connection to database failed: %s", err))
	}
	log.Debug("Connection Opened to Database")

	//chmErr := os.Chmod(fullDbName, 0777)
	//if chmErr != nil {
	//	log.Panicf(chmErr.Error())
	//}

	dBConn = db
}

func MigrateDatabase(dbObject interface{}) {
	log.Debug("MigrateDatabase")
	err := dBConn.AutoMigrate(dbObject)
	if err != nil {
		log.Panic(fmt.Sprintf("database migration failed: %s", err))
	}
	log.Debug("database was migrated successfully")
}

func CloseDatabase() {
	log.Debug("close database")
}

func getDbPath(dbName string) string {
	if !strings.Contains(dbName, ":memory:") {
		path := fmt.Sprintf("%s.db", dbName)
		return path
	} else {
		return dbName
	}
}

func GetDb() *gorm.DB {
	return dBConn
}
