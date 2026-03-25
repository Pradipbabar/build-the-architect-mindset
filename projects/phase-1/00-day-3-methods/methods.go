// Write a Go program that does the following:

// 1. Define an interface called Shape with two methods:
//    - Area() float64
//    - Perimeter() float64

// 2. Define two structs: Rectangle (width, height float64) and Circle (radius float64)

// 3. Implement the Shape interface on both Rectangle and Circle
//    - Rectangle Area    = width * height
//    - Rectangle Perimeter = 2 * (width + height)
//    - Circle Area       = π * radius * radius  (use math.Pi)
//    - Circle Perimeter  = 2 * π * radius

// 4. Write a function called printShape(s Shape) that prints:
//    - The area and perimeter of the shape
//    - The concrete type of the shape (hint: use a type switch)

// 5. In main(), create one Rectangle and one Circle, call printShape() on each

// Allowed imports: "fmt", "math"

package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	redius float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c Circle) Area() float64 {
	return math.Pi * c.redius * c.redius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.redius
}

func printShape(s Shape) {
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())

	// Add-on
	switch v := s.(type) {
	case Rectangle:
		fmt.Println("Rectangle", v)
	case Circle:
		fmt.Println("Circle", v)
	default:
		fmt.Println("Unkown")
	}

}
func main() {

	rectangle := Rectangle{12.44, 23.22}
	// fmt.Println(rectangle.Area())
	// fmt.Println(rectangle.Perimeter())

	circle := Circle{13.2}
	// fmt.Println(circle.Area())
	// fmt.Println(circle.Perimeter())

	printShape(rectangle)
	printShape(circle)

}
