package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var fmeng = false

func produce(threadid int, wg *sync.WaitGroup, ch chan string) {
	count := 0
	for !fmeng {
		time.Sleep(time.Second)
		count++
		data := strconv.Itoa(threadid) + "--------" + strconv.Itoa(count)
		fmt.Printf("produce,%s\n", data)
		ch <- data
	}
	wg.Done()
}

func consume(wg *sync.WaitGroup, ch chan string) {
	for data := range ch {
		time.Sleep(time.Second)
		fmt.Printf("consume,%s\n", data)
	}
	wg.Done()
}

func main() {
	chansteam := make(chan string, 10)
	prodwg := new(sync.WaitGroup)
	conswg := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		prodwg.Add(1)
		go produce(i, prodwg, chansteam)
	}
	for j := 0; j < 2; j++ {
		conswg.Add(1)
		go consume(conswg, chansteam)
	}
	go func() {
		time.Sleep(time.Second * 3)
		fmeng = true
	}()
	prodwg.Wait()
	close(chansteam)
	conswg.Wait()
}
