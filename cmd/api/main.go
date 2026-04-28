package main

import (
	"fmt"

	"github.com/Djuanzz/mutasiq-backend/internal/config"
	"github.com/Djuanzz/mutasiq-backend/internal/model"
	"github.com/Djuanzz/mutasiq-backend/internal/router"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println("===== Cashlens Backend =====")
	db := config.ConnectDatabase()
	defer config.CloseDatabase(db)

	if err := model.Migrate(db); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	r := router.SetupRouter(db)
	r.Run(":5000")
}
