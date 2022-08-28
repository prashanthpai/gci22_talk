package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// START OMIT

func main() {
	srv := &http.Server{Addr: ":8083"}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("ListenAndServe() failed: %w", err)
		}
	}()

	c := make(chan os.Signal, 1) // buffered ch
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	log.Printf("received signal: %s", <-c)

	if err := srv.Shutdown(context.TODO()); err != nil {
		log.Printf("Shutdown() failed: %w", err)
	}

	wg.Wait()
}

// END OMIT
