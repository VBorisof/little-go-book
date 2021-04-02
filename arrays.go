package main

import (
    "fmt"
    "strings"
    crypto_rand "crypto/rand"
    math_rand "math/rand"
    "encoding/binary"
    "sort"
)

func main() {
    // Some capacity flexes
    println("---------------\nCapacity tests...")
    PrintScores()

    // Using the 2-arg make() version makes the slice length
    // default to it's capacity. Use the 3-arg version to
    // specify the initial size as a second argument.
    println("---------------\nSome other append() tests...")
    scores := make([]int, /*0,*/ 5)
    scores = append(scores, 12)
    fmt.Println("Scores: ", scores)

    // When you get a slice, you actually get the portion 
    // of that array, pointing at the same values -- not a copy.
    println("---------------\nSlicing tests...")
    someNums := []int { 1, 2, 3, 4, 5, 6 }
    slice := someNums[2:4]
    slice[0] = 999
    fmt.Println(someNums)

    // Find first space after some length:
    someString := "The quick brown fox jumps over lazy dog."
    firstSpaceIdx := strings.Index(someString[5:6], " ")
    fmt.Printf("First space index: %d\n", firstSpaceIdx)

    println("---------------\nRemoveAtIndex() test.")
    source := []int { 1, 2, 3, 4, 5 }
    source = RemoveAtIndex(source, 4)
    fmt.Println(source)

    println("---------------\ncopy() tests.")
    CopyTests()

    println("---------------\nDone boy.")
}

func PrintScores() {
    scores := make([]int, 0, 10)

    c := cap(scores)

    fmt.Printf("Initial Capacity: %d\n", c)

    for i := 0; i < 25; i++ {
        scores = append(scores, i)

        if (cap(scores) != c) {
            fmt.Printf("Capacity has changed: %d -> %d\n", c, cap(scores))

            c = cap(scores)
        }
    }

    for index, value := range scores {
        fmt.Printf("[%d] -- %d\n", index, value)
    }
}

// Removes an element found at given index in the source slice.
// Most probably will mess up the order of the slice.
func RemoveAtIndex(source []int, index int) []int {
    lastIndex := len(source) - 1
    source[index], source[lastIndex] = source[lastIndex], source[index]
    return source[:lastIndex]
}

func CopyTests() {
    // Seed the rng

    // KISS: Seed with time
    // rand.Seed(time.Now().UTC().UnixNano())

    // More involved: seed without relying on nanosecond
    // precision, by using crypto/rand.
    var bytes [8]byte
    _, err := crypto_rand.Read(bytes[:])
    if err != nil {
        panic("[!] Cannot seed the RNG! Aborting CopyTests()")
        return
    }
    math_rand.Seed(int64(binary.LittleEndian.Uint64(bytes[:])))

    // Generate 100 random ints [0, 1000)
    scores := make([]int, 100)
    for i := 0; i < 100; i++ {
        scores[i] = math_rand.Intn(1000)
    }
    sort.Ints(scores)

    // Make a new slice of 5 lowest ints
    lowest := make([]int, 5)
    copy(lowest, scores[:5])
    fmt.Println(lowest)
}



