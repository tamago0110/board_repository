package handler

import (
	"golang_api/auth"
	"golang_api/controller"
	"golang_api/db"
	"golang_api/model"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
)

var (
	database         *gorm.DB
	testUser         string
	testPass         string
	testDbName       string
	testDbProtocol   string
	connectErr       error
	oldUser1         model.User
	jwtOldUser1      auth.JWT
	jwtOldUser1Err   error
	tokenSecret      string
	tokenIss         string
	oldUser2         model.User
	jwtOldUser2      auth.JWT
	jwtOldUser2Err   error
	oldUser3         model.User
	jwtOldUser3      auth.JWT
	jwtOldUser3Err   error
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
	tokenSecret = "secretKey"
	tokenIss = "tokenIss"

	if err := controller.CreateUuid(&oldUser1); err != nil {
		panic(err)
	}
	oldUser1.Email = "oldUser1@gmail.com"
	oldUser1.Password = "oldUser1"
	if err := controller.ValidateEmail(&oldUser1); err != nil {
		panic(err)
	}
	if err := controller.HashPassword(&oldUser1); err != nil {
		panic(err)
	}
	if err := controller.CreateUser(&oldUser1, database); err != nil {
		panic(err)
	}
	jwtOldUser1.Token, jwtOldUser1Err = auth.CreateToken(oldUser1.ID, tokenSecret, tokenIss)
	if jwtOldUser1Err != nil {
		panic(jwtOldUser1Err)
	}

	if err := controller.CreateUuid(&oldUser2); err != nil {
		panic(err)
	}
	oldUser2.Email = "oldUser2@gmail.com"
	oldUser2.Password = "oldUser2"
	if err := controller.ValidateEmail(&oldUser2); err != nil {
		panic(err)
	}
	if err := controller.HashPassword(&oldUser2); err != nil {
		panic(err)
	}
	if err := controller.CreateUser(&oldUser2, database); err != nil {
		panic(err)
	}
	jwtOldUser2.Token, jwtOldUser2Err = auth.CreateToken(oldUser2.ID, tokenSecret, tokenIss)
	if jwtOldUser2Err != nil {
		panic(jwtOldUser2Err)
	}

	if err := controller.CreateUuid(&oldUser3); err != nil {
		panic(err)
	}
	oldUser3.Email = "oldUser3@gmail.com"
	oldUser3.Password = "oldUser3"
	if err := controller.ValidateEmail(&oldUser3); err != nil {
		panic(err)
	}
	if err := controller.HashPassword(&oldUser3); err != nil {
		panic(err)
	}
	if err := controller.CreateUser(&oldUser3, database); err != nil {
		panic(err)
	}
	jwtOldUser3.Token, jwtOldUser3Err = auth.CreateToken(oldUser3.ID, tokenSecret, tokenIss)
	if jwtOldUser3Err != nil {
		panic(jwtOldUser3Err)
	}

	if err := controller.CreateLeadAsConsumer(&oldUser2, &user1ToUser2Lead, database); err != nil {
		panic(err)
	}
	user1ToUser2Lead.Producer = oldUser1.ID
	if err := controller.CreateLeadAsProducer(&oldUser1, &user1ToUser2Lead, database); err != nil {
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
