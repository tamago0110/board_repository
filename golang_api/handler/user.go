package handler

import (
	"fmt"
	"golang_api/auth"
	"golang_api/controller"
	"golang_api/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Signup(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		if err := controller.CreateUuid(&user); err != nil {
			c.String(http.StatusBadRequest, "Request failed"+err.Error())
			return
		}

		if err := controller.ValidateEmail(&user); err != nil {
			c.String(http.StatusBadRequest, "Request failed"+err.Error())
			return
		}

		if err := controller.HashPassword(&user); err != nil {
			c.String(http.StatusBadRequest, "Request failed"+err.Error())
			return
		}

		if err := controller.CreateUser(&user, db); err != nil {
			c.String(http.StatusBadRequest, "Request failed"+err.Error())
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func Login(db *gorm.DB, tokenSecret, tokenIss string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		var jwtToken auth.JWT

		if err := c.ShouldBindJSON(&user); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		rowPassword := user.Password

		if err := controller.SelectUserWhereEmail(&user, db); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		if err := controller.CheckPassword(&user, rowPassword); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		token, tokenErr := auth.CreateToken(user.ID, tokenSecret, tokenIss)
		if tokenErr != nil {
			c.String(http.StatusUnprocessableEntity, tokenErr.Error())
			return
		}

		jwtToken.Token = token
		c.JSON(http.StatusOK, jwtToken)
	}
}

func GetMyProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var profile model.Profile

		userID, ok := c.Get("userID")
		if !ok {
			c.String(http.StatusUnauthorized, "You are not authorized user.")
			return
		}
		strUserID := fmt.Sprintf("%v", userID)

		if err := controller.SelectProfileWhereUserid(strUserID, &profile, db); err != nil {
			c.String(http.StatusBadRequest, "Request failed"+err.Error())
			return
		}

		c.JSON(http.StatusOK, profile)
	}
}

func ListProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var profiles []model.Profile
		var leads []model.Lead
		var consumers []string

		userID, ok := c.Get("userID")
		if !ok {
			c.String(http.StatusUnauthorized, "You are not authorized user.")
			return
		}
		strUserID := fmt.Sprintf("%v", userID)

		if err := controller.SelectAllLeads(strUserID, &leads, db); err != nil {
			c.String(http.StatusBadRequest, "Request failed"+err.Error())
			return
		}

		for _, v := range leads {
			consumers = append(consumers, v.Consumer)
		}

		if err := controller.SelectAllConsumerProfiles(consumers, &profiles, db); err != nil {
			c.String(http.StatusBadRequest, "Request failed"+err.Error())
			return
		}

		c.JSON(http.StatusOK, profiles)
	}
}

func GetProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var profile model.Profile

		_, ok := c.Get("userID")
		if !ok {
			c.String(http.StatusUnauthorized, "You are not authorized user.")
			return
		}

		searchUserID := c.Param("uuid")

		if err := controller.SelectProfileWhereUserid(searchUserID, &profile, db); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		c.JSON(http.StatusOK, profile)
	}
}

func PostProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		var profile model.Profile

		if err := c.ShouldBindJSON(&profile); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		userID, ok := c.Get("userID")
		if !ok {
			c.String(http.StatusUnauthorized, "You are not authorized user.")
			return
		}
		strUserID := fmt.Sprintf("%v", userID)

		if err := controller.SelectUserWhereId(strUserID, &user, db); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		if err := controller.CreateProfile(&user, &profile, db); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		c.JSON(http.StatusOK, profile)
	}
}

func PutProfile(db *gorm.DB, imageDir, imgDirPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var oldProfile model.Profile

		userID, ok := c.Get("userID")
		if !ok {
			c.String(http.StatusUnauthorized, "You are not authorized user.")
			return
		}
		strUserID := fmt.Sprintf("%v", userID)

		if err := controller.SelectProfileWhereUserid(strUserID, &oldProfile, db); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		if oldProfile.UserID != userID {
			c.String(http.StatusUnauthorized, "Request failed: You can't edit this profile")
			return
		}

		newName := c.PostForm("name")
		var newImg string

		if file, err := c.FormFile("image"); file != nil && err == nil {
			if oldProfile.Image != "" {
				oldImage := strings.TrimPrefix(oldProfile.Image, imageDir)
				if removeErr := controller.RemoveImg(imgDirPath + oldImage); removeErr != nil {
					c.String(http.StatusBadRequest, "Request failed: "+removeErr.Error())
					return
				}
			}

			uniqueImgName := strUserID + "_" + file.Filename
			fileName := imgDirPath + uniqueImgName

			if upErr := c.SaveUploadedFile(file, fileName); upErr != nil {
				c.String(http.StatusBadRequest, "Request failed: "+err.Error())
				return
			}

			newImg = imageDir + uniqueImgName

		} else if file != nil && err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return

		} else {
			if oldProfile.Image != "" {
				newImg = oldProfile.Image
			} else {
				newImg = ""
			}
		}

		var newProfile model.Profile
		if err := controller.UpdateProfile(&oldProfile, &newProfile, newName, newImg, db); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		c.JSON(http.StatusOK, newProfile)
	}
}

func PostLead(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var consumerUser model.User
		var producerUser model.User
		var lead model.Lead

		if err := c.ShouldBindJSON(&lead); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		userID, ok := c.Get("userID")
		if !ok {
			c.String(http.StatusUnauthorized, "You are not authorized user.")
			return
		}
		strUserID := fmt.Sprintf("%v", userID)

		if err := controller.SelectUserWhereId(strUserID, &consumerUser, db); err != nil {
			c.String(http.StatusUnauthorized, "Request failed: You can't edit this profile")
			return
		}

		if err := controller.SelectUserWhereId(lead.Producer, &producerUser, db); err != nil {
			c.String(http.StatusUnauthorized, "Request failed: You can't edit this profile")
			return
		}

		if err := controller.CreateLeadAsConsumer(&consumerUser, &lead, db); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		if err := controller.CreateLeadAsProducer(&producerUser, &lead, db); err != nil {
			c.String(http.StatusBadRequest, "Request failed: "+err.Error())
			return
		}

		c.JSON(http.StatusOK, lead)
	}
}
