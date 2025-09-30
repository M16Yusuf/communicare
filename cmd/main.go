package main

import (
	"log"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/m16yusuf/communicare/internal/routers"
)

// @title 											communicare
// @version 										1.0
// @description 								communicare your social app
// @host												127.0.0.1:3000
// @securityDefinitions.apikey 	JWTtoken
// @in header
// @name Authorization
func main() {
	// load env manual
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load ENV \nCause: ", err.Error())
	}

	// inisialsation database for this project

	// inizialization redish for this project

	// inizialization engine gin
	router := routers.InitRouter()
	//  run the engine gin
	// Run this project on 127.0.0.1:3000 or localhost:3000
	if runtime.GOOS == "windows" {
		router.Run("127.0.0.1:3000")
	} else {
		router.Run(":3000")
	}
}
