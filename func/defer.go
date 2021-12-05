package main

import "fmt"

func a() {
	i := 0
	//defer相当于finally但可以在结构体内灵活使用
	defer fmt.Println(i)
	i++
	return
}

//当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func main() {
	a()
	f()
}
