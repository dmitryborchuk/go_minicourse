package main

import "fmt"

type Rectangle struct {
	height uint
	width  uint
}

func (d Rectangle) getArea() uint {
	return d.width * d.height
}

func main() {
	var test Rectangle = Rectangle{height: 10, width: 20}
	fmt.Println(test.getArea())
}
