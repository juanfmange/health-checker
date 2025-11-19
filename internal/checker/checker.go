package checker

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/juanfmange/health-checker/internal/config"
)

type ServiceStatus struct {
	Name         string `json:"name"`
	URL          string `json:"url"`
	Status       string `json:"status"` // healthy or unhealthy
	ResponseTime int64  `json:"responseTime"`
	LastChecked  string `json:"lastChecked"`
	Message      string `json:"message,omitempty"`
}

func CheckServices(ctx context.Context, services []config.ServiceConfig) []ServiceStatus {
	ch := make(chan ServiceStatus)
	results := make([]ServiceStatus, len(services))

	for _, svc := range services {
		go func(svc config.ServiceConfig) {
			start := time.Now()

			client := http.Client{
				Timeout: time.Duration(svc.Timeout) * time.Second,
			}

			url := fmt.Sprintf("%s://%s%s", svc.Protocol, svc.Host, svc.Path)

			req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)

			resp, err := client.Do(req)
			responseTime := time.Since(start).Milliseconds()
			now := time.Now().UTC().Format(time.RFC3339)

			if err != nil {
				ch <- ServiceStatus{
					Name:         svc.Name,
					URL:          url,
					Status:       "unhealthy",
					ResponseTime: responseTime,
					LastChecked:  now,
					Message:      err.Error(),
				}
				return
			}
			defer resp.Body.Close()

			status := "healthy"
			msg := ""

			if resp.StatusCode >= 500 {
				status = "unhealthy"
				msg = fmt.Sprintf("Server error (%d)", resp.StatusCode)
			}

			ch <- ServiceStatus{
				Name:         svc.Name,
				URL:          url,
				Status:       status,
				ResponseTime: responseTime,
				LastChecked:  now,
				Message:      msg,
			}
		}(svc)
	}

	for i := range results {
		results[i] = <-ch
	}

	return results
}
