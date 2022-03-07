package main

import (
	"context"
	"fmt"
	"time"
)

type producer struct {
	ch       chan<- string
	ctx      context.Context
	interval time.Duration
	i        int
}

type consumer struct {
	ch       <-chan string
	ctx      context.Context
	interval time.Duration
}

func main() {
	message_c := make(chan string, 10)
	defer close(message_c)
	var deviceMap = make(map[int]string, 5)
	deviceMap[1] = "10.5.5.10"
	deviceMap[2] = "10.5.5.40"
	deviceMap[3] = "10.5.5.70"
	deviceMap[4] = "10.5.5.100"
	deviceMap[5] = "10.5.5.130"
	deviceMap[7] = "done"
	p := new(producer)
	ctx, cancel := context.WithTimeout(context.TODO(), 8*time.Second)
	defer cancel()
	p.ctx = ctx
	p.ch = message_c
	p.interval = time.Second
	p.i = 0
	c := &consumer{
		ch:       message_c,
		ctx:      ctx,
		interval: time.Second,
	}
	go p.product(deviceMap)
	go c.consume()
	time.Sleep(time.Second * 10)
	fmt.Println("main process is exit")

}

func (p *producer) product(message map[int]string) {
	ticker := time.NewTicker(p.interval)
	defer ticker.Stop()
	for _ = range ticker.C {
		select {
		case <-p.ctx.Done():
			fmt.Println("producer process is done <-")
			return
		default:
			if message[p.i] != "" {
				p.ch <- message[p.i]
				fmt.Printf("[produce message] %v\n", message[p.i])
			}
			p.i++

		}
	}
}

func (c *consumer) consume() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for _ = range ticker.C {
		select {
		case <-c.ctx.Done():
			fmt.Println("<- consumer process is done")
			return
		default:
			fmt.Printf("[consume message] %v\n", <-c.ch)

		}
	}
}
