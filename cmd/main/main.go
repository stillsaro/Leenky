package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/stillsaro/Leenky/internal/database"
	"github.com/stillsaro/Leenky/internal/server"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loadig .env file: %v", err)
	}

	db, queries, err := database.ConnectToDB(os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := server.New(os.Getenv("LISTEN_ADDR"), db, queries, os.Getenv("JWT_KEY"))
	log.Fatal(server.Run())
}
