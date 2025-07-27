package main

import (
	"context"
	"time"
)

func main() {
	// create a context
	ctx := context.Background()
	// create cancelable context with a timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel() // ensure resources are cleaned up
	myRoutine(ctx)
}

func myRoutine(ctx context.Context) {
	<-ctx.Done()
}
