package main

import (
	"fmt"

	"github.com/Djuanzz/mutasiq-backend/internal/router"
)

type Transaction struct {
	Date   string  `json:"date"`
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
	Desc   string  `json:"desc"`
}

func main() {
	fmt.Println("===== Cashlens Backend =====")

	r := router.SetupRouter()
	r.Run(":5000")
}
