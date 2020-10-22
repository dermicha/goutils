package database_pg

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dBConn *gorm.DB
)

func CloseDatabase() {
	log.Debug("close database")
}

func GetDb() *gorm.DB {
	return dBConn
}

func InitDatabase(dbDSN string, conf *gorm.Config) {
	log.Println("testing postgres...")

	var gConf *gorm.Config
	if conf != nil {
		gConf = conf
	} else {
		gConf = &gorm.Config{
			DryRun:            false,
			PrepareStmt:       false,
			AllowGlobalUpdate: true,
			//Logger: logger.,
			Logger: logger.Default.LogMode(logger.Warn),
		}
	}

	if dbDSN == "" {
		dbDSN = "user=postgres dbname=gorm host=localhost port=5432 sslmode=disable TimeZone=Europe/Paris"
	}
	db, _ := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbDSN,
		PreferSimpleProtocol: true,
	}), gConf)

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
