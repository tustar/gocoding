package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown")
	select {
	case <-time.After(10 * time.Second):
		// Do nothing
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}

	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
