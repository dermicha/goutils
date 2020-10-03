package database

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"log"
	"testing"
)

type TestObject struct {
	gorm.Model
	UniqueName string `gorm:"column:unique_name" gorm:"uniqueIndex"`
}

var (
	dbName     = "test_db"
	testDbName = fmt.Sprintf("%s_test", dbName)
)

func setup() {
	log.Println("test setup")
	InitDatabase(testDbName)
	MigrateDatabase(&TestObject{})
	CleanUpDb(testDbName)
}

func tearDown() {
	log.Println("test teardown")
	CloseDatabase()
	//CleanUpDb(testDbName)
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
	utils.AssertEqual(resCreate.RowsAffected, 1)
	utils.AssertEqual(tObj.ID, 1)

	tObj2 := TestObject{}
	tObj2.UniqueName = "Name 2"
	resCreate2 := GetDb().Create(&tObj2)
	utils.AssertEqual(resCreate2.RowsAffected, 0)
	utils.AssertEqual(tObj2.ID, 0)

	resDel := GetDb().Delete(tObj)
	utils.AssertEqual(resDel.RowsAffected, 1)
}
