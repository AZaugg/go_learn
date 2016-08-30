package main

import (
  "fmt"
  )


func add(num1 int, num2 int) int {
    return num1 + num2
}

func isgreater(num int) bool {
    if num > 10 {
        return true
    }
    return false
}

func throwaway() (bool, string) {
    return true, "aaa"
}

func main() {
  z := add(1,1)

  if isgreater(z) {
      fmt.Printf("Number is greater than 10")
  }

  _, text := throwaway()
  fmt.Printf("\nValue is %s\n", text)

  fmt.Printf("The added numbers are: %d", z)
}
