package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}

func main() {
	m := myTime{time.Now()}
	fmt.Printf("Full Time Now is:%s\n", m.String())
	fmt.Printf("Time Now's First 3 chars:%v\n", m.first3Chars())
}
