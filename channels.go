package main

import "fmt"
import "time"

func main() {

  messages := make(chan string)

  go func() {
    time.Sleep(1*time.Second)
    fmt.Println("> Sending message to channel...")
    messages <- "ping"
    fmt.Println("Sent.")
    time.Sleep(1*time.Second)
  }()

  fmt.Println("> Receiving message from another goroutine via channel...")
  msg := <- messages
  fmt.Println("Received.")
  fmt.Println(msg)
}
