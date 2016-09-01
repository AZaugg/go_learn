package main

import (
   "fmt"
   )
type Person struct {
    Name string
    Age int
}
func (p *Person) age(){
    p.Age += 10
}

func five(data *Person) {
    data.Age += 5
}

func main() {


    Andrew := &Person{
        "Andy", 25,
    }

    fmt.Printf("Name is %s and age is %d\n", Andrew.Name, Andrew.Age)
    five(Andrew)
    fmt.Printf("Name is %s and age in 5 pointer years is: %d\n", Andrew.Name, Andrew.Age)
    Andrew.age()
    fmt.Printf("Agew in 10 will be %d\n", Andrew.Age)

    George := new(Person)
    George.Name = "George"

    fmt.Printf("New Persons name is %s\n", George.Name)
}
