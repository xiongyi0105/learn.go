package main

import (
	"context"
	"fmt"
	"time"
)

//ctx := context.WithValue(bgctx, "a", "b")
//go func(c context.Context) {
//	fmt.Println(c.Value("a"))
//}(ctx)
func main() {
	bgctx := context.Background()
	timeoutctx, cancel := context.WithTimeout(bgctx, time.Second*5)
	defer cancel()

	go func(ctx context.Context) {
		deadline, _ := ctx.Deadline()
		fmt.Println(ctx)
		fmt.Printf("Child Process's Deadline:%v\n", deadline)
		ticker := time.NewTicker(time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("Child process interrupt")
				return
			default:
				fmt.Println("enter default:", time.Now().String())
			}
		}
	}(timeoutctx)
	time.Sleep(10 * time.Second)
	select {
	case <-timeoutctx.Done():
		time.Sleep(5 * time.Second)
		fmt.Println("main process exit", time.Now().String())
	}
}
