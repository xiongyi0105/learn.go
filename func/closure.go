package main

import (
	"fmt"
)

var x = 0

func main() {

	p := closure()
	x1 := p(10)
	x2 := p(20)
	x3 := p(30)
	fmt.Println(x1 + x2 + x3)
	fib := fibo()
	ll := fib(1)
	l1 := fib(2)
	l2 := fib(3)
	l3 := fib(4)
	fmt.Println(ll)
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(l3)
}

func closure() func(int) int {
	//var x = 0
	return func(delta int) int {
		x += delta
		return x
	}

}

func fibo() func(i int) int {
	var c = 1
	var j = 0
	return func(i int) int {
		if i < 10 {
			defer func() {

				fmt.Println(j)
			}()
			c += j
		}
		return c
	}
}
