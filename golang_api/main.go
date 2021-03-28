package main

import (
	"github.com/joho/godotenv"
	"golang_api/db"
	"golang_api/server"
	"os"
)

func main() {
	if loadErr := godotenv.Load(".env"); loadErr != nil {
		panic(loadErr)
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbProtocol := os.Getenv("DB_PROTOCOL")
	tokenSecret := os.Getenv("TOKEN_SECRET")
	tokenIss := os.Getenv("TOKEN_ISS")
	imageDir := os.Getenv("IMAGE_DIR")
	imgDirPath := os.Getenv("IMGDIR_PATH")

	database, databaseErr := db.Init(dbUser, dbPass, dbName, dbProtocol)
	if databaseErr != nil {
		panic(databaseErr)
	}
	if serverErr := server.Init(database, tokenSecret, tokenIss, imageDir, imgDirPath); serverErr != nil {
		panic(serverErr)
	}
	if closeErr := db.Close(database); closeErr != nil {
		panic(closeErr)
	}
}
