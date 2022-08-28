package main

import (
	"context"
	"os/signal"
	"sync"
	"syscall"
)

// START OMIT

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go work(ctx, wg)
	}

	wg.Wait()
}

func work(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// some work
		}
	}
}

// END OMIT
