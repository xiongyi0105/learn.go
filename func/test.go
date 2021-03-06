package main

import (
	"bytes"
	"fmt"
	"math/big"
	"strconv"
)

/*
练习 6.4
重写本节中生成斐波那契数列的程序并返回两个命名返回值（详见第 6.2 节），即数列中的位置和对应的值，例如 5 与 4，89 与 10。
练习 6.5
使用递归函数从 10 打印到 1。
练习 6.6
实现一个输出前 30 个整数的阶乘的程序。
n! 的阶乘定义为：n! = n * (n-1)!, 0! = 1，因此它非常适合使用递归函数来实现。
然后，使用命名返回值来实现这个程序的第二个版本。
特别注意的是，使用 int 类型最多只能计算到 12 的阶乘，因为一般情况下 int 类型的大小为 32 位，继续计算会导致溢出错误。那么，如何才能解决这个问题呢？
最好的解决方案就是使用 big 包（详见第 9.4 节）。
*/

var buffer *bytes.Buffer = &bytes.Buffer{}
var sum int64 = 1
var bigSum = big.NewInt(sum)

func main() {

	num := 10
	ll := for10(&num)
	fmt.Println(ll)
	x := factorialForthirty(30)
	fmt.Printf("x:%v", x)
}

func for10(num *int) string {
	fmt.Println(*num)
	buffer.WriteString(strconv.Itoa(*num))
	if *num == 0 {
		return buffer.String()
	}
	*num -= 1
	for10(num)
	return buffer.String()
}

func factorialForthirty(factorialNum int64) (res *big.Int) {
	bigFactorialNum := big.NewInt(factorialNum)
	if factorialNum == 0 {
		return
	}
	bigSum.Mul(bigSum, bigFactorialNum)
	fmt.Println(bigSum)
	res = bigSum
	factorialForthirty(factorialNum - 1)
	return
}
