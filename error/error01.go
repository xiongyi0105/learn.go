package main

import (
	"errors"
	"fmt"
)

//Go 对于 error只提供了一个 error的接口，实现Error()方法

var errorNotFound error = errors.New("NotFound")

func main() {
	// ll := errorNotFound.Error()
	// fmt.Println(ll)
	fmt.Println(errorNotFound)
}
