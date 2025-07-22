package main

import (
	"myapp/internal/router"
)

func main() {
	r := router.New()
	r.Run(":8080")
}
