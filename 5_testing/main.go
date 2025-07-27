package mai

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestDeadlineExceeded(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://example.com", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			t.Log("Request timed out as expected")
		} else {
			t.Fatalf("Request failed: %v", err)
		}
		return
	}
	defer resp.Body.Close()

	t.Log("Request completed successfully with status:", resp.Status)
}
