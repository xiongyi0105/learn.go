package main

import "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

type car struct {
	make  string
	mode1 string
	price float32
}

type valuable interface {
	getValue() float32
	Shaper
}

type circle struct {
	price float32
}

type Shaper interface {
	getAll()
}

func (c circle) getValue() float32 {
	return c.price
}

func (c circle) getAll() {
	fmt.Println(c)
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

func (s stockPosition) getAll() {
	fmt.Println(s)
}

func (c car) getValue() float32 {
	return c.price
}

func (c car) getAll() {
	fmt.Println(c)
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

func main() {
	var o valuable = stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	var c valuable = car{"BMW", "G", 30000}
	showValue(c)
	var shape valuable = circle{2000}
	shape.getAll()
}
