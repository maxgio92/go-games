package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

  ctx := req.Context()
  fmt.Println("> Server: Hello handler started")
  defer fmt.Println("> Server: Hello handler ended")

  select {
  case <-time.After(2 * time.Second):
    fmt.Println(w, "hello\n")
  case <-ctx.Done():
    err := ctx.Err()
    fmt.Println("server:", err)
    internalError := http.StatusInternalServerError
    http.Error(w, err.Error(), internalError)
  }
}

func main() {
  http.HandleFunc("/hello", hello)
  http.ListenAndServe(":8090", nil)
}
