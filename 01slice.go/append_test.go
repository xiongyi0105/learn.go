package main

import "fmt"

/*
练习 7.9
给定 slice s[]int 和一个 int 类型的因子 factor，扩展 s 使其长度为 len(s) * factor。
练习 7.10
用顺序函数过滤容器：s 是前 10 个整型的切片。构造一个函数 Filter，第一个参数是 s，第二个参数是一个 fn func(int) bool，返回满足函数 fn 的元素切片。通过 fn 测试方法测试当整型值是偶数时的情况。
练习 7.11
写一个函数 InsertStringSlice 将切片插入到另一个切片的指定位置。
练习 7.12
写一个函数 RemoveStringSlice 将从 start 到 end 索引的元素从切片 中移除。
*/

func main() {
	mySlice := AppendInt([]int{1, 3, 6, 7}, 9, 5, 1, 4, 3)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
}

func AppendInt(slice []int, factor ...int) (sliceReturn []int) {
	oldSliceLen := len(slice)
	lenOfFactor := len(factor)
	newSliceLen := oldSliceLen * lenOfFactor
	newSlice := make([]int, newSliceLen)
	copy(newSlice, slice)
	slice = newSlice
	appendFactorLen := oldSliceLen + lenOfFactor
	copy(slice[oldSliceLen:appendFactorLen], factor)
	return slice
}
