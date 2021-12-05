package main

import "fmt"

// 值类型的变量 直接指向内存中的值
// 将值类型的变量赋值给另一个变量相当于是在内存中的值进行了拷贝，地址改变
// 值类型的变量的值存储在栈中。

func main() {
	a := 7
	var i int
	i = a
	fmt.Println(&a) // 0xc0000aa058
	fmt.Println(&i) // 0xc0000aa070
	a = 8
	fmt.Println(a)
	fmt.Println(i)

}
