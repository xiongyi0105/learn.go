package main

import (
	"fmt"
	"strconv"
)

type T struct {
	a int
	b float32
	c string
}

func main() {
	t := &T{7, -2.35, "abc\\tdef"}
	fmt.Printf("%v\n", t) //Printf等都会调用String(),故不要在String方法中再次使用调用String的方法会导致内存泄漏
}

func (t *T) String() string {
	return strconv.Itoa(t.a) + "/" + strconv.FormatFloat(float64(t.b), 'f', 6, 64) + "/" + "\"" + t.c + "\""
}
