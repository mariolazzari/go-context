package main

import (
	"context"
	"net/http"
	"time"
)

func main() {
	// creates a new context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // ensure the cancel function is called to release resources

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://example.com", nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	_, err := client.Do(req)

}
