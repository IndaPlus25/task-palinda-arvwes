package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	p := 0
	c := 1
	n := 0
	fn := func() int {
		n = p + c
		p = c
		c = n
		return n
	}
	return fn
}

func fibo() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
