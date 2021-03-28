package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang_api/auth"
	"golang_api/handler"
)

func Init(db *gorm.DB, tokenSecret, tokenIss, imageDir, imgDirPath string) error {
	r := Router(db, tokenSecret, tokenIss, imageDir, imgDirPath)
	if err := r.Run(":8080"); err != nil {
		return err
	}
	return nil
}

func Router(db *gorm.DB, tokenSecret, tokenIss, imageDir, imgDirPath string) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	r.POST("/signup", handler.Signup(db))
	r.POST("/login", handler.Login(db, tokenSecret, tokenIss))
	user := r.Group("/user").Use(auth.Authz(tokenSecret, tokenIss))
	{
		user.GET("/me", handler.GetMyProfile(db))
		user.PUT("/me", handler.PutProfile(db, imageDir, imgDirPath))
		user.GET("/profile", handler.ListProfile(db))
		user.GET("/profile/:uuid", handler.GetProfile(db))
		user.POST("/profile", handler.PostProfile(db))
		user.POST("/lead", handler.PostLead(db))
	}
	board := r.Group("/board").Use(auth.Authz(tokenSecret, tokenIss))
	{
		board.POST("/reject", handler.RejectBoard(db))
		board.GET("/mine", handler.GetMyBoard(db))
		board.GET("/item", handler.ListBoard(db))
		board.POST("/item", handler.PostBoard(db))
		board.PUT("/item/:id", handler.PutBoard(db))
		board.DELETE("/item/:id", handler.DeleteBoard(db))
	}
	r.Static("/image", "./image")

	return r
}
