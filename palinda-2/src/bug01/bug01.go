package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
// does not work due to chanel never being able to let go of the value
// causing th program to freeze on the ch <- "Hello world!" line
func main() {
	ch := make(chan string, 1)
	ch <- "Hello world!"
	fmt.Println(<-ch)
}

// creating a buffered chanel fixes it since the ch can letgo of the value
//and place it in the buffer (as long as the buffer is not full) and move on
//and print
