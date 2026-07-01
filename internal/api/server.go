package api

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/cyberangel1/atlas/internal/config"
	"github.com/cyberangel1/atlas/internal/health"
)

func collect(cfg *config.Config) []health.Status {

	var results []health.Status

	for _, s := range cfg.Services {
		results = append(results, health.CheckService(s))
	}

	return results
}

func Start(cfg *config.Config) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		t := template.Must(template.ParseFiles("web/templates/index.html"))

		t.Execute(w, collect(cfg))

	})

	http.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(collect(cfg))

	})

	http.ListenAndServe(":8080", nil)

}
