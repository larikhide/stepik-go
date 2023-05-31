// вы уже внутри main()

package main

import (
	"fmt"
	"time"
)

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
	time.Sleep(2 * time.Second)
	fmt.Println("Mouse: i'm...")
	time.Sleep(2 * time.Second)
	fmt.Println("Mouse: i did")
}
