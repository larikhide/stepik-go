package main

import "fmt"

func task2(ch chan string, s string) {
	for i := 0; i < 5; i++ {
		ch <- s + " "
	}
}

func main() {
	ch := make(chan string)

	go task2(ch, "Hey, DJ!")

	for i := 0; i < 5; i++ {
		result := <-ch
		fmt.Print(result)
	}
}
