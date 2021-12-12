package main

import (
	"errors"
	"fmt"
)

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

	// 7.9
	mySlice := AppendInt([]int{1, 3, 6, 7}, 9, 5, 1, 4, 3)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))

	// 7.10
	fn := func(i int) bool {
		if i%2 == 0 {
			return true
		}
		return false
	}
	newSlice := Filter([]int{96, 12, 53, 123, 54, 34, 13, 16, 75, 90}, fn)
	fmt.Println(newSlice)

	// 7.11
	ll, err := InsertStringSlice([]int{1, 3, 4, 6, 123, 1233, 123, 123, 8, 123, 42365, 123, 11}, []int{1, 2, 0, 4, 3, 8, 1, 3, 5, 7, 8}, -1)
	fmt.Println(ll, err)

	//	7.12
	afterSlice := RemoveStringSlice([]string{"熊易", "杨苏文", "谢谨伊", "严加文", "hjl"}, 2, 4)
	fmt.Println(afterSlice)
}

func AppendInt(slice []int, factor ...int) []int {
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

func Filter(s []int, fn func(int) bool) []int {

	newSlice := make([]int, 0)
	for _, intValue := range s {
		if fn(intValue) {
			newSlice = append(newSlice, intValue)
		}
	}
	return newSlice

}

func InsertStringSlice(insertSlice []int, oldSlice []int, start int) ([]int, error) {
	end := len(insertSlice) + start
	if start < len(oldSlice) && start >= 0 {
		finalslice := make([]int, len(oldSlice)+len(insertSlice))
		copy(finalslice[:start], oldSlice[:start])
		copy(finalslice[start:end], insertSlice)
		copy(finalslice[end:], oldSlice[start:])
		return finalslice, nil
	}
	return []int{}, errors.New("The start index is out of index")

}
func RemoveStringSlice(stringSlice []string, start int, end int) []string {
	fmt.Println(stringSlice[end:])
	newStringSlice := append(stringSlice[:start], stringSlice[end:]...)
	stringSlice = newStringSlice
	return stringSlice
}
