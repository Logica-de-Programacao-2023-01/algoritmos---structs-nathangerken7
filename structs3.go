package main

import "fmt"

type Triangulo struct {
	base   int
	altura int
}

func area(triangulo Triangulo) int {

	Area := triangulo.base * (triangulo.altura / 2)
	return Area
}

func main() {

	p := Triangulo{
		base:   0,
		altura: 0,
	}
	fmt.Println("altura")
	fmt.Scan(&p.altura)
	fmt.Println("base")
	fmt.Scan(&p.base)

	fmt.Println("A area é:", area(p))

}
