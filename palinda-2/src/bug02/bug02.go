package main

import (
	"fmt"
	"sync"
	"time"
)

// This program should go to 11, but it seemingly only prints 1 to 10.
// this program does not count the final value because main closes before the
// print function has time to print
// Fixed by cereating a wait group so main waits for the go print thread to finish
func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go Print(ch, &wg)

	for i := 1; i <= 11; i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()

}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}

}
