package main

import (
	"fmt"
)

type Circulo struct {
	raio float64
}

func Area(circulo Circulo) float64 {
	pi := 3.14
	area := pi * circulo.raio * circulo.raio
	return area

}

func main() {

	p := Circulo{raio: 0}

	fmt.Println("raio do circulo")
	fmt.Scan(&p.raio)

	fmt.Printf("A area é : %.2f", Area(p))

}
