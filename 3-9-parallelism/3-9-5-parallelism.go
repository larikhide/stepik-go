package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go work(wg)
	}
	wg.Wait()

}

func work(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Shark: do something")
	time.Sleep(2 * time.Second)
	fmt.Println("Mouse: i'm...")
	time.Sleep(2 * time.Second)
	fmt.Println("Mouse: i did")

}
