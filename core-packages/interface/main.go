package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length, Width float64
}

func (r *Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func main() {
	var myRectangle Shape = &Rectangle{10, 20}
	fmt.Println(myRectangle.Area())
	fmt.Println(myRectangle.Perimeter())

	var myCircle Shape = &Circle{10}
	fmt.Println(myCircle.Area())
	fmt.Println(myCircle.Perimeter())
}
