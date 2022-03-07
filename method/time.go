package main

import (
	"fmt"
	"time"
)

func main() {
	t := new(time.Time)
	fmt.Printf("现在是当地时间：%s", t.Local().String())
}
