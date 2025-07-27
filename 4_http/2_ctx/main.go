package main

import (
	"context"
	"net/http"
)

func main() {
	http.Handle("/", contextMiddleware(http.HandlerFunc(requestHandler)))
	// Start the HTTP server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}

func contextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Add a value to the context
		ctx = context.WithValue(ctx, "key", "value")
		// Pass the modified context to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Retrieve the value from the context
	if val, ok := ctx.Value("key").(string); ok {
		w.Write([]byte("Value from context: " + val))
	} else {
		w.Write([]byte("No value found in context"))
	}
}
