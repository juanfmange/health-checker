package checker

import (
	"context"
	"net/http"
	"time"
)

type ServiceStatus struct {
	URL        string `json:"url"`
	Alive      bool   `json:"alive"`
	StatusCode int    `json:"status_code,omitempty"`
	Error      string `json:"error,omitempty"`
}

func CheckServices(ctx context.Context, urls []string, timeoutSeconds int) []ServiceStatus {
	results := make([]ServiceStatus, len(urls))
	ch := make(chan ServiceStatus)

	for _, url := range urls {
		go func(url string) {
			client := &http.Client{
				Timeout: time.Duration(timeoutSeconds) * time.Second,
			}

			req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)

			resp, err := client.Do(req)
			if err != nil {
				ch <- ServiceStatus{URL: url, Alive: false, Error: err.Error()}
				return
			}
			defer resp.Body.Close()

			ch <- ServiceStatus{URL: url, Alive: resp.StatusCode < 500, StatusCode: resp.StatusCode}
		}(url)
	}

	for i := range results {
		results[i] = <-ch
	}

	return results
}
