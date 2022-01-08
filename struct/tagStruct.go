package main

import (
	"fmt"
	"reflect"
)

type structTag struct {
	domain string "domain of Lightwan"
	ip     string "ip address"
	no     int    "no"
}

func main() {
	appextag := structTag{"www.lightwan.com", "10.5.0.254", 2}
	for i := 0; i < 3; i++ {
		reftag(appextag, i)

	}
}

//reflect.Typeof 获取变量的正确类型，如果变量是一个结构体类型就可以Field来索引结构体字段，然后获取Tag属性
func reftag(tag structTag, ix int) {
	tagType := reflect.TypeOf(tag)
	ixField := tagType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}
