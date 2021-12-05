package main

import "fmt"

//事实证明切片相加是会出错的
func main() {
	s := []int{2, 3, 4, 5, 6, 7}
	s1 := s[2:3]
	s2 := s[1:2]
	//不给编译
	//if s == s1+s2
	fmt.Println(s1)
	fmt.Println(s2)
	//不给编译
	//fmt.Println(s1+s2)
}
