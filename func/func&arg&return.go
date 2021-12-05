package main

import "fmt"

var num = 10
var num1, num2 int

func main() {
	num1, num2 = getReturnvalue1(num)
	PrintValue()
	num1, num2 = getReturnvalue2(num)
	PrintValue()
	//引用类型传参可以改变实参的值，而值类型的传参只会复制一份实参
	num1, num2 = getReturnValueAndChangearg(&num)
	PrintValue()
}

func getReturnvalue1(input int) (int, int) {
	return input * 2, input * 3
}

func getReturnvalue2(input int) (x1, x2 int) {
	x1 = input * 2
	x2 = input * 3
	return
}

func getReturnValueAndChangearg(input *int) (x1, x2 int) {
	//改变被引用量的值
	*input = *input * 2
	x1 = *input * 2
	x2 = *input * 3
	return
}

func PrintValue() {
	fmt.Printf("ReturnValue is %d and %d\n", num1, num2)
	fmt.Println("Now num is:", num)
}
