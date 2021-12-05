package main

import (
	"fmt"
)

func main() {
	var lenString string = "hello world"
	//可以重新赋值字符串 lenString = "xiongyi"
	println(lenString)
	ll := lenString[len(lenString)-1] // 函数len()获取的是字节的长度
	println(ll)                       // 100   通过索引只能拿到字符串的纯字节内容，且只对ASCII码有效 ，取出来的是一个uint8类型
	joinstring := "by Xy"
	newString := lenString + joinstring
	fmt.Printf("%+v\n", newString)
	fmt.Println("输入一个字符？")
	var (
	//_astring int
	//_bint    int
	//_cint    int
	//_newW *bufio.Writer
	)
	//xx, err := fmt.Scanln(&astring)
	//xx, err := fmt.Scanf("%d %d %d", &astring, &bint, &cint)
	//newW = bufio.NewWriter(os.Stdout)

	//xx, err := fmt.Fprintln(newW, &astring)
	//fmt.Printf("扫描得到的个数%d，出现的错误%s,扫描得到的参数%d", xx, err, astring)
	//fmt.Printf("hai%v", astring)
	//fmt.Println(astring)
	//xx, err := fmt.Sscanln("hello", &astring)
	//fmt.Println(xx, err)
	//fmt.Println(astring)
	var xy string //默认是None
	fmt.Println(xy)
}
