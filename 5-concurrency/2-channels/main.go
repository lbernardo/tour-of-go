package main

import "fmt"

/**
ch <- v // Send v to channel ch
v := <- ch // Receive from ch and assign value to v
*/

func testChannel(c chan string) {
	for i := 0; i < 1000; i++ {
		c <- fmt.Sprintf("Ballon")
	}
}

func main() {
	c := make(chan string)
	go testChannel(c)
	value := <-c
	fmt.Println(value)
}
