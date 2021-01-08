package main

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"
  "time"
)

func main() {

  sigs := make(chan os.Signal, 1)
  done := make(chan bool, 1)

  signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

  go func() {
    sig := <-sigs
    fmt.Println()
    fmt.Println("Signal received")
    fmt.Println(sig)
    fmt.Println()
    time.Sleep(2*time.Second)
    done <- true
  }()

  fmt.Println("> Starting")
  time.Sleep(1*time.Second)
  fmt.Println("Done.")
  fmt.Println("> Awaiting signal")
  time.Sleep(1*time.Second)

  <- done

  time.Sleep(1*time.Second)
  fmt.Println("> Exiting")
  time.Sleep(1*time.Second)
  fmt.Println("Done.")
}
