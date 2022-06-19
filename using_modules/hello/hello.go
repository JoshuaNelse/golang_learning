package main

import (
  "fmt"
  "log"

  "example/greeting"
)

func main() {
  log.SetPrefix("hello module: ")
  log.SetFlags(0)

  message, err := greeting.Hello("Josh")

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(message)
}
