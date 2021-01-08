package main

import (
	"context"
	"fmt"
)

func main() {
	// gen generates integers in a separate gouroutine and
	// sends them to the returner channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal gouroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the gouroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 10 {
			break
		}
	}
}
