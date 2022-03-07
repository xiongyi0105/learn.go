package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "/tmp/cpuprofile", "write cpu profile to file")

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err.Error())
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	//messages := make(chan string, 11)
	//defer close(messages)
	bgctx := context.Background()
	ctx := context.WithValue(bgctx, "a", "b")
	go func(c context.Context) {
		fmt.Println(c.Value("a"))
	}(ctx)
	timeoutctx, cancel := context.WithTimeout(bgctx, time.Second*30)
	defer cancel()
	go func(ctx context.Context) {
		deadline, _ := ctx.Deadline()
		fmt.Println(ctx)
		/*
			contextName(c.cancelCtx.Context) + ".WithDeadline(" +
					c.deadline.String() + " [" +
					time.Until(c.deadline).String() + "])"
		*/
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
	time.Sleep(15 * time.Second)
	select {
	case <-timeoutctx.Done():
		time.Sleep(6 * time.Second)
		fmt.Println("main process exit", time.Now().String())
	}

}
