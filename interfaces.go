package main

import (
    "fmt"
    "time"
)

type ILogger interface {
    Info(message string)
    Debug(message string)
}

type ConsoleLogger struct { }
func (cl *ConsoleLogger) Debug(message string) {
    fmt.Printf("[D][%v] %s\n", time.Now(), message)
}
func (cl *ConsoleLogger) Info(message string) {
    fmt.Printf("[I][%v] %s\n", time.Now(), message)
}


func Write(l ILogger, message string) {
    l.Info(message)
    l.Debug(message)
}

func main() {
    logger := &ConsoleLogger { }
    Write(logger, "Ssup")
}
