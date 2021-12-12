package main

import (
	"bytes"
	"fmt"
	"reflect"
	"time"
)

func main() {
	// 实际就是new一个Buffer实例  直接赋值也可以，其实是对Buffer结构体的访问
	start := time.Now()
	var buffer *bytes.Buffer = &(bytes.Buffer{})
	bytesSlice := []string{"类似于", "dqw ", "qwoidhqoifhqwo", "12351235123"}
	for _, oneele := range bytesSlice {
		if reflect.TypeOf("1") == reflect.TypeOf(oneele) {
			buffer.WriteString(oneele)

		}
	}
	fmt.Println(buffer.String(), "\n")
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("Runfunc time:%s", delta)
}
