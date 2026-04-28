package main

import (
	"fmt"

	"github.com/Djuanzz/mutasiq-backend/internal/config"
	"github.com/Djuanzz/mutasiq-backend/internal/router"
)

func main() {
	fmt.Println("===== Cashlens Backend =====")
	db := config.ConnectDatabase()

	defer config.CloseDatabase(db)

	r := router.SetupRouter()
	r.Run(":5000")
}
