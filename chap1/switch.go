package main

import (
  "fmt"
  "runtime"
)

func main(){
  switch os:= runtime.GOOS; os {
  case "darwin":
      fmt.Println("Mac OS")
  case "linux":
      fmt.Println("Linux")
  default:
      fmt.Println("%s.", os)
  }
}
