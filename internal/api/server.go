package api

import (
	"encoding/json"
	"net/http"

	"github.com/cyberangel1/atlas/internal/config"
	"github.com/cyberangel1/atlas/internal/health"
)

func Start(cfg *config.Config) {

	http.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {

		var results []health.Status

		for _, service := range cfg.Services {
			results = append(results, health.CheckService(service))
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	})

	http.ListenAndServe(":8080", nil)
}
