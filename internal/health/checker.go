package health

import (
	"net/http"
	"time"

	"github.com/cyberangel1/atlas/internal/config"
)

type Status struct {
	Name       string
	URL        string
	Type       string
	State      string
	StatusCode int
	Latency    time.Duration
	Error      string
	CheckedAt  time.Time
}

func CheckService(service config.Service) Status {
	start := time.Now()

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(service.URL)
	latency := time.Since(start)

	status := Status{
		Name:      service.Name,
		URL:       service.URL,
		Type:      service.Type,
		Latency:   latency,
		CheckedAt: time.Now(),
	}

	if err != nil {
		status.State = "DOWN"
		status.Error = err.Error()
		return status
	}

	defer resp.Body.Close()

	status.StatusCode = resp.StatusCode

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		status.State = "UP"
	} else {
		status.State = "DEGRADED"
	}

	return status
}
