package main

import (
  "fmt"
)

func atoi(input string) int {
  result := 0
  for _, c := range input {
    fmt.Println(c)
  }
  return result
}

func main(){
  atoi("a")
}
