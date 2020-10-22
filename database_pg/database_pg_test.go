package database_pg

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"testing"
)

type TestObject struct {
	gorm.Model
	UniqueName string `gorm:"column:unique_name" gorm:"uniqueIndex"`
}

func cleanUpDb() {
	db := GetDb()
	db.Exec("drop table if exists test_objects;")
}

func setup() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetLevel(log.InfoLevel)

	log.Info("test setup")

	gConf := &gorm.Config{
		DryRun:            false,
		PrepareStmt:       false,
		AllowGlobalUpdate: true,
		//Logger: logger.,
		Logger: logger.Default.LogMode(logger.Info),
	}

	InitDatabase("", gConf)
	cleanUpDb()
	MigrateDatabase(&TestObject{})
}

func tearDown() {
	log.Info("test teardown")
	CloseDatabase()
	cleanUpDb()
}

func TestSetup(t *testing.T) {
	setup()
	defer tearDown()
}

func TestCrud(t *testing.T) {
	setup()
	defer tearDown()

	tObj := TestObject{}
	tObj.UniqueName = "Name 1"
	resCreate := GetDb().Create(&tObj)
	if resCreate.RowsAffected != 1 {
		t.Fatalf("wrong number of rows effected: %d", resCreate.RowsAffected)
	}
	if tObj.ID < 1 {
		t.Fatalf("wrong object id: %d", tObj.ID)
	}
	utils.AssertEqual(resCreate.RowsAffected, 1)
	utils.AssertEqual(tObj.ID, 1)

	tObj2 := TestObject{}
	tObj2.UniqueName = "Name 2"
	resCreate2 := GetDb().Create(&tObj2)
	if resCreate2.RowsAffected != 1 {
		t.Fatalf("wrong number of rows effected: %d", resCreate2.RowsAffected)
	}
	if tObj2.ID < 1 {
		t.Fatalf("wrong object id: %d", tObj2.ID)
	}

	resDel := GetDb().Delete(tObj)
	if resDel.RowsAffected != 1 {
		t.Fatalf("wrong number of rows effected: %d", resDel.RowsAffected)
	}

	resDel2 := GetDb().Delete(tObj2)
	if resDel2.RowsAffected != 1 {
		t.Fatalf("wrong number of rows effected: %d", resDel2.RowsAffected)
	}
}

func TestUniqueIndex(t *testing.T) {
	setup()
	defer tearDown()

	tObj := TestObject{}
	tObj.UniqueName = "UniqueTest Name 1"
	resCreate1 := GetDb().Create(&tObj)
	resCreate2 := GetDb().Create(&tObj)

	if resCreate1.RowsAffected != 1 {
		t.Fatalf("wrong number of rows effected: %d", resCreate1.RowsAffected)
	}
	if tObj.ID < 1 {
		t.Fatalf("wrong object id: %d", tObj.ID)
	}
	if resCreate2.RowsAffected != 0 {
		t.Fatalf("wrong number of rows effected: %d", resCreate2.RowsAffected)
	}
	if resCreate2.Error.Error() != "ERROR: duplicate key value violates unique constraint \"test_objects_pkey\" (SQLSTATE 23505)" {
		t.Fatalf("unique index failed")
	}

	//utils.AssertEqual(t, resCreate1.RowsAffected, 1)
	//utils.AssertEqual(tObj.ID, 1)
	//utils.AssertEqual(resCreate2.RowsAffected, -1)
	//utils.AssertEqual(resCreate2.Error.Error(), "UNIQUE constraint failed: test_objects.id")

}
