package main

import (
  "fmt"
  "os"
)

func main() {
    fmt.Println("Length", len(os.Args))

    if len(os.Args) < 2 {
        fmt.Println("Not enough args")
    } else {
        fmt.Println("Arg 1 is", os.Args[1])
        fmt.Println("Arg 0 is", os.Args[0])
    }
}

