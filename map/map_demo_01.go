package main

import "fmt"

func main() {
	var week_map = make(map[int]string)
	var week_slice = [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Satday", "Sunday"}
	for number, word := range week_slice {
		week_map[number+1] = word
	}
	for number, word := range week_map {
		fmt.Printf("This is digital week :%d and this is string week day %s\n", number, word)
	}
	// 判断key 7 存不存在，若值为空时，都会返回空，故要都过ok判断bool值
	if value, ok := week_map[7]; ok {
		fmt.Printf("key:7 is present and value is %v\n", value)
	}
	// 删除元素7
	delete(week_map, 7)
	if value, ok := week_map[7]; !ok {
		fmt.Printf("key:7 is not present and value is %#v\n", value)
	}
}
