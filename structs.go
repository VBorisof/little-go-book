package main

import (
    "fmt"
)

type Person struct {
    Name string
}

func (p *Person) Introduce() {
    fmt.Printf("Hi! I am %s\n", p.Name)
}

type Saiyan struct {
    *Person
    Power int
}

func NewSaiyan(name string, power int) *Saiyan {
    return &Saiyan{
        Person: &Person { name },
        Power: power,
    }
}

func (s *Saiyan) Super() {
    s.Power += 10000
}

func (s *Saiyan) GetPower() int {
    return s.Power
}

func main() {
    goku := NewSaiyan("Goku", 9000)

    goku.Person.Introduce()
    fmt.Printf("%s's power is %d\n", goku.Name, goku.GetPower())
    goku.Super()
    fmt.Printf("%s's power is %d\n", goku.Name, goku.GetPower())
}

