// вы уже внутри main()

package main

import "fmt"

func main() {
	done := make(chan struct{})
	go func(done chan struct{}) {
		work()
		close(done)
	}(done)
	<-done
}

func work() {
	fmt.Println("Shark: do something")
	fmt.Println("Mouse: i'm...i did")
}
