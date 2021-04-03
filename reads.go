package main

import (
    "fmt"
    "io"
)

func main() {
    var input int
    _, err := fmt.Scan(&input)
    if err == io.EOF {
        fmt.Printf("No more input after %s", input)
    }
}
