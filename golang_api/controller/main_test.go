package controller

import (
	"github.com/jinzhu/gorm"
	"golang_api/db"
	"golang_api/model"
	"os"
	"testing"
)

var (
	database         *gorm.DB
	testUser         string
	testPass         string
	testDbName       string
	testDbProtocol   string
	connectErr       error
	oldUser1         model.User
	oldUser2         model.User
	oldUser3         model.User
	user1ToUser2Lead model.Lead
	board1           model.Board
	board2           model.Board
)

func initDb() error {
	testUser = "test_user"
	testPass = "test"
	testDbName = "test_db"
	testDbProtocol = "tcp(mysql_container:3306)"
	database, connectErr = db.Init(testUser, testPass, testDbName, testDbProtocol)
	if connectErr != nil {
		return connectErr
	}
	return nil
}

func TestMain(m *testing.M) {
	if err := initDb(); err != nil {
		panic(err)
	}

	if err := CreateUuid(&oldUser1); err != nil {
		panic(err)
	}
	oldUser1.Email = "oldUser1@gmail.com"
	oldUser1.Password = "oldUser1"
	if err := ValidateEmail(&oldUser1); err != nil {
		panic(err)
	}
	if err := HashPassword(&oldUser1); err != nil {
		panic(err)
	}
	if err := CreateUser(&oldUser1, database); err != nil {
		panic(err)
	}

	if err := CreateUuid(&oldUser2); err != nil {
		panic(err)
	}
	oldUser2.Email = "oldUser2@gmail.com"
	oldUser2.Password = "oldUser2"
	if err := ValidateEmail(&oldUser2); err != nil {
		panic(err)
	}
	if err := HashPassword(&oldUser2); err != nil {
		panic(err)
	}
	if err := CreateUser(&oldUser2, database); err != nil {
		panic(err)
	}

	if err := CreateUuid(&oldUser3); err != nil {
		panic(err)
	}
	oldUser3.Email = "oldUser3@gmail.com"
	oldUser3.Password = "oldUser3"
	if err := ValidateEmail(&oldUser3); err != nil {
		panic(err)
	}
	if err := HashPassword(&oldUser3); err != nil {
		panic(err)
	}
	if err := CreateUser(&oldUser3, database); err != nil {
		panic(err)
	}

	if err := CreateLeadAsConsumer(&oldUser2, &user1ToUser2Lead, database); err != nil {
		panic(err)
	}
	user1ToUser2Lead.Producer = oldUser1.ID
	if err := CreateLeadAsProducer(&oldUser1, &user1ToUser2Lead, database); err != nil {
		panic(err)
	}

	code := m.Run()

	if err := database.Delete(&oldUser1).Error; err != nil {
		panic(err)
	}
	if err := database.Delete(&oldUser2).Error; err != nil {
		panic(err)
	}
	if err := database.Delete(&oldUser3).Error; err != nil {
		panic(err)
	}
	if err := database.Delete(&newUser).Error; err != nil {
		panic(err)
	}
	if err := database.Delete(&newUserPro).Error; err != nil {
		panic(err)
	}
	if err := database.Delete(&newUserToOldUser1).Error; err != nil {
		panic(err)
	}
	if err := database.Delete(&display).Error; err != nil {
		panic(err)
	}
	if err := database.Delete(&user1ToUser2Lead).Error; err != nil {
		panic(err)
	}

	if err := db.Close(database); err != nil {
		panic(err)
	}

	os.Exit(code)
}
