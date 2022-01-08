package main

import (
	"fmt"
	"learn.go/struct/matrix"
)

func main() {
	wrong := new(matrix.matrix) // 编译失败（matrix 是私有的）
	right := matrix.Newmatrix()
	fmt.Println(right)
}
