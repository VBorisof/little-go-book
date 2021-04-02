package main

import (
    "fmt"
)

func main() {
    m := map[string]int {
        "goku": 9001,
        "gohan": 2044,
        "vegeta": 1,
    }

    fmt.Printf("Map capacity: %d\n", len(m))

    // Interesting: The order in which the map will be printed is random,
    // to prevent DOS: https://github.com/golang/go/issues/2630
    for key, value := range m {
        fmt.Printf("%s, power %d\n", key, value)
    }
}
