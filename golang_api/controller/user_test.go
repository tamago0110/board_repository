package controller

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"golang_api/model"
	"os"
	"testing"
)

var (
	newUser           model.User
	newUserPro        model.Profile
	newUserToOldUser1 model.Lead
	leads             []model.Lead
	consumers         []string
	consumerPros      []model.Profile
)

func TestCreateUuid(t *testing.T) {
	err := CreateUuid(&newUser)
	assert.NoError(t, err)
	assert.NotEmpty(t, newUser.ID)
}

func TestValidateEmail(t *testing.T) {
	newUser.Email = "test"
	noNilErr := ValidateEmail(&newUser)
	assert.Error(t, noNilErr)
	newUser.Email = "test@gmail.com"
	nilErr := ValidateEmail(&newUser)
	assert.NoError(t, nilErr)
}

func TestHashPassword(t *testing.T) {
	newUser.Password = "password"
	err := HashPassword(&newUser)
	assert.NoError(t, err)
	assert.NotEmpty(t, newUser.Password)
	var testStr string
	assert.IsType(t, newUser.Password, testStr)
	comparePassErr := bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte("password"))
	assert.NoError(t, comparePassErr)
}

func TestCheckPassword(t *testing.T) {
	nilErr := CheckPassword(&newUser, "password")
	assert.NoError(t, nilErr)
	noNilErr := CheckPassword(&newUser, "failPass")
	assert.Error(t, noNilErr)
}

func TestCreateUser(t *testing.T) {
	err := CreateUser(&newUser, database)
	assert.NoError(t, err)
}

func TestSelectUserWhereEmail(t *testing.T) {
	checkID := newUser.ID
	checkEmail := newUser.Email
	checkPass := newUser.Password
	getErr := SelectUserWhereEmail(&newUser, database)
	assert.NoError(t, getErr)
	assert.Equal(t, checkID, newUser.ID)
	assert.Equal(t, checkEmail, newUser.Email)
	assert.Equal(t, checkPass, newUser.Password)
}

func TestSelectUserWhereId(t *testing.T) {
	checkID := newUser.ID
	checkEmail := newUser.Email
	checkPass := newUser.Password
	getErr := SelectUserWhereId(newUser.ID, &newUser, database)
	assert.NoError(t, getErr)
	assert.Equal(t, checkID, newUser.ID)
	assert.Equal(t, checkEmail, newUser.Email)
	assert.Equal(t, checkPass, newUser.Password)
}

func TestCreateProfile(t *testing.T) {
	newUserPro.Name = "testPro"
	recordErr := CreateProfile(&newUser, &newUserPro, database)
	assert.NoError(t, recordErr)
}

func TestSelectProfileWhereUserid(t *testing.T) {
	getErr := SelectProfileWhereUserid(newUser.ID, &newUserPro, database)
	assert.NoError(t, getErr)
	assert.Equal(t, "testPro", newUserPro.Name)
	assert.Equal(t, newUser.ID, newUserPro.UserID)
}

func TestCreateLeadAsConsumer(t *testing.T) {
	newUserToOldUser1.Producer = oldUser1.ID
	recordErr := CreateLeadAsConsumer(&newUser, &newUserToOldUser1, database)
	assert.NoError(t, recordErr)
}

func TestCreateLeadAsProducer(t *testing.T) {
	recordErr := CreateLeadAsProducer(&oldUser1, &newUserToOldUser1, database)
	assert.NoError(t, recordErr)
}

func TestSelectAllLeads(t *testing.T) {
	getErr := SelectAllLeads(oldUser1.ID, &leads, database)
	assert.NoError(t, getErr)
	assert.Equal(t, 2, len(leads))
	assert.Equal(t, oldUser1.ID, leads[0].Producer)
	assert.Equal(t, oldUser1.ID, leads[1].Producer)
}

func TestSelectAllConsumerProfiles(t *testing.T) {
	consumers = append(consumers, newUser.ID)
	getErr := SelectAllConsumerProfiles(consumers, &consumerPros, database)
	assert.NoError(t, getErr)
	assert.Equal(t, newUser.ID, consumerPros[0].UserID)
	assert.Equal(t, newUserPro.Name, consumerPros[0].Name)
}

func TestUpdateProfile(t *testing.T) {
	newName := "newName"
	img, imgUpErr := os.Create("../image/test.png")
	if imgUpErr != nil {
		panic(imgUpErr)
	}
	newImg := img.Name()

	var newUserNewPro model.Profile
	updateErr := UpdateProfile(&newUserPro, &newUserNewPro, newName, newImg, database)
	assert.NoError(t, updateErr)
	assert.Equal(t, newUserPro.ID, newUserNewPro.ID)
	assert.Equal(t, newUser.ID, newUserNewPro.UserID)
	assert.Equal(t, newName, newUserNewPro.Name)
	assert.Equal(t, newImg, newUserNewPro.Image)
}

func TestRemoveImg(t *testing.T) {
	removeErr := RemoveImg("../image/test.png")
	assert.NoError(t, removeErr)
}
