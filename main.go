package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	ctx, cancelCtx := context.WithTimeout(ctx, 1500*time.Millisecond)
	defer cancelCtx()

	printCh := make(chan int)
	go doAnother(ctx, printCh)

	for i := 0; i < 10; i++ {
		select {
		case printCh <- i:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			return
		}
	}

	cancelCtx()

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Doing something finished")

}
func doAnother(ctx context.Context, printCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Println("do another error:", err)
			}
			fmt.Println("do another finished")
			return
		case num := <-printCh:
			fmt.Println("do another:", num)
		}
	}
}

func main() {
	ctx := context.Background()

	doSomething(ctx)
}
