package main

import (
	"fmt"
	"log"

	"github.com/cyberangel1/atlas/internal/config"
)

func main() {
	cfg, err := config.Load("configs/services.yaml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Atlas configuration loaded successfully")

	for _, service := range cfg.Services {
		fmt.Printf("- %s [%s] -> %s\n", service.Name, service.Type, service.URL)
	}
}
