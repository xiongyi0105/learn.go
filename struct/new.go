package main

import "fmt"

type vcard struct {
	name    rune
	address int64
	birth   string
}

func main() {
	type T struct {
		a int
		b string
	}

	// 这是类型T的一个实例或对象
	var t *T = new(T)
	fmt.Println(t.a)
	fmt.Println(t.b)
	// 初始化一个结构体实例的更简便和惯用方式
	type struct1 struct {
		a int
		b float64
		c [5]int
	}
	m := &struct1{1, 1.5, [5]int{1, 2, 3, 4, 5}}
	fmt.Println(m.c)
	Vcard := new(vcard)
	Vcard.Vcard(123123412341)
}

func (Vcard *vcard) Vcard(address int64) {
	Vcard.address = address
	fmt.Printf("%v", Vcard.address)
}
