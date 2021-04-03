package main

func main() {
    print("Mood? ")
    if mood := getMood(); mood > 8 {
        println("Fucking perfect")
    } else if mood > 5 {
        println("Eh, okay")
    } else {
        println("Sadboy")
    }
}

func getMood() int {
    return 10
}
