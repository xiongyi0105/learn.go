package main

import "fmt"

//声明的空切片是长度容量为0的切片，且不会引用数组。是一个nil切片
func main() {
	var nilslice []int
	fmt.Println(nilslice)
	if nilslice == nil {
		fmt.Println("This is a nil slice!") //空切片是nil ，且切片只能==nil来比较，除非可以使用DeepEqual函数
	}
	fmt.Println("len为 ：", len(nilslice)) //长度容量都为0
	fmt.Println("cap为：", cap(nilslice))
}
