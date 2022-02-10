package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		ch <- "data"
		fmt.Println("child: sent signal")
	}()

	d := <-ch
	fmt.Println("parent: recv'd signal: ", d)
	time.Sleep(time.Second) // make sure child task finished
}

func fanOut() {
	children := 2000
	ch := make(chan string, children) // won't block the senders

	for c := 0; c < children; c++ {
		go func(child int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "data"
			fmt.Printf("child %d: sent signal\n", child)
		}(c)
	}

	for children > 0 {
		d := <-ch
		children--
		fmt.Println("parent: recv'd ", d)
	}
	time.Sleep(time.Second)
}

func waitForTask() {
	ch := make(chan string)
	go func() {
		d := <-ch
		fmt.Println("child: recv'd ", d)
	}()

	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	ch <- "data"
	fmt.Println("parent: sent signal")
	time.Sleep(time.Second)
}

func pooling() {
	ch := make(chan string)
	grs := runtime.GOMAXPROCS(0) // query available thread numbers
	for g := 0; g < grs; g++ {
		go func(child int) {
			for d := range ch {
				fmt.Printf("child %d: recv'd signal: %s\n", child, d)
			}
			fmt.Printf("child %d: recv'd shutdown signal\n", child)
		}(g)
	}

	const work = 100
	for w := 0; w < work; w++ {
		ch <- "data"
		fmt.Println("parent: sent signal ", w)
	}
	close(ch)
	fmt.Println("parent: sent shutdown signal")
	time.Sleep(time.Second)
}

func drop() {
	const cap = 100
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("child : recv'd signal :", p)
		}
	}()

	const work = 2000
	for w := 0; w < work; w++ {
		select {
		case ch <- "data":
			fmt.Println("parent : sent signal :", w)
		default:
			fmt.Println("parent : dropped data :", w)
		}
	}
	close(ch)
	fmt.Println("parent : sent shutdown signal")
	time.Sleep(time.Second)
}

func cancellation() {
	duration := 100 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)
	go func() {
		// The child Goroutine needs to be able to send its signal,
		// whether or not the parent Goroutine is around to receive it.
		// If a non-buffered channel is used,
		// the child Goroutine will block forever and become a memory leak.
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		ch <- "data"
	}()

	select {
	case d := <-ch:
		fmt.Println("work complete", d)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
	time.Sleep(time.Second)
}

func main() {
	cancellation()
}
