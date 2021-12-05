package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	aa()
	log1()
}

func trace(task string) string {
	fmt.Println("enter:", task)
	return task
}

func untrace(task string) {
	fmt.Println("leave:", task)
}

func ff() {
	defer untrace(trace("ff"))
	fmt.Println("in ff ")
}

func aa() {
	defer untrace(trace("aa"))
	fmt.Println("in aa")
	ff()
}

func log1() (num int, err error) {
	defer func() { log.Printf("func log1:%d ,%v", num, err) }()
	return 7, io.EOF
}

//package main
//
//import "fmt"
//
//func trace(s string) string {
//	fmt.Println("entering:", s)
//	return s
//}
//
//func un(s string) {
//	fmt.Println("leaving:", s)
//}
//
//func a() {
//	defer un(trace("a"))
//	fmt.Println("in a")
//}
//
//func b() {
//	defer un(trace("b"))
//	fmt.Println("in b")
//	a()
//}
//
//func main() {
//	b()
//}
