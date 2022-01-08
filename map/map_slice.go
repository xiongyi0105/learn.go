package main

import "fmt"

// map slice需要使用两次make，第一次分配切片，第二次分配每个map元素
func main() {
	var items = make([]map[string]int, 5)
	for i := range items {
		// 此时需要通过索引拿出切片里的值，不然拿到的是切片的副本不会实际改变原始底层数据。
		items[i] = make(map[string]int)
		items[i]["hello,this is key1"] = 1
		items[i]["hello,this is key2"] = 2

	}
	fmt.Println(items)
}
