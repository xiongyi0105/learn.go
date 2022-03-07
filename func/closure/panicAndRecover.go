package main

import "fmt"

func main() {
	standard := judge()
	standard(-1)
}

func judge() func(bmi int) {

	return func(bmi int) {
		// 若 panic 引起程序错误，recover可抓到错误且不让程序退出，defer得放在程序出错前，不然放在后面func还没准备好就准备退出程序？
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("程序错误")
			}
		}()
		if bmi > 10 {
			fmt.Println("you are fat")
		} else if bmi > 0 {
			fmt.Println("you are strong")
		} else {
			panic("bmi is illegal")
		}

	}
}
