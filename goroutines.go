package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("[-] Start.")
    go process()

    time.Sleep(time.Millisecond * 1000)

    fmt.Println("[+] Done.")
}

func process() {
    fmt.Println("...processing...")
}
