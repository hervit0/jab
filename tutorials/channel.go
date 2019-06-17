package main

import (
    "fmt"
    "sync"
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

func main() {
    // messages := make(chan string, 42) Buffered to 42 messages
    messages := make(chan string)
    go func() { messages <- "ping" }()
    msg := <-messages
    fmt.Println(msg)

    threadedFunction(5)
    fmt.Scanln()

    messagesChan := make(chan int)
    var waitingGroup sync.WaitGroup
    threadedFunctionWithChan(messagesChan, &waitingGroup, 5)

    go func(wg *sync.WaitGroup, messages chan string) {
		fmt.Println("waiting")
		wg.Wait()
		fmt.Println("done waiting")
		close(messagesChan)
    }(&waitingGroup, messages)

    // for msg := range messagesChan {
    //     fmt.Println(msg)
    // }

    // msgCollection := <-messagesChan
    // waitingGroup.Wait()
    // close(messagesChan)

    // fmt.Println("Collected messages", msgCollection)
}
