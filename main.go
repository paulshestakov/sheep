package main

import (
	"main/app"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	databaseUrl := os.Getenv("DATABASE_URL")

	err := migrate(databaseUrl)
	if err != nil {
		panic(err)
	}

	app.App(databaseUrl, port)
}