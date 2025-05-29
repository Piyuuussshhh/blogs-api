package main

import (
	"blog-api/api"
	"blog-api/db"
	"context"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = db.Init(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Disconnect()

	log.Fatal(api.Serve())
}
