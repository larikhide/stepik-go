package main

import (
	"fmt"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	resChan := make(chan int)

	go func() {
		defer close(resChan)
		select {
		case <-stopChan:
			return
		case num := <-firstChan:
			resChan <- num * num
		case num := <-secondChan:
			resChan <- num * 3
		default:
			break
		}
	}()
	return resChan
}

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	stop := make(chan struct{})

	r := calculator(ch1, ch2, stop)
	//ch1 <- 3
	//ch2 <- 3
	//close(stop)
	fmt.Println(<-r)

}
