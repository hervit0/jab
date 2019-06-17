package main

import (
    "fmt"
    "sync"
	// "time"
)

func threadedFunction(upperBound int) {
    fmt.Println("Upper bound - not storing:", upperBound)
    resultSlice := make([]int, upperBound)

    // for i := 0; i < upperBound; i++ {
    for index, _ := range resultSlice {
        go func(current int) int {
            result := current * 2
            fmt.Println("Yo:", result)
            return result
        }(index)
    }
}

func threadedFunctionWithChan(sendingChan chan<- int, waitingGroup *sync.WaitGroup, upperBound int) {
    fmt.Println("Upper bound - sending to chan:", upperBound)
    resultSlice := make([]int, upperBound)

    // for i := 0; i < upperBound; i++ {
    for index, _ := range resultSlice {
        waitingGroup.Add(1)

        go func(current int) int {
            defer waitingGroup.Done()

            result := current * 2
            fmt.Println("Yo sending:", result)
            sendingChan <- result
            return result
        }(index)
    }
}

func sayhello(count int, wg *sync.WaitGroup, messages *chan int) {
	defer wg.Done()
    r := count * 2

	fmt.Println("[Say Hello Routine] result computed: ", r)
	*messages <- r
	fmt.Println("[Say Hello Routine] sent message: ", r)
}

func main() {
    // messages := make(chan string, 42) Buffered to 42 messages
    messages := make(chan string)
    go func() { messages <- "ping" }()
    msg := <-messages
    fmt.Println(msg)

    threadedFunction(5)
    fmt.Scanln()

    // Another one
    var waitingGroup sync.WaitGroup
    messagesChan := make(chan int)
    threadedFunctionWithChan(messagesChan, &waitingGroup, 5)

    go func(wg *sync.WaitGroup, messages chan int) {
		fmt.Println("[Closer routine] Waiting")
		wg.Wait()
        fmt.Println("[Closer routine] Done waiting")

		close(messagesChan)
    }(&waitingGroup, messagesChan)
    fmt.Scanln()

    // Another one
    var wg sync.WaitGroup
    msgChan := make(chan int)
    s := make([]int, 10)

	for x := 1; x <= len(s); x++ {
        wg.Add(1)
		go sayhello(x, &wg, &msgChan)
    }

    go func(wg *sync.WaitGroup, msgChan chan int) {
		fmt.Println("[Closer routine] Waiting")
		wg.Wait()
        fmt.Println("[Closer routine] Done waiting")
		close(msgChan)
    }(&wg, msgChan)

    // for msg := range msgChan {
        // Interesting behavior with a timer - it's a blocking op
	    // time.Sleep(time.Millisecond * time.Duration(1000 * 5))
    //     fmt.Println("[Channel Reader]", msg)
    // }
    select {
    case msg := <-msgChan:
        fmt.Println("[Select Fire] received: ", msg)
    default:
        fmt.Println("[Select Fire] Nein!")
    }

    // msgCollection := <-messagesChan
    // waitingGroup.Wait()
    // close(messagesChan)

    // fmt.Println("Collected messages", msgCollection)
}
