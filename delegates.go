package main

import "fmt"

type Adder func (a int, b int) int

func GetSum(adder Adder) int {
    return adder(1, 2)
}


func main() {
    fmt.Println(GetSum(func(a int, b int) int {
        return a + b
    }))
}
