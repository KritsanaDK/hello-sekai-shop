package main

import (
	"context"
	"log"
	"os"

	"github.com/KaDingMeaw/hello-sekai-shop/config"
	"github.com/KaDingMeaw/hello-sekai-shop/pkg/database"
)

func main() {
	ctx := context.Background()

	// Initialize config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())

	// Database connection
	db := database.DbConn(ctx, &cfg)

	defer db.Disconnect(ctx)
	log.Println(db)

}
