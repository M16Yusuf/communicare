package main

import (
	"context"
	"log"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/m16yusuf/communicare/internal/configs"
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
	db, err := configs.InitDB()
	if err != nil {
		log.Println("FAILED TO CONNECT DB")
		return
	}
	defer db.Close()
	// test ping database
	err = configs.PingDB(db)
	if err != nil {
		log.Println("PING TO DB FAILED", err.Error())
		return
	}
	log.Println("database connected")

	// inizialization redish for this project
	rdb := configs.InitRedis()
	cmd := rdb.Ping(context.Background())
	if cmd.Err() != nil {
		log.Println("failed ping on redis \nCause:", cmd.Err().Error())
		return
	}
	log.Println("Redis Connected")
	defer rdb.Close()

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
