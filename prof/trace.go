package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

//go run trace.go 2> trace.out
//go tool trace trace.out
func main() {
	fmt.Println("hi")
	trace.Start(os.Stderr)
	defer trace.Stop()
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int)
	const C int = 5
	go func() {
		defer wg.Done()
		for i := 0; i < C; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < C; i++ {
			v := <-ch
			_ = v
		}

	}()
	wg.Wait()

}
