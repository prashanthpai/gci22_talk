package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// START OMIT

func producer(wg *sync.WaitGroup, ch chan<- int, stopCh <-chan struct{}) {
	defer wg.Done()

	for i := 0; ; i++ {
		select {
		case <-stopCh:
			return
		default:
			ch <- i
		}
	}
}

func consumer(wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()

	for i := range ch {
		fmt.Println(i)
	}
}

// END OMIT

func main() {
	ch := make(chan int, 10)
	stopCh := make(chan struct{})

	wgc := sync.WaitGroup{}
	wgc.Add(10)
	for i := 0; i < 10; i++ {
		go consumer(&wgc, ch)
	}

	wgp := sync.WaitGroup{}
	wgp.Add(5)
	for i := 0; i < 5; i++ {
		go producer(&wgp, ch, stopCh)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	<-sigCh
	close(stopCh)

	wgp.Wait()
	close(ch)
	wgc.Wait()
}

// END OMIT2
