package main

import "fmt"

//闭包实现斐波那契函数
func main() {
	fibo := closure()
	for i := 0; i < 10; i++ {
		ll := fibo(i)
		fmt.Println(ll)
	}

}

// 闭包
func closure() func(num int) int {
	var (
		i int = 1
		k int = 0
		j int = 0
	)
	return func(num int) int {
		j = k + i
		defer func() {
			i, k = j, i
		}()
		return j
	}
}
