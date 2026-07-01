package main

import (
	"fmt"
	"log"

	"github.com/cyberangel1/atlas/internal/config"
	"github.com/cyberangel1/atlas/internal/health"
)

func main() {
	cfg, err := config.Load("configs/services.yaml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Atlas Health Engine")
	fmt.Println("-------------------")

	for _, service := range cfg.Services {
		status := health.CheckService(service)

		if status.State == "UP" {
			fmt.Printf("✓ %s [%s] %s | code=%d | latency=%s\n", status.Name, status.Type, status.State, status.StatusCode, status.Latency)
		} else {
			fmt.Printf("✗ %s [%s] %s | code=%d | latency=%s | error=%s\n", status.Name, status.Type, status.State, status.StatusCode, status.Latency, status.Error)
		}
	}
}
