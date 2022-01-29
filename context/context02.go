package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*2)

	go func(ctx context.Context) {
		t := time.NewTicker(1 * time.Second)
		for _ = range t.C {
			select {
			case <-ctx.Done():
				fmt.Println("1")
				return
			default:
				fmt.Println("default")

			}
		}

	}(ctx)

	time.Sleep(4 * time.Second)

	select {
	case <-ctx.Done():
		fmt.Println("exist")
	}
}
