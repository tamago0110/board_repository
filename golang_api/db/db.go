package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang_api/model"
)

func Init(dbUser, dbPass, dbName, dbProtocol string) (*gorm.DB, error) {
	DBMS := "mysql"
	CONNECT := dbUser + ":" + dbPass + "@" + dbProtocol + "/" + dbName + "?parseTime=true"

	db, connectErr := gorm.Open(DBMS, CONNECT)
	if connectErr != nil {
		return nil, connectErr
	}

	if err := AutoMigrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func Close(db *gorm.DB) error {
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.User{}).Error; err != nil {
		return err
	}

	if err := db.AutoMigrate(&model.Profile{}).Error; err != nil {
		return err
	}

	if err := db.AutoMigrate(&model.Lead{}).Error; err != nil {
		return err
	}

	if err := db.AutoMigrate(&model.Board{}).Error; err != nil {
		return err
	}

	if err := db.AutoMigrate(&model.Display{}).Error; err != nil {
		return err
	}

	return nil
}
