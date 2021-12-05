package main

import "fmt"

//主函数调用一个使用序列个数作为参数的函数，该函数返回一个大小为序列个数的 Fibonacci 切片。

func main() {
	fibonacciInstance := Fibonacci(10)
	fmt.Println(fibonacciInstance)
}

func Fibonacci(numLen int) []int {
	sliceFibo := make([]int, numLen)
	for i := 0; i < numLen-1; i++ {
		if i == 0 {
			sliceFibo[i+1] = i + 1
			continue
		}
		// 从下标为1的切片元素开始每个元素的后一个值等于现在加之前
		sliceFibo[i+1] = sliceFibo[i] + sliceFibo[i-1]
	}
	return sliceFibo
}
