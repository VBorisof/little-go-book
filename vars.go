package main

import (
    "fmt"
)

func main() {
    _, isGoruValid := getPowerOf("Goru")
    if (isGoruValid) {
        fmt.Printf("Goru's power is defined\n")
    }
    _, isGokuValid := getPowerOf("Goku")
    if (isGokuValid) {
        fmt.Printf("Goku's power is defined: %d\n")
    }
}

func getPower() int {
    return 9001
}

func getPowerOf(name string) (int, bool) {
    if name == "Goku" {
        return 9000, true
    }

    return -1, false
}

