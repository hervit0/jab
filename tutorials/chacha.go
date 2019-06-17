package main

import (
	"fmt"
	"sync"
	"time"
	"log"
)

func main() {
	wg := new(sync.WaitGroup)
	messages := make(chan string)
	for x := 1; x <= 5; x++ {
		wg.Add(1)
		go sayhello(x, wg, &messages)
	}

	go func(wg *sync.WaitGroup, messages chan string) {
		log.Println("waiting")
		wg.Wait()
		log.Println("done waiting")
		close(messages)
	}(wg, messages)

	for msg := range messages {
		fmt.Println(msg)
	}
}

func sayhello(count int, wg *sync.WaitGroup, messages *chan string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(1000 * count))
	*messages <- fmt.Sprintf("hello: %d", count)
	log.Println("sent message ", count)
}
