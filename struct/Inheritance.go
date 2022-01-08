package main

import "fmt"

type D struct {
	D float64
	b int
}

func main() {
	d := new(D)
	d.D = 6.1
	d.b = 3
	fmt.Printf("%+v\n", *d)
}
