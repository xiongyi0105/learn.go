package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 1, 0, 8, 10}
	var minValue = min(slice...) // 此处传参时也要使用 x ...
	fmt.Println(minValue)
}

//传递变长参数  x ...视作一个类似slice结构的参数
func min(s ...int) (min int) {
	//使用命名式返回值，会默认初始化声明变量为0值，若是非命名返回值需要新声明变量
	min = s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return
}
