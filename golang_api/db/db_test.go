package db

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	db             *gorm.DB
	testUser       string
	testPass       string
	testDbName     string
	testDbProtocol string
	err            error
)

func TestInit(t *testing.T) {
	testUser = "test_user"
	testPass = "test"
	testDbName = "test_db"
	testDbProtocol = "tcp(mysql_container:3306)"
	db, err = Init(testUser, testPass, testDbName, testDbProtocol)

	var checkDatabase *gorm.DB
	assert.NoError(t, err)
	assert.NotEmpty(t, db)
	assert.IsType(t, db, checkDatabase)
}

func TestAutoMigrate(t *testing.T) {
	err = AutoMigrate(db)
	assert.NoError(t, err)
}

func TestClose(t *testing.T) {
	err = Close(db)
	assert.NoError(t, err)
}
