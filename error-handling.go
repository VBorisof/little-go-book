package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) != 2 {
        panic("[!] Supply exactly one argument.")
        os.Exit(1)
    }

    n, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Printf("[!] Cannot convert %s to a number. Sorry eh, try again. Error: %v\n", os.Args[1], err)
    } else {
        fmt.Printf("[+] Your number: %d\n", n)
    }
}

