package main

import (
  "context"
  "fmt"
)

func main() {
  type favContextKey string

  f := func(ctx context.Context, k favContextKey) {
    if v := ctx.Value(k); v != nil {
      fmt.Println("found value:", v)
      return
    }
    fmt.Println("key not found:", k)
  }

  k := favContextKey("pizza")
  ctx := context.WithValue(context.Background(), k, "Margherita")

  f(ctx, k)
  f(ctx, favContextKey("burger"))
}
