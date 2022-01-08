package main

import "fmt"

type B struct {
	thing int
	hello int
}

func (b B) write() string {
	return fmt.Sprint(b)
}

func (b *B) change() {
	b.thing = 1

}

func main() {
	var b1 B
	b1.change()
	fmt.Println(b1.write())
}
