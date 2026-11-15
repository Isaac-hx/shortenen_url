package main

import (
	"context"
	"log"
	"url_shortener_api/internal/config"
	"url_shortener_api/internal/database"

	"github.com/joho/godotenv"
)

func main() {
	//Load .env file to parsing
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[ERROR] Load .env file :%v", err)
	}

	//Call construct object config
	cfg := config.NewConfig()
	//Calling context background
	ctx := context.Background()

	//Calling function database postgres to open connection
	// db, err := database.NewDbConn(ctx, cfg)
	// if err != nil {
	// 	log.Fatalf("%v", err)
	// }
	// defer db.Close()

	//calling a database supabase to open connection
	db, err := database.NewDbSupabase(ctx, cfg)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer db.Close(ctx)
}
