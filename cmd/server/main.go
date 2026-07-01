package main

import (
	"fmt"
	"log"

	"github.com/cyberangel1/atlas/internal/api"
	"github.com/cyberangel1/atlas/internal/config"
)

func main() {

	cfg, err := config.Load("configs/services.yaml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("===================================")
	fmt.Println(" Atlas Infrastructure Control Plane")
	fmt.Println("===================================")

	fmt.Println("API running at http://localhost:8080/services")

	api.Start(cfg)
}
