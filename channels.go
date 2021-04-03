package main

import (
    "fmt"
    "time"
    "math/rand"
)

type Worker struct {
    id int
}

func (w Worker) process(c chan int) {
    for {
        data := <-c
        fmt.Printf("%d-", data);
    }
}

func main() {
    // Seed the RNG
    rand.Seed(time.Now().UTC().UnixNano())

    c := make(chan int)
    for i := 0; i < 100; i++ {
        worker := &Worker { id: i }
        go worker.process(c)
    }

    for {
        randInt := rand.Int()
        select {
        case c <- randInt:
            // All good, do nothing
        case <- time.After(time.Millisecond * 100):
            fmt.Printf("\n[!] Timeout for message: %d\n", randInt)
        default:
            fmt.Printf("\n[!] Message dropped: %d\n", randInt)
        }

        time.Sleep(time.Millisecond * 50)
    }
}

