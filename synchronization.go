package main

import (
    "fmt"
    "time"
    "sync"
)

var (
    counter = 0
    lock sync.Mutex
)

func main() {
    for i := 0; i < 20; i++ {
        go func() {
            lock.Lock()
            defer lock.Unlock()

            counter++
            fmt.Printf("[+] Just increased the timer! %d->%d\n", counter-1, counter)
        }()
    }

    time.Sleep(time.Millisecond * 100)
}

