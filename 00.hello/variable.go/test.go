////package main
////
////var a = "G"
////
////func main() {
////	n()
////	m()
////	n()
////}
////
////func n() { print(a) }
////
////func m() {
////	a := "O"
////	print(a)
////}
//package main
//
//var a = "G"
//
//func main() {
//	n()
//	m()
//	n()
//}
//
//func n() {
//	print(a)
//}
//
//func m() {
//	a = "O"
//	print(a)
//}

//
package main

import (
	"fmt"
	"strings"
)

var newReader *strings.Reader

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string %#v have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	newReader = strings.NewReader(str)
	fmt.Println(newReader.ReadByte())
}
