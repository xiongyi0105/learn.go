package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

func main() {
	myString := "熊叉和谢叉伊"
	stringSlice := []byte(myString)
	stringSlice[1] = 'y'
	stringSlice[4] = 'j'
	fmt.Println(string(stringSlice)) //�y��j�和谢叉伊
	runeSlice := []rune(myString)
	runeSlice[1] = '易'
	runeSlice[4] = '谨'
	myRuneSlice := string(runeSlice)
	fmt.Printf("myRuneSlice:%c,%v\n", myRuneSlice, myRuneSlice) //myRuneSlice:%!c(string=熊易和谢谨伊),熊易和谢谨伊
	fmt.Printf("myRuneSlice's length:%v\n", utf8.RuneCountInString(myRuneSlice))
	newSlice := strings.Split(myRuneSlice, "")
	fmt.Println(newSlice)
	//sort.Strings(stringSlice)
	sort.Strings(newSlice)
	fmt.Println(newSlice)
	stringsIndex := sort.SearchStrings(newSlice, "和")
	fmt.Println(stringsIndex)
}
