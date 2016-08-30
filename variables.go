package main

import(
  "fmt"
  "reflect"
  )

func main() {
    var number int
    var text string

    number = 20
    text = "mytext"

    fmt.Printf("My variable is %d\n", number)
    fmt.Printf("My bool var is is %s\n", text)

    infernumber := 20
    inferstring := "my infered string"

    fmt.Println(reflect.TypeOf(infernumber))
    fmt.Println(reflect.TypeOf(inferstring))

    name, age := "Andy", 20

    fmt.Printf("Name is %s, age is %d", name, age)
}
