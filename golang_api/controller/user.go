package controller

import (
	"golang_api/model"
	"os"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

func CreateUuid(user *model.User) error {
	uuidObj, uuidErr := uuid.NewUUID()
	if uuidErr != nil {
		return uuidErr
	}
	user.ID = uuidObj.String()
	return nil
}

func ValidateEmail(user *model.User) error {
	if err := validator.New().Struct(user); err != nil {
		return err
	}
	return nil
}

func HashPassword(user *model.User) error {
	hash, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if hashErr != nil {
		return hashErr
	}
	user.Password = string(hash)
	return nil
}

func CreateUser(user *model.User, db *gorm.DB) error {
	if err := db.Create(user); err != nil {
		return err.Error
	}
	return nil
}

func SelectUserWhereEmail(user *model.User, db *gorm.DB) error {
	if err := db.Where("email = ?", user.Email).First(user); err != nil {
		return err.Error
	}
	return nil
}

func CheckPassword(user *model.User, rowPassword string) error {
	inputPassErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(rowPassword))
	if inputPassErr != nil {
		return inputPassErr
	}
	return nil
}

func SelectProfileWhereUserid(userID string, profile *model.Profile, db *gorm.DB) error {
	if err := db.Where("user_id = ?", userID).First(profile); err != nil {
		return err.Error
	}
	return nil
}

func SelectAllLeads(userID string, leads *[]model.Lead, db *gorm.DB) error {
	if err := db.Where("producer = ?", userID).Find(leads); err != nil {
		return err.Error
	}
	return nil
}

func SelectAllConsumerProfiles(consumers []string, profiles *[]model.Profile, db *gorm.DB) error {
	if err := db.Where("user_id in (?)", consumers).Find(profiles); err != nil {
		return err.Error
	}
	return nil
}

func SelectUserWhereId(userID string, user *model.User, db *gorm.DB) error {
	if err := db.Where("id = ?", userID).First(user); err != nil {
		return err.Error
	}
	return nil
}

func CreateProfile(user *model.User, profile *model.Profile, db *gorm.DB) error {
	profile.UserID = user.ID
	if err := db.Model(user).Association("Profile").Append(profile); err != nil {
		return err.Error
	}
	return nil
}

func CreateLeadAsConsumer(user *model.User, lead *model.Lead, db *gorm.DB) error {
	lead.Consumer = user.ID
	if err := db.Model(user).Association("Consumer").Append(lead); err != nil {
		return err.Error
	}
	return nil
}

func CreateLeadAsProducer(user *model.User, lead *model.Lead, db *gorm.DB) error {
	if err := db.Model(user).Association("Producer").Append(lead); err != nil {
		return err.Error
	}
	return nil
}

func RemoveImg(imgName string) error {
	if err := os.Remove(imgName); err != nil {
		return err
	}
	return nil
}

func UpdateProfile(oldProfile *model.Profile, newProfile *model.Profile, newName string, newImg string, db *gorm.DB) error {
	newProfile.ID = oldProfile.ID
	newProfile.UserID = oldProfile.UserID
	newProfile.Name = newName
	newProfile.Image = newImg
	if err := db.Model(oldProfile).Update(newProfile); err != nil {
		return err.Error
	}
	return nil
}
