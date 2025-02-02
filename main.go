package main

import (
	routes "butler_backend/routes"
	"log"
)

func main() {
	r := routes.Router()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}

	r.Run(":8080")
}
