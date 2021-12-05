package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
包 strings 中的 Map 函数和 strings.IndexFunc() 一样都是非常好的使用例子。请学习它的源代码并基于该函数书写一个程序，
要求将指定文本内的所有非 ASCII 字符替换成 ? 或空格。您需要怎么做才能删除这些字符呢？
*/
func main() {
	var ll = "熊易&ioasjfhda(8*$_@!"
	fmt.Println(ll)
	// runell := []rune(ll)
	ll = strings.Map(replaceNotASCII, ll)
	fmt.Println(ll)
}

func replaceNotASCII(r rune) rune {
	if r > unicode.MaxASCII {
		r = '?'
	}
	return r
}
