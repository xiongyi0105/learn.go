package main

import "fmt"

func fff() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
func main() {
	fmt.Println(fff())
}
