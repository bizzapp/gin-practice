package main

import (
	"myapp/internal/config"
	"myapp/internal/router"
)

func main() {
	// Connect to PostgreSQL

	r := router.New()
	config.ConnectDatabase()
	r.Run(":8080")
}
