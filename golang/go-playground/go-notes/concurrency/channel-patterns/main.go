package main

import (
	"fmt"
	"math/rand"
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

func main() {
	waitForTask()
}
